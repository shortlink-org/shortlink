/*
API
*/

package api_application

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/batazor/shortlink/internal/pkg/logger"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	link_cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
)

// API - general describe of API
type API interface { // nolint unused
	Run(
		ctx context.Context,
		config api_type.Config,
		log logger.Logger,
		tracer *opentracing.Tracer,

		// delivery
		link_rpc link_rpc.LinkServiceClient,
		link_command link_cqrs.LinkCommandServiceClient,
		link_query link_cqrs.LinkQueryServiceClient,
		sitemap_rpc sitemap_rpc.SitemapServiceClient,
	) error
}
