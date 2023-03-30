package http_chi

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/riandyrn/otelchi"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/shortlink-org/shortlink/internal/pkg/i18n"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	cqrs_api "github.com/shortlink-org/shortlink/internal/services/api/application/http-chi/controllers/cqrs"
	link_api "github.com/shortlink-org/shortlink/internal/services/api/application/http-chi/controllers/link"
	sitemap_api "github.com/shortlink-org/shortlink/internal/services/api/application/http-chi/controllers/sitemap"
	link_cqrs "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/shortlink-org/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	"github.com/shortlink-org/shortlink/pkg/http/handler"
	additionalMiddleware "github.com/shortlink-org/shortlink/pkg/http/middleware"
	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
)

const MAX_AGE = 300

// Run HTTP-server
// @title Shortlink API
// @version 1.0
// @description Shortlink API
// @termsOfService http://swagger.io/terms/
//
// @contact.name Shortlink repository
// @contact.url https://github.com/shortlink-org/shortlink/issues
// @contact.email support@swagger.io
//
// @license.name MIT
// @license.url http://www.opensource.org/licenses/MIT
//
// @host shortlink.best
// @BasePath /api
// @schemes http https
func (api *API) Run(
	ctx context.Context,
	i18n *message.Printer,
	config http_server.Config,
	log logger.Logger,
	tracer *trace.TracerProvider,

	// Delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) error {

	api.ctx = ctx
	api.jsonpb = protojson.MarshalOptions{
		UseProtoNames: true,
	}

	log.Info("Run HTTP-CHI API")

	r := chi.NewRouter()

	// CORS
	corsPolicy := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           MAX_AGE,
		// Debug:            true,
	})

	r.Use(corsPolicy.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// A good base middleware stack
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(config.Timeout))

	// Additional middleware
	r.Use(otelchi.Middleware(viper.GetString("SERVICE_NAME")))
	r.Use(additionalMiddleware.Logger(log))

	r.NotFound(handler.NotFoundHandler)

	r.Mount("/api/links", link_api.Routes(link_rpc))
	r.Mount("/api/cqrs", cqrs_api.Routes(link_command, link_query))
	r.Mount("/api/sitemap", sitemap_api.Routes(sitemap_rpc))

	srv := http_server.New(ctx, r, config)

	// start HTTP-server
	log.Info(i18n.Sprintf("API run on port %d", config.Port))
	err := srv.ListenAndServe()

	return err
}
