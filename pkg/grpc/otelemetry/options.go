// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package otelemetry

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

// Instrumentation interface
type Instrumentation interface {
	Extract(ctx context.Context, metadata *metadata.MD) (baggage.Baggage, trace.SpanContext)
	Inject(ctx context.Context, metadata *metadata.MD)
}

// InstrumentationOptions is a group of options for this instrumentation.
type InstrumentationOptions struct {
	propagators    propagation.TextMapPropagator
	tracerProvider trace.TracerProvider
	name           string
}

func (o *InstrumentationOptions) Inject(ctx context.Context, metadata *metadata.MD) {
	o.propagators.Inject(ctx, &metadataInfo{
		metadata: metadata,
	})
}

func (o *InstrumentationOptions) Extract(ctx context.Context, metadata *metadata.MD) (baggage.Baggage, trace.SpanContext) {
	ctx = o.propagators.Extract(ctx, &metadataInfo{
		metadata: metadata,
	})

	return baggage.FromContext(ctx), trace.SpanContextFromContext(ctx)
}

// InstrumentationOption function
type InstrumentationOption func(*InstrumentationOptions)

// NewInstrumentation creates a configuration for an instrumentation using a set of given options
func NewInstrumentation(opts []InstrumentationOption) *InstrumentationOptions {
	c := &InstrumentationOptions{
		propagators:    otel.GetTextMapPropagator(),
		tracerProvider: trace.NewNoopTracerProvider(),
		name:           "",
	}
	for _, option := range opts {
		option(c)
	}

	return c
}

func (o InstrumentationOptions) apply(opts ...InstrumentationOption) {
	for _, opt := range opts {
		opt(&o)
	}
}

func (o InstrumentationOptions) NewTracer(opts ...trace.TracerOption) trace.Tracer {
	return o.tracerProvider.Tracer(o.name, opts...)
}

// WithPropagators sets propagators
func WithPropagators(propagators propagation.TextMapPropagator) InstrumentationOption {
	return func(options *InstrumentationOptions) {
		options.propagators = propagators
	}
}

// WithTraceProvider sets trace provider
func WithTraceProvider(tracerProvider trace.TracerProvider) InstrumentationOption {
	return func(options *InstrumentationOptions) {
		options.tracerProvider = tracerProvider
	}
}

// WithInstrumentationName sets instrumentation name
func WithInstrumentationName(name string) InstrumentationOption {
	return func(options *InstrumentationOptions) {
		options.name = name
	}
}

var _ Instrumentation = &InstrumentationOptions{}

type metadataInfo struct {
	metadata *metadata.MD
}

// assert that metadataSupplier implements the TextMapCarrier interface
var _ propagation.TextMapCarrier = &metadataInfo{}

func (s *metadataInfo) Get(key string) string {
	values := s.metadata.Get(key)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

func (s *metadataInfo) Set(key string, value string) {
	s.metadata.Set(key, value)
}

func (s *metadataInfo) Keys() []string {
	out := make([]string, 0, len(*s.metadata))
	for key := range *s.metadata {
		out = append(out, key)
	}
	return out
}
