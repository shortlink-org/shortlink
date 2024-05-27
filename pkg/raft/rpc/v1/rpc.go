package v1

import (
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type Server struct {
	RaftServiceServer
}

func NewServer(runRPCServer *rpc.Server) (*Server, error) {
	server := &Server{}

	// Register services
	if runRPCServer != nil {
		RegisterRaftServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}

type Client struct {
	RaftServiceClient
}

func NewClient(runRPCClient *grpc.ClientConn) (*Client, error) {
	client := &Client{}

	// Register services
	if runRPCClient != nil {
		client.RaftServiceClient = NewRaftServiceClient(runRPCClient)
	}

	return client, nil
}
