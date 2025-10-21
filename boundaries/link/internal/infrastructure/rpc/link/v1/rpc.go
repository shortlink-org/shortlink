/*
Link UC. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package v1

import (
	rpc "github.com/shortlink-org/go-sdk/grpc"
	"github.com/shortlink-org/go-sdk/logger"
	link_application "github.com/shortlink-org/shortlink/boundaries/link/internal/usecases/link"
)

type LinkRPC struct {
	LinkServiceServer

	log logger.Logger

	// Application
	service *link_application.UC
}

func New(runRPCServer *rpc.Server, application *link_application.UC, log logger.Logger) (*LinkRPC, error) {
	server := &LinkRPC{
		// Create UC Application
		service: application,

		log: log,
	}

	// Register services
	if runRPCServer != nil {
		RegisterLinkServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}
