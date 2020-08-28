package rpc

import (
	"context"

	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/metadata/application"
	rpc "github.com/batazor/shortlink/internal/metadata/domain"
)

type MetadataServer struct {
	rpc.UnimplementedMetadataServer
	db db.DB
}

func (m *MetadataServer) Get(ctx context.Context, req *rpc.GetMetaRequest) (*rpc.GetMetaResponse, error) {
	service := application.Repository{
		DB: m.db,
	}
	meta, err := service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &rpc.GetMetaResponse{
		Meta: meta,
	}, nil
}

func (m *MetadataServer) Set(ctx context.Context, req *rpc.SetMetaRequest) (*rpc.SetMetaResponse, error) {
	service := application.Repository{
		DB: m.db,
	}
	meta, err := service.Set(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &rpc.SetMetaResponse{
		Meta: meta,
	}, nil
}

func New(runRPCServer *di.RPCServer, db db.DB) (*MetadataServer, error) {
	server := MetadataServer{
		db: db,
	}

	// Register services
	rpc.RegisterMetadataServer(runRPCServer.Server, &server)
	runRPCServer.Run()

	return &server, nil
}
