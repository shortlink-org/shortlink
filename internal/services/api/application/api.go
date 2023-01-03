/*
API
*/

package api_application

import (
	"context"

	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	link_cqrs "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
)

// API - general describe of API
type API interface {
	Run(
		ctx context.Context,
		i18n *message.Printer,
		config http_server.Config,
		log logger.Logger,
		tracer *trace.TracerProvider,

		// delivery
		link_rpc link_rpc.LinkServiceClient,
		link_command link_cqrs.LinkCommandServiceClient,
		link_query link_cqrs.LinkQueryServiceClient,
		sitemap_rpc sitemap_rpc.SitemapServiceClient,
	) error
}
