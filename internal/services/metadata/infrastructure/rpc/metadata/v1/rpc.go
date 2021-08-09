/*
Metadata Service. Infrastructure layer
*/
package v1

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/logger"
	metadata "github.com/batazor/shortlink/internal/services/metadata/application"
	"github.com/batazor/shortlink/pkg/rpc"
)

type Metadata struct {
	MetadataServiceServer

	service *metadata.Service
	log     logger.Logger
}

func New(runRPCServer *rpc.RPCServer, application *metadata.Service, log logger.Logger) (*Metadata, error) {
	server := &Metadata{
		service: application,
		log:     log,
	}

	// Register services
	RegisterMetadataServiceServer(runRPCServer.Server, server)
	go runRPCServer.Run()

	return server, nil
}

func (m *Metadata) Get(ctx context.Context, req *MetadataServiceGetRequest) (*MetadataServiceGetResponse, error) {
	meta, err := m.service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &MetadataServiceGetResponse{
		Meta: meta,
	}, nil
}

func (m *Metadata) Set(ctx context.Context, req *MetadataServiceSetRequest) (*MetadataServiceSetResponse, error) {
	meta, err := m.service.Set(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &MetadataServiceSetResponse{
		Meta: meta,
	}, nil
}

func (m *Metadata) mustEmbedUnimplementedMetadataServer() {}
