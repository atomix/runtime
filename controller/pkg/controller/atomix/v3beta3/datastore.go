// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v3beta3

import (
	"context"
	"fmt"
	proxyv1 "github.com/atomix/runtime/api/atomix/proxy/v1"
	atomixv3beta2 "github.com/atomix/runtime/controller/pkg/apis/atomix/v3beta3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	dataStoreFinalizer = "atomix.io/store"
)

func addDataStoreController(mgr manager.Manager) error {
	// Create a new controller
	c, err := controller.New("data-store-controller", mgr, controller.Options{
		Reconciler: &DataStoreReconciler{
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

	// Watch for changes to Profiles
	err = c.Watch(&source.Kind{Type: &atomixv3beta2.DataStore{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to ConfigMap
	err = c.Watch(&source.Kind{Type: &corev1.ConfigMap{}}, &handler.EnqueueRequestForOwner{
		OwnerType: &atomixv3beta2.DataStore{},
	})
	if err != nil {
		return err
	}
	return nil
}

// DataStoreReconciler is a Reconciler for Profiles
type DataStoreReconciler struct {
	client client.Client
	scheme *runtime.Scheme
	config *rest.Config
	events record.EventRecorder
}

// Reconcile reconciles DataStore resources
func (r *DataStoreReconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	log.Infof("Reconciling DataStore '%s'", request.NamespacedName)
	dataStore := &atomixv3beta2.DataStore{}
	err := r.client.Get(context.TODO(), request.NamespacedName, dataStore)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		log.Error(err)
		return reconcile.Result{}, err
	}

	if dataStore.DeletionTimestamp != nil {
		if !hasFinalizer(dataStore, dataStoreFinalizer) {
			return reconcile.Result{}, nil
		}

		if changed, err := r.reconcileStatus(ctx, dataStore); err != nil {
			return reconcile.Result{}, err
		} else if changed {
			return reconcile.Result{Requeue: true}, nil
		}

		for _, podStatus := range dataStore.Status.PodStatuses {
			if podStatus.State != atomixv3beta2.PodBindingDisconnected {
				return reconcile.Result{}, nil
			}
		}
		removeFinalizer(dataStore, dataStoreFinalizer)
		if err := r.client.Update(ctx, dataStore); err != nil {
			log.Error(err)
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	if !hasFinalizer(dataStore, dataStoreFinalizer) {
		addFinalizer(dataStore, dataStoreFinalizer)
		if err := r.client.Update(ctx, dataStore); err != nil {
			log.Error(err)
			return reconcile.Result{}, err
		}
		return reconcile.Result{}, nil
	}

	if changed, err := r.reconcileStatus(ctx, dataStore); err != nil {
		return reconcile.Result{}, err
	} else if changed {
		return reconcile.Result{Requeue: true}, nil
	}
	return reconcile.Result{}, nil
}

func (r *DataStoreReconciler) reconcileStatus(ctx context.Context, dataStore *atomixv3beta2.DataStore) (bool, error) {
	statusChanged := false
	podStatuses := make([]atomixv3beta2.PodStatus, 0, len(dataStore.Status.PodStatuses))
	for _, podStatus := range dataStore.Status.PodStatuses {
		podName := types.NamespacedName{
			Namespace: podStatus.Namespace,
			Name:      podStatus.Name,
		}
		pod := &corev1.Pod{}
		if err := r.client.Get(ctx, podName, pod); err != nil {
			if !k8serrors.IsNotFound(err) {
				log.Error(err)
				return false, err
			} else {
				statusChanged = true
			}
		} else {
			if changed, err := r.reconcilePodStatus(ctx, dataStore, pod, &podStatus); err != nil {
				log.Error(err)
				return false, err
			} else if changed {
				statusChanged = true
			}
			podStatuses = append(podStatuses, podStatus)
		}
	}
	if statusChanged {
		dataStore.Status.PodStatuses = podStatuses
		if err := r.client.Status().Update(ctx, dataStore); err != nil {
			log.Error(err)
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func (r *DataStoreReconciler) reconcilePodStatus(ctx context.Context, dataStore *atomixv3beta2.DataStore, pod *corev1.Pod, status *atomixv3beta2.PodStatus) (bool, error) {
	if dataStore.DeletionTimestamp != nil {
		if status.State == atomixv3beta2.PodBindingPending {
			return false, nil
		}

		switch status.State {
		case atomixv3beta2.PodBindingConnecting, atomixv3beta2.PodBindingConnected, atomixv3beta2.PodBindingConfiguring:
			status.State = atomixv3beta2.PodBindingDisconnecting
			return true, nil
		case atomixv3beta2.PodBindingDisconnecting:
			conn, err := connect(ctx, pod)
			if err != nil {
				log.Error(err)
				return false, err
			}

			r.events.Eventf(pod, "Normal", "DisconnectStore", "Disconnecting store '%s'", getNamespacedName(dataStore))
			client := proxyv1.NewProxyClient(conn)
			request := &proxyv1.DisconnectRequest{
				StoreID: proxyv1.StoreId{
					Namespace: dataStore.Namespace,
					Name:      dataStore.Name,
				},
			}
			_, err = client.Disconnect(ctx, request)
			if err != nil {
				log.Error(err)
				r.events.Eventf(pod, "Warning", "DisconnectStoreFailed", "Failed disconnecting from store '%s': %s", getNamespacedName(dataStore), err)
				return false, err
			}
			r.events.Eventf(pod, "Normal", "DisconnectStoreSucceeded", "Successfully disconnected from store '%s'", getNamespacedName(dataStore))

			status.State = atomixv3beta2.PodBindingDisconnected
			status.Version = ""
			return true, nil
		}
		return false, nil
	}

	if status.State == atomixv3beta2.PodBindingPending {
		status.State = atomixv3beta2.PodBindingConnecting
		status.Version = dataStore.ResourceVersion
		return true, nil
	}

	switch status.State {
	case atomixv3beta2.PodBindingConnecting:
		if status.Version != dataStore.ResourceVersion {
			status.Version = dataStore.ResourceVersion
			return true, nil
		}

		conn, err := connect(ctx, pod)
		if err != nil {
			log.Error(err)
			return false, err
		}

		r.events.Eventf(pod, "Normal", "ConnectStore", "Connecting store '%s'", getNamespacedName(dataStore))
		client := proxyv1.NewProxyClient(conn)
		request := &proxyv1.ConnectRequest{
			StoreID: proxyv1.StoreId{
				Namespace: dataStore.Namespace,
				Name:      dataStore.Name,
			},
			DriverID: proxyv1.DriverId{
				Name:    dataStore.Spec.Driver.Name,
				Version: dataStore.Spec.Driver.Version,
			},
			Config: dataStore.Spec.Config.Raw,
		}
		_, err = client.Connect(ctx, request)
		if err != nil {
			log.Error(err)
			r.events.Eventf(pod, "Warning", "ConnectStoreFailed", "Failed connecting to store '%s': %s", getNamespacedName(dataStore), err)
			return false, err
		}
		r.events.Eventf(pod, "Normal", "ConnectStoreSucceeded", "Successfully connected to store '%s'", getNamespacedName(dataStore))

		status.State = atomixv3beta2.PodBindingConnected
		return true, nil
	case atomixv3beta2.PodBindingConnected:
		if status.Version != dataStore.ResourceVersion {
			status.State = atomixv3beta2.PodBindingConfiguring
			status.Version = dataStore.ResourceVersion
			return true, nil
		}
	case atomixv3beta2.PodBindingConfiguring:
		if status.Version != dataStore.ResourceVersion {
			status.Version = dataStore.ResourceVersion
			return true, nil
		}

		conn, err := connect(ctx, pod)
		if err != nil {
			log.Error(err)
			return false, err
		}

		r.events.Eventf(pod, "Normal", "ConfigureStore", "Configuring store '%s'", getNamespacedName(dataStore))
		client := proxyv1.NewProxyClient(conn)
		request := &proxyv1.ConfigureRequest{
			StoreID: proxyv1.StoreId{
				Namespace: dataStore.Namespace,
				Name:      dataStore.Name,
			},
			Config: dataStore.Spec.Config.Raw,
		}
		_, err = client.Configure(ctx, request)
		if err != nil {
			r.events.Eventf(pod, "Warning", "ConfigureStoreFailed", "Failed reconfiguring store '%s': %s", getNamespacedName(dataStore), err)
			log.Error(err)
			return false, err
		}
		r.events.Eventf(pod, "Normal", "ConfigureStoreSucceeded", "Successfully configured store '%s'", getNamespacedName(dataStore))

		status.State = atomixv3beta2.PodBindingConnected
		return true, nil
	default:
		status.State = atomixv3beta2.PodBindingConnecting
		status.Version = dataStore.ResourceVersion
		return true, nil
	}
	return false, nil
}

func connect(ctx context.Context, pod *corev1.Pod) (*grpc.ClientConn, error) {
	target := fmt.Sprintf("%s:%d", pod.Status.PodIP, defaultProxyPort)
	return grpc.DialContext(ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
