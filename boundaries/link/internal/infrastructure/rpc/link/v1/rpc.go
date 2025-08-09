/*
Link UC. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package v1

import (
	link_application "github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/link"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
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
