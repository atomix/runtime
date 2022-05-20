// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	"github.com/atomix/runtime/api/atomix/map/v1"
	"github.com/atomix/runtime/pkg/errors"
	"github.com/atomix/runtime/pkg/primitive"
)

func newMapV1ManagerServer(proxies *primitive.Service[Map]) v1.MapManagerServer {
	return &mapV1ManagerServer{
		proxies: proxies,
	}
}

type mapV1ManagerServer struct {
	proxies *primitive.Service[Map]
}

func (s *mapV1ManagerServer) Create(ctx context.Context, request *v1.CreateRequest) (*v1.CreateResponse, error) {
	conn, err := s.proxies.Connect(ctx, request.Primitive)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = conn.Create(ctx, request.Primitive.PrimitiveID)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &v1.CreateResponse{}, nil
}

func (s *mapV1ManagerServer) Close(ctx context.Context, request *v1.CloseRequest) (*v1.CloseResponse, error) {
	conn, err := s.proxies.GetConn(request.PrimitiveID)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	err = conn.Close(ctx, request.PrimitiveID)
	if err != nil {
		return nil, errors.ToProto(err)
	}
	return &v1.CloseResponse{}, nil
}

var _ v1.MapManagerServer = (*mapV1ManagerServer)(nil)