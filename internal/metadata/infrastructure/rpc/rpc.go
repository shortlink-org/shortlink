package rpc

import (
	"context"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/metadata/application"
	rpc "github.com/batazor/shortlink/internal/metadata/domain"
)

type MetadataServer struct {
	rpc.UnimplementedMetadataServer

	config struct {
		grpc_port string
	}
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

func New() (*MetadataServer, error) {
	server := MetadataServer{}
	server.setConfig()

	// Run gRPC server
	lis, err := net.Listen("tcp", server.config.grpc_port)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()
	rpc.RegisterMetadataServer(s, &server)
	if err := s.Serve(lis); err != nil {
		return nil, err
	}

	return &server, nil
}

// setConfig - set configuration
func (s *MetadataServer) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("GRPC_PORT", ":50051") // gRPC port
	s.config.grpc_port = viper.GetString("GRPC_PORT")
}
