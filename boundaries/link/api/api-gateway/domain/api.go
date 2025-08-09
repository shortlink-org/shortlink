/*
API Gateway - is a common entry point for all an external requests.
*/

package domain

import (
	"context"

	link_cqrs "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/cqrs/link/v1/linkv1grpc"
	link_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/link/v1/linkv1grpc"
	sitemap_rpc "buf.build/gen/go/shortlink-org/shortlink-link-link/grpc/go/infrastructure/rpc/sitemap/v1/sitemapv1grpc"
	"go.opentelemetry.io/otel/trace"

	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

// API - general describe of API
type API interface {
	Run(
		ctx context.Context,
		config http_server.Config,
		log logger.Logger,
		tracer trace.TracerProvider,

		// delivery
		link_rpc link_rpc.LinkServiceClient,
		link_command link_cqrs.LinkCommandServiceClient,
		link_query link_cqrs.LinkQueryServiceClient,
		sitemap_rpc sitemap_rpc.SitemapServiceClient,
	) error
}
