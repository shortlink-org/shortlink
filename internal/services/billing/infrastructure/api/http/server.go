package api

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	account_application "github.com/batazor/shortlink/internal/services/billing/application/account"
	order_application "github.com/batazor/shortlink/internal/services/billing/application/order"
	payment_application "github.com/batazor/shortlink/internal/services/billing/application/payment"
	tariff_application "github.com/batazor/shortlink/internal/services/billing/application/tariff"
	http_chi "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http/http-chi"
	api_type "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http/type"
)

// API - general describe of API
type API interface { // nolint unused
	Run(
		ctx context.Context,
		db *db.Store,
		config api_type.Config,
		log logger.Logger,
		tracer *opentracing.Tracer,

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
	tracer *opentracing.Tracer,

	// services
	accountService *account_application.AccountService,
	orderService *order_application.OrderService,
	paymentService *payment_application.PaymentService,
	tariffService *tariff_application.TariffService,
) (*Server, error) {
	var server API

	viper.SetDefault("API_TYPE", "http-chi") // Select: http-chi
	viper.SetDefault("API_PORT", 7070)       // API port
	viper.SetDefault("API_TIMEOUT", 60)      // Request Timeout (seconds)

	config := api_type.Config{
		Port:    viper.GetInt("API_PORT"),
		Timeout: viper.GetDuration("API_TIMEOUT") * time.Second, // nolint durationcheck
	}

	serverType := viper.GetString("API_TYPE")

	switch serverType {
	case "http-chi":
		server = &http_chi.API{}
	default:
		server = &http_chi.API{}
	}

	if err := server.Run(
		ctx,
		db,
		config,
		log,
		tracer,

		accountService,
		orderService,
		paymentService,
		tariffService,
	); err != nil {
		return nil, err
	}

	return s, nil
}
