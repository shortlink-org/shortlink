/*
Link Service. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package v1

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	sitemap_application "github.com/shortlink-org/shortlink/internal/services/link/application/sitemap"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type Sitemap struct {
	SitemapServiceServer

	service *sitemap_application.Service
	log     logger.Logger
}

func New(runRPCServer *rpc.RPCServer, application *sitemap_application.Service, log logger.Logger) (*Sitemap, error) {
	server := &Sitemap{
		// Create Service Application
		service: application,

		log: log,
	}

	// Register services
	if runRPCServer != nil {
		RegisterSitemapServiceServer(runRPCServer.Server, server)
	}

	return server, nil
}
