/*
Link Service. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package v1

import (
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/link_cqrs"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type LinkRPC struct {
	LinkCommandServiceServer
	LinkQueryServiceServer

	cqrs *link_cqrs.Service
	log  logger.Logger
}

func New(runRPCServer *rpc.Server, application *link_cqrs.Service, log logger.Logger) (*LinkRPC, error) {
	server := &LinkRPC{
		// Create Service Application
		cqrs: application,

		log: log,
	}

	// Register services
	if runRPCServer != nil {
		RegisterLinkCommandServiceServer(runRPCServer.Server, server)
		RegisterLinkQueryServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}
