/*
Link Service. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package v1

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	link_application "github.com/shortlink-org/shortlink/internal/services/link/application/link"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type Link struct {
	LinkServiceServer

	log logger.Logger

	// Application
	service *link_application.Service
}

func New(runRPCServer *rpc.RPCServer, application *link_application.Service, log logger.Logger) (*Link, error) {
	server := &Link{
		// Create Service Application
		service: application,

		log: log,
	}

	// Register services
	if runRPCServer != nil {
		RegisterLinkServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}
