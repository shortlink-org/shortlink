/*
API Gateway - is a common entry point for all an external requests.
*/

package domain

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"

	link_cqrs "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/sitemap/v1"
	http_server "github.com/shortlink-org/shortlink/internal/pkg/http/server"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

// API - general describe of API
type API interface {
	Run(
		ctx context.Context,
		i18n *message.Printer,
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
