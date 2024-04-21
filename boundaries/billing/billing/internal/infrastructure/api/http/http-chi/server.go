package http_chi

import (
	"context"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/riandyrn/otelchi"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/http/http-chi/controllers/account"
	"github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/http/http-chi/controllers/balance"
	"github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/http/http-chi/controllers/order"
	"github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/http/http-chi/controllers/payment"
	"github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/http/http-chi/controllers/tariff"
	account_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/account"
	order_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/order"
	payment_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/payment"
	tariff_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/tariff"
	"github.com/shortlink-org/shortlink/pkg/http/handler"
	auth_middleware "github.com/shortlink-org/shortlink/pkg/http/middleware/auth"
	logger_middleware "github.com/shortlink-org/shortlink/pkg/http/middleware/logger"
	metrics_middleware "github.com/shortlink-org/shortlink/pkg/http/middleware/metrics"
	pprof_labels_middleware "github.com/shortlink-org/shortlink/pkg/http/middleware/pprof_labels"
	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

// Run HTTP-server
func (api *API) Run(
	ctx context.Context,
	config http_server.Config,
	log logger.Logger,
	tracer trace.TracerProvider,

	// Services
	accountService *account_application.AccountService,
	orderService *order_application.OrderService,
	paymentService *payment_application.PaymentService,
	tariffService *tariff_application.TariffService,
) error {
	api.jsonpb = protojson.MarshalOptions{
		UseProtoNames: true,
	}

	log.Info("Run HTTP-CHI API")

	r := chi.NewRouter()

	// CORS
	corsHandler := cors.New(cors.Options{ //nolint:gocritic // ignore
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, //nolint:mnd // ignore
		// Debug:            true,
	})

	r.Use(corsHandler.Handler)
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

	metrics, err := metrics_middleware.NewMetrics()
	if err != nil {
		return err
	}
	r.Use(metrics)

	r.NotFound(handler.NotFoundHandler)

	// Init routes
	accountRoutes, err := account.New(accountService)
	if err != nil {
		return err
	}

	balanceRoutes, err := balance.New(paymentService)
	if err != nil {
		return err
	}

	orderRoutes, err := order.New(orderService)
	if err != nil {
		return err
	}

	paymentRoutes, err := payment.New(paymentService)
	if err != nil {
		return err
	}

	tariffRoutes, err := tariff.New(tariffService)
	if err != nil {
		return err
	}

	r.Mount("/api/billing", r.Group(func(router chi.Router) {
		accountRoutes.Routes(router)
		balanceRoutes.Routes(router)
		orderRoutes.Routes(router)
		paymentRoutes.Routes(router)
		tariffRoutes.Routes(router)
	}))

	srv := http_server.New(ctx, r, config, tracer)

	// start HTTP-server
	log.Info(fmt.Sprintf("API run on port %d", config.Port))
	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
