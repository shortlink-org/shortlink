/*
Link Service. Infrastructure layer. RPC EndpointRPC Endpoint
*/

package v1

import (
	sitemap_application "github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/sitemap"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type Sitemap struct {
	SitemapServiceServer

	service *sitemap_application.Service
	log     logger.Logger
}

func New(runRPCServer *rpc.Server, application *sitemap_application.Service, log logger.Logger) (*Sitemap, error) {
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
