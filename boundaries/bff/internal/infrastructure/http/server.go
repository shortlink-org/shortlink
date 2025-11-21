package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	cors2 "github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/riandyrn/otelchi"
	flight_trace_middleware "github.com/shortlink-org/go-sdk/http/middleware/flight_trace"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/go-sdk/http/handler"
	auth_middleware "github.com/shortlink-org/go-sdk/http/middleware/auth"
	csrf_middleware "github.com/shortlink-org/go-sdk/http/middleware/csrf"
	logger_middleware "github.com/shortlink-org/go-sdk/http/middleware/logger"
	metrics_middleware "github.com/shortlink-org/go-sdk/http/middleware/metrics"
	pprof_labels_middleware "github.com/shortlink-org/go-sdk/http/middleware/pprof_labels"
	span_middleware "github.com/shortlink-org/go-sdk/http/middleware/span"
	http_server "github.com/shortlink-org/go-sdk/http/server"

	serverAPI "github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/api"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/controllers/cqrs"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/controllers/link"
	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/controllers/sitemap"
)

// MAX_AGE CORS - 5 minutes
const MAX_AGE = 300

// Run HTTP-server
func (api *Server) run(config Config) error {
	viper.SetDefault("BASE_PATH", "/api") // Base path for API endpoints

	api.ctx = config.Ctx
	api.jsonpb = protojson.MarshalOptions{
		UseProtoNames: true,
	}

	r := chi.NewRouter()

	// CORS configuration with CSRF support
	cors := cors2.New(cors2.Options{
		AllowedOrigins: []string{"*"},
		// HTTP methods: Added PUT and PATCH to support RESTful API operations
		// These methods are commonly used for updating resources (PUT for full updates, PATCH for partial updates)
		// Along with standard methods: GET (read), POST (create), DELETE (remove), OPTIONS (preflight)
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		// Headers: Include CSRF token headers and common request headers
		// X-CSRF-Token: Required for CSRF protection token validation
		// X-Requested-With: Standard header for AJAX requests, helps identify the request type
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Requested-With"},
		// Expose CSRF token header so clients can read it for subsequent requests
		ExposedHeaders:   []string{"X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           MAX_AGE,
	})

	r.Use(cors.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// A good base middleware stack
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(config.Http.Timeout))

	// CSRF Protection - must be early in the middleware stack
	r.Use(csrf_middleware.Middleware(api.log, config.Config))

	// Additional middleware
	r.Use(otelchi.Middleware(viper.GetString("SERVICE_NAME")))
	r.Use(logger_middleware.Logger(config.Log))
	r.Use(middleware.Recoverer)
	r.Use(span_middleware.Span())
	r.Use(auth_middleware.Auth(config.Config))
	r.Use(pprof_labels_middleware.Labels)
	r.Use(flight_trace_middleware.DebugTraceMiddleware(config.FlightTrace, config.Log, config.Config))

	// Metrics
	metrics, err := metrics_middleware.NewMetrics()
	if err != nil {
		return err
	}
	r.Use(metrics)

	r.NotFound(handler.NotFoundHandler)

	// Init routes
	controller := &Controller{
		Controller: link.NewController(config.Log, config.Link_rpc),
		LinkCQRSController: cqrs.LinkCQRSController{
			LinkCommandServiceClient: config.Link_command,
			LinkQueryServiceClient:   config.Link_query,
		},
		SitemapController: sitemap.SitemapController{
			// SitemapServiceClient: config.Sitemap_rpc,
		},
	}

	r.Mount(viper.GetString("BASE_PATH"), serverAPI.HandlerFromMux(controller, r))

	srv := http_server.New(config.Ctx, r, config.Http, config.Tracer, config.Config)

	// start HTTP-server
	config.Log.Info(config.I18n.Sprintf("BFF Web run on port %d", config.Http.Port))
	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// New API Provider for DI
func New(params Config, log logger.Logger) (*Server, error) {
	// API port
	viper.SetDefault("API_PORT", 7070) //nolint:mnd
	// Request Timeout (seconds)
	viper.SetDefault("API_TIMEOUT", "60s")

	// CSRF Protection settings
	viper.SetDefault("CSRF_TRUSTED_ORIGINS", "")
	viper.SetDefault("CSRF_TRUSTED_ORIGINS_ENV", "CSRF_TRUSTED_ORIGINS")

	params.Http = http_server.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	api := &Server{
		log: log,
	}
	g := errgroup.Group{}

	g.Go(func() error {
		return api.run(params)
	})

	return api, nil
}
