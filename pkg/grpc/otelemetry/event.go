// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package otelemetry

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Event telemetry event
type event struct {
	id          int
	message     interface{}
	messageType attribute.KeyValue
}

// Send sends a telemetry event
func (e event) Send(ctx context.Context) {
	span := trace.SpanFromContext(ctx)
	if p, ok := e.message.(proto.Message); ok {
		span.AddEvent("event", trace.WithAttributes(
			e.messageType,
			otelgrpc.RPCMessageIDKey.Int(e.id),
			otelgrpc.RPCMessageUncompressedSizeKey.Int(proto.Size(p)),
		))
	} else {
		span.AddEvent("event", trace.WithAttributes(
			e.messageType,
			otelgrpc.RPCMessageIDKey.Int(e.id),
		))
	}

}
