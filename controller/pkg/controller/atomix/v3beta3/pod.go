// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v3beta3

import (
	"context"
	atomixv3beta2 "github.com/atomix/runtime/controller/pkg/apis/atomix/v3beta3"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"time"
)

const (
	podFinalizer = "atomix.io/pod"
)

func addPodController(mgr manager.Manager) error {
	// Create a new controller
	c, err := controller.New("pod-controller", mgr, controller.Options{
		Reconciler: &PodReconciler{
			client: mgr.GetClient(),
			scheme: mgr.GetScheme(),
			config: mgr.GetConfig(),
			events: mgr.GetEventRecorderFor("atomix"),
		},
		RateLimiter: workqueue.NewItemExponentialFailureRateLimiter(time.Millisecond*10, time.Second*5),
	})
	if err != nil {
		return err
	}

	// Watch for changes to Pods
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
		if object.GetAnnotations()[proxyInjectStatusAnnotation] == injectedStatus &&
			object.GetAnnotations()[proxyProfileAnnotation] == object.GetName() {
			return []reconcile.Request{
				{
					NamespacedName: types.NamespacedName{
						Namespace: object.GetNamespace(),
						Name:      object.GetName(),
					},
				},
			}
		}
		return nil
	}))
	if err != nil {
		return err
	}

	// Watch for changes to StorageProfiles
	err = c.Watch(&source.Kind{Type: &atomixv3beta2.StorageProfile{}}, handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
		podList := &corev1.PodList{}
		if err := mgr.GetClient().List(context.Background(), podList, &client.ListOptions{Namespace: object.GetNamespace()}); err != nil {
			log.Error(err)
			return nil
		}

		var requests []reconcile.Request
		for _, pod := range podList.Items {
			if pod.Annotations[proxyInjectStatusAnnotation] == injectedStatus &&
				pod.Annotations[proxyProfileAnnotation] == object.GetName() {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: pod.Namespace,
						Name:      pod.Name,
					},
				})
			}
		}
		return requests
	}))
	if err != nil {
		return err
	}

	// Watch for changes to DataStores
	err = c.Watch(&source.Kind{Type: &atomixv3beta2.DataStore{}}, handler.EnqueueRequestsFromMapFunc(func(object client.Object) []reconcile.Request {
		podList := &corev1.PodList{}
		if err := mgr.GetClient().List(context.Background(), podList, &client.ListOptions{}); err != nil {
			log.Error(err)
			return nil
		}

		var requests []reconcile.Request
		for _, pod := range podList.Items {
			if pod.Annotations[proxyInjectStatusAnnotation] == injectedStatus {
				requests = append(requests, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: pod.Namespace,
						Name:      pod.Name,
					},
				})
			}
		}
		return requests
	}))
	if err != nil {
		return err
	}
	return nil
}

// PodReconciler is a Reconciler for Profiles
type PodReconciler struct {
	client client.Client
	scheme *runtime.Scheme
	config *rest.Config
	events record.EventRecorder
}

// Reconcile reconciles StorageProfile resources
func (r *PodReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log.Infof("Reconciling Pod '%s'", request.NamespacedName)
	pod := &corev1.Pod{}
	err := r.client.Get(context.TODO(), request.NamespacedName, pod)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		log.Error(err)
		return reconcile.Result{}, err
	}

	profileName, ok := pod.Annotations[proxyProfileAnnotation]
	if !ok {
		return reconcile.Result{}, nil
	}

	profileNamespacedName := types.NamespacedName{
		Namespace: pod.Namespace,
		Name:      profileName,
	}
	profile := &atomixv3beta2.StorageProfile{}
	if err := r.client.Get(ctx, profileNamespacedName, profile); err != nil {
		if !k8serrors.IsNotFound(err) {
			log.Error(err)
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	for _, binding := range profile.Spec.Bindings {
		storeName := types.NamespacedName{
			Namespace: binding.Store.Namespace,
			Name:      binding.Store.Name,
		}
		if storeName.Namespace == "" {
			storeName.Namespace = profile.Namespace
		}
		store := &atomixv3beta2.DataStore{}
		if err := r.client.Get(ctx, storeName, store); err != nil {
			if !k8serrors.IsNotFound(err) {
				log.Error(err)
				return reconcile.Result{}, err
			}
		} else {
			var podStatus *atomixv3beta2.PodStatus
			for _, ps := range store.Status.PodStatuses {
				if ps.Namespace == store.Namespace && ps.Name == store.Name {
					podStatus = &ps
					break
				}
			}
			if podStatus == nil {
				store.Status.PodStatuses = append(store.Status.PodStatuses, atomixv3beta2.PodStatus{
					ObjectReference: corev1.ObjectReference{
						Namespace: pod.Namespace,
						Name:      pod.Name,
					},
					State: atomixv3beta2.PodBindingPending,
				})
				if err := r.client.Status().Update(ctx, store); err != nil {
					log.Error(err)
					return reconcile.Result{}, err
				}
			}
		}
	}
	return reconcile.Result{}, nil
}
