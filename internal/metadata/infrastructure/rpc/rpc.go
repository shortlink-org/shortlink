//go:generate protoc -I../../domain --go_out=Minternal/metadata/domain/rpc.proto=.:. --go-grpc_out=Minternal/metadata/domain/rpc.proto=.:. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative rpc.proto

/*
Metadata Service. Infrastructure layer
*/
package rpc

import (
	"context"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/metadata/application"
	meta_store "github.com/batazor/shortlink/internal/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

func New(runRPCServer *rpc.RPCServer, st *meta_store.MetaStore, log logger.Logger) (*MetadataServer, error) {
	server := MetadataServer{
		// Create Service Application
		service: &application.Service{
			Store: st,
		},
		log: log,
	}

	// Register services
	RegisterMetadataServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return &server, nil
}

func (m *MetadataServer) Get(ctx context.Context, req *GetMetaRequest) (*GetMetaResponse, error) {
	meta, err := m.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &GetMetaResponse{
		Meta: meta,
	}, nil
}

func (m *MetadataServer) Set(ctx context.Context, req *SetMetaRequest) (*SetMetaResponse, error) {
	meta, err := m.service.Set(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &SetMetaResponse{
		Meta: meta,
	}, nil
}

func (m *MetadataServer) mustEmbedUnimplementedMetadataServer() {}
