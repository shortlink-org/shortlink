package rpc

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/metadata/application"
	rpc "github.com/batazor/shortlink/internal/metadata/domain"
)

type MetadataServer struct {
	rpc.UnimplementedMetadataServer
}

func (m *MetadataServer) Get(ctx context.Context, req *rpc.GetMetaRequest) (*rpc.GetMetaResponse, error) {
	service := application.Repository{}
	meta, err := service.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &rpc.GetMetaResponse{
		Meta: meta,
	}, nil
}

func (m *MetadataServer) Set(ctx context.Context, req *rpc.SetMetaRequest) (*rpc.SetMetaResponse, error) {
	service := application.Repository{}
	meta, err := service.Set(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &rpc.SetMetaResponse{
		Meta: meta,
	}, nil
}

func New(grpcServer *grpc.Server) (*MetadataServer, error) {
	server := MetadataServer{}
	server.setConfig()

	// Register services
	rpc.RegisterMetadataServer(grpcServer, &server)

	return &server, nil
}

// setConfig - set configuration
func (s *MetadataServer) setConfig() { // nolint unused
	viper.AutomaticEnv()
}
