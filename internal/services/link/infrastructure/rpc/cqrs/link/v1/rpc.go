/*
Link Service. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package v1

import (
	"github.com/batazor/shortlink/internal/pkg/logger"
	link_application "github.com/batazor/shortlink/internal/services/link/application/link"
	"github.com/batazor/shortlink/pkg/rpc"
)

type Link struct {
	LinkCommandServiceServer
	LinkQueryServiceServer

	service *link_application.Service
	log     logger.Logger
}

func New(runRPCServer *rpc.RPCServer, application *link_application.Service, log logger.Logger) (*Link, error) {
	server := &Link{
		// Create Service Application
		service: application,

		log: log,
	}

	// Register services
	RegisterLinkCommandServiceServer(runRPCServer.Server, server)
	RegisterLinkQueryServiceServer(runRPCServer.Server, server)

	return server, nil
}
