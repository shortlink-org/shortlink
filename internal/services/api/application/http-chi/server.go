package http_chi

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/text/message"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/batazor/shortlink/internal/pkg/i18n"
	"github.com/batazor/shortlink/internal/pkg/logger"
	cqrs_api "github.com/batazor/shortlink/internal/services/api/application/http-chi/controllers/cqrs"
	link_api "github.com/batazor/shortlink/internal/services/api/application/http-chi/controllers/link"
	sitemap_api "github.com/batazor/shortlink/internal/services/api/application/http-chi/controllers/sitemap"
	api_type "github.com/batazor/shortlink/internal/services/api/application/type"
	link_cqrs "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/cqrs/link/v1"
	link_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/link/v1"
	sitemap_rpc "github.com/batazor/shortlink/internal/services/link/infrastructure/rpc/sitemap/v1"
	"github.com/batazor/shortlink/pkg/http/handler"
	additionalMiddleware "github.com/batazor/shortlink/pkg/http/middleware"
)

// Run HTTP-server
func (api *API) Run(
	ctx context.Context,
	i18n *message.Printer,
	config api_type.Config,
	log logger.Logger,
	tracer *opentracing.Tracer,

	// delivery
	link_rpc link_rpc.LinkServiceClient,
	link_command link_cqrs.LinkCommandServiceClient,
	link_query link_cqrs.LinkQueryServiceClient,
	sitemap_rpc sitemap_rpc.SitemapServiceClient,
) error { // nolint unparam
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
		MaxAge:           300,
		//Debug:            true,
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
	r.Use(additionalMiddleware.NewTracing(tracer))
	r.Use(additionalMiddleware.Logger(log))

	r.NotFound(handler.NotFoundHandler)

	r.Mount("/api/link", link_api.Routes(link_rpc))
	r.Mount("/api/cqrs", cqrs_api.Routes(link_command, link_query))
	r.Mount("/api/sitemap", sitemap_api.Routes(sitemap_rpc))

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: r,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},

		ReadTimeout:       1 * time.Second,                 // the maximum duration for reading the entire request, including the body
		WriteTimeout:      config.Timeout + 30*time.Second, // the maximum duration before timing out writes of the response
		IdleTimeout:       30 * time.Second,                // the maximum amount of time to wait for the next request when keep-alive is enabled
		ReadHeaderTimeout: 2 * time.Second,                 // the amount of time allowed to read request headers
	}

	// start HTTP-server
	log.Info(i18n.Sprintf("API run on port %d", config.Port))
	err := srv.ListenAndServe()
	return err
}
