/*
Metadata Service. Infrastructure layer
*/
package rpc

import (
	"context"

	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/metadata/application"
	rpc "github.com/batazor/shortlink/internal/metadata/domain"
	meta_store "github.com/batazor/shortlink/internal/metadata/infrastructure/store"
)

type MetadataServer struct {
	service *application.Service
}

func New(runRPCServer *di.RPCServer, st *meta_store.MetaStore) (*MetadataServer, error) {
	server := MetadataServer{
		// Create Service Application
		service: &application.Service{
			Store: st,
		},
	}

	service := &rpc.MetadataService{
		Get: server.Get,
		Set: server.Set,
	}

	// Register services
	rpc.RegisterMetadataService(runRPCServer.Server, service)
	runRPCServer.Run()

	return &server, nil
}

func (m *MetadataServer) Get(ctx context.Context, req *rpc.GetMetaRequest) (*rpc.GetMetaResponse, error) {
	meta, err := m.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &rpc.GetMetaResponse{
		Meta: meta,
	}, nil
}

func (m *MetadataServer) Set(ctx context.Context, req *rpc.SetMetaRequest) (*rpc.SetMetaResponse, error) {
	meta, err := m.service.Set(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &rpc.SetMetaResponse{
		Meta: meta,
	}, nil
}
