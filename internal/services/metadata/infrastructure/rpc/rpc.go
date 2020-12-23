//go:generate protoc -I. -I../../domain --go_out=Minternal/metadata/domain/meta.proto=.:. --go-grpc_out=Minternal/metadata/domain/meta.proto=.:. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative metadata_rpc.proto

/*
Metadata Service. Infrastructure layer
*/
package metadata_rpc

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/services/metadata/application"
	meta_store "github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

type Metadata struct {
	service *application.Service
	log     logger.Logger
}

func New(runRPCServer *rpc.RPCServer, st *meta_store.MetaStore, log logger.Logger) (*Metadata, error) {
	server := &Metadata{
		// Create Service Application
		service: &application.Service{
			Store: st,
		},
		log: log,
	}

	// Register services
	RegisterMetadataServer(runRPCServer.Server, server)
	runRPCServer.Run()

	return server, nil
}

func (m *Metadata) Get(ctx context.Context, req *GetMetaRequest) (*GetMetaResponse, error) {
	meta, err := m.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &GetMetaResponse{
		Meta: meta,
	}, nil
}

func (m *Metadata) Set(ctx context.Context, req *SetMetaRequest) (*SetMetaResponse, error) {
	meta, err := m.service.Set(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &SetMetaResponse{
		Meta: meta,
	}, nil
}

func (m *Metadata) mustEmbedUnimplementedMetadataServer() {}
