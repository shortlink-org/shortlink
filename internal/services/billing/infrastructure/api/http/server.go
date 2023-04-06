package api

import (
	"context"

	http_server "github.com/shortlink-org/shortlink/pkg/http/server"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	account_application "github.com/shortlink-org/shortlink/internal/services/billing/application/account"
	order_application "github.com/shortlink-org/shortlink/internal/services/billing/application/order"
	payment_application "github.com/shortlink-org/shortlink/internal/services/billing/application/payment"
	tariff_application "github.com/shortlink-org/shortlink/internal/services/billing/application/tariff"
	http_chi "github.com/shortlink-org/shortlink/internal/services/billing/infrastructure/api/http/http-chi"
)

// API - general describe of API
type API interface {
	Run(
		ctx context.Context,
		db *db.Store,
		config http_server.Config,
		log logger.Logger,
		tracer *trace.TracerProvider,

		// services
		accountService *account_application.AccountService,
		orderService *order_application.OrderService,
		paymentService *payment_application.PaymentService,
		tariffService *tariff_application.TariffService,
	) error
}

type Server struct{}

func (s *Server) Use(
	ctx context.Context,
	db *db.Store,
	log logger.Logger,
	tracer *trace.TracerProvider,

	// services
	accountService *account_application.AccountService,
	orderService *order_application.OrderService,
	paymentService *payment_application.PaymentService,
	tariffService *tariff_application.TariffService,
) (*Server, error) {
	var server API

	viper.SetDefault("API_TYPE", "http-chi") // Select: http-chi
	// API port
	viper.SetDefault("API_PORT", 7070) // nolint:gomnd
	// Request Timeout (seconds)
	viper.SetDefault("API_TIMEOUT", "60s")

	config := http_server.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT"),
	}

	serverType := viper.GetString("API_TYPE")

	switch serverType {
	case "http-chi":
		server = &http_chi.API{}
	default:
		server = &http_chi.API{}
	}

	g := errgroup.Group{}

	g.Go(func() error {
		return server.Run(
			ctx,
			db,
			config,
			log,
			tracer,

			accountService,
			orderService,
			paymentService,
			tariffService,
		)
	})

	return s, nil
}
