// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package otelemetry

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	grpccodes "google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// UnaryServerTelemetryInterceptor returns a grpc.UnaryServerInterceptor suitable
// for use in a grpc.NewServer call.
func UnaryServerTelemetryInterceptor(opts ...InstrumentationOption) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
	) (interface{}, error) {
		requestMetadata, _ := metadata.FromIncomingContext(ctx)
		metadataCopy := requestMetadata.Copy()

		instrumentation := NewInstrumentation(opts)
		bags, spanCtx := instrumentation.Extract(ctx, &metadataCopy)
		ctx = baggage.ContextWithBaggage(ctx, bags)

		tracer := instrumentation.NewTracer(trace.WithInstrumentationVersion("1.0.0"))

		name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx))
		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			name,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(attr...),
		)

		defer span.End()
		e := event{
			messageType: otelgrpc.RPCMessageTypeReceived,
			message:     req,
			id:          1,
		}
		e.Send(ctx)

		resp, err := handler(ctx, req)
		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(codes.Error, s.Message())
			span.SetAttributes(statusCodeAttr(s.Code()))
			e = event{
				messageType: otelgrpc.RPCMessageTypeSent,
				message:     s.Proto(),
				id:          1,
			}
			e.Send(ctx)

		} else {
			span.SetAttributes(statusCodeAttr(grpccodes.OK))
			e = event{
				messageType: otelgrpc.RPCMessageTypeSent,
				message:     resp,
				id:          1,
			}
		}

		return resp, err
	}
}

// StreamTelemetryServerInterceptor returns a grpc.StreamServerInterceptor suitable
// for use in a grpc.NewServer call.
func StreamTelemetryServerInterceptor(opts ...InstrumentationOption) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := stream.Context()

		requestMetadata, _ := metadata.FromIncomingContext(ctx)
		metadataCopy := requestMetadata.Copy()

		instrumentation := NewInstrumentation(opts)
		bags, spanCtx := instrumentation.Extract(ctx, &metadataCopy)
		ctx = baggage.ContextWithBaggage(ctx, bags)

		tracer := instrumentation.NewTracer(trace.WithInstrumentationVersion("1.0.0"))

		name, attr := spanInfo(info.FullMethod, peerFromCtx(ctx))
		ctx, span := tracer.Start(
			trace.ContextWithRemoteSpanContext(ctx, spanCtx),
			name,
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(attr...),
		)

		defer span.End()
		err := handler(srv, wrapServerStream(ctx, stream))

		if err != nil {
			s, _ := status.FromError(err)
			span.SetStatus(codes.Error, s.Message())
			span.SetAttributes(statusCodeAttr(s.Code()))
		} else {
			span.SetAttributes(statusCodeAttr(grpccodes.OK))
		}

		return err
	}
}
