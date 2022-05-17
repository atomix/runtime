// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package otelemetry

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

type wrappedServerStream struct {
	grpc.ServerStream
	ctx                    context.Context
	receivedMessageCounter int
	sentMessageCounter     int
}

func (w *wrappedServerStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)

	if err != nil {
		return err
	}
	w.receivedMessageCounter++
	e := event{
		messageType: otelgrpc.RPCMessageTypeReceived,
		id:          w.receivedMessageCounter,
		message:     m,
	}
	e.Send(w.Context())

	return nil
}

func (w *wrappedServerStream) SendMsg(m interface{}) error {
	err := w.ServerStream.SendMsg(m)
	w.sentMessageCounter++
	e := event{
		messageType: otelgrpc.RPCMessageTypeSent,
		message:     m,
		id:          w.sentMessageCounter,
	}
	e.Send(w.Context())
	return err
}

func wrapServerStream(ctx context.Context, stream grpc.ServerStream) *wrappedServerStream {
	return &wrappedServerStream{
		ServerStream: stream,
		ctx:          ctx,
	}
}
