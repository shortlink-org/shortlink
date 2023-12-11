package http

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	cors2 "github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/riandyrn/otelchi"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/internal/pkg/http/handler"
	auth_middleware "github.com/shortlink-org/shortlink/internal/pkg/http/middleware/auth"
	logger_middleware "github.com/shortlink-org/shortlink/internal/pkg/http/middleware/logger"
	metrics_middleware "github.com/shortlink-org/shortlink/internal/pkg/http/middleware/metrics"
	pprof_labels_middleware "github.com/shortlink-org/shortlink/internal/pkg/http/middleware/pprof_labels"

	http_server "github.com/shortlink-org/shortlink/internal/pkg/http/server"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	serverAPI "github.com/shortlink-org/shortlink/internal/services/bff-web/infrastructure/http/api"
)

// Run HTTP-server
func (api *Server) run(
	// Common
	ctx context.Context,
	config http_server.Config,
	log logger.Logger,
	tracer trace.TracerProvider,
) error {
	api.ctx = ctx
	api.jsonpb = protojson.MarshalOptions{
		UseProtoNames: true,
	}

	r := chi.NewRouter()

	// CORS
	cors := cors2.New(cors2.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{},
		AllowCredentials: true,
		MaxAge:           300, // nolint:gomnd
	})

	r.Use(cors.Handler)
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
	r.Use(logger_middleware.Logger(log))
	r.Use(auth_middleware.Auth())
	r.Use(pprof_labels_middleware.Labels)

	// Additional middlewares
	metrics, err := metrics_middleware.NewMetrics()
	if err != nil {
		return err
	}
	r.Use(metrics)

	r.NotFound(handler.NotFoundHandler)

	// Init routes
	r.Mount("/bff/web", serverAPI.HandlerFromMux(nil, r))

	srv := http_server.New(ctx, r, config, tracer)

	// start HTTP-server
	log.Info(fmt.Sprintf("BFF Web run on port %d", config.Port))
	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// API Provider for DI
func (s *Server) Run(
	// Common
	ctx context.Context,
	log logger.Logger,
	tracer trace.TracerProvider,
) (*Server, error) {
	// API port
	viper.SetDefault("API_PORT", 7070) // nolint:gomnd
	// Request Timeout (seconds)
	viper.SetDefault("API_TIMEOUT", "60s")

	config := http_server.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	g := errgroup.Group{}

	g.Go(func() error {
		return s.run(
			ctx,
			config,
			log,
			tracer,
		)
	})

	return s, nil
}
