//go:generate wire
//go:build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Billing Service DI-package
*/
package billing_di

import (
	"context"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/trace"

	api "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/http"
	order_rpc "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/rpc/order/v1"
	payment_rpc "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/rpc/payment/v1"
	tariff_rpc "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/api/rpc/tariff/v1"
	account_repository "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/repository/account"
	tariff_repository "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/infrastructure/repository/tariff"
	account_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/account"
	order_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/order"
	payment_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/payment"
	tariff_application "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/usecases/tariff"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/store"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

type BillingService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	httpAPIServer    *api.Server
	orderRPCServer   *order_rpc.Order
	paymentRPCServer *payment_rpc.Payment
	tariffRPCServer  *tariff_rpc.Tariff

	// Repository
	accountRepository    account_repository.Repository
	tariffRepository     tariff_repository.Repository
	eventStoreRepository *eventsourcing.EventSourcing
}

// BillingService ======================================================================================================
var BillingSet = wire.NewSet(
	di.DefaultSet,

	// Delivery
	rpc.InitServer,
	rpc.InitClient,
	store.New,

	// Infrastructure
	NewBillingAPIServer,

	// repository
	eventsourcing.New,

	// application
	NewTariffApplication,
	NewAccountApplication,
	NewOrderApplication,
	NewPaymentApplication,

	NewBillingService,
)

func NewAccountApplication(ctx context.Context, log logger.Logger, db db.DB) (*account_application.AccountService, error) {
	accountService, err := account_application.New(ctx, log, db)
	if err != nil {
		return nil, err
	}

	return accountService, nil
}

func NewOrderApplication(log logger.Logger, eventStore eventsourcing.EventSourcing) (*order_application.OrderService, error) {
	orderService, err := order_application.New(log, eventStore)
	if err != nil {
		return nil, err
	}

	return orderService, nil
}

func NewPaymentApplication(log logger.Logger, eventStore eventsourcing.EventSourcing) (*payment_application.PaymentService, error) {
	paymentService, err := payment_application.New(log, eventStore)
	if err != nil {
		return nil, err
	}

	return paymentService, nil
}

func NewTariffApplication(ctx context.Context, log logger.Logger, db db.DB) (*tariff_application.TariffService, error) {
	tariffService, err := tariff_application.New(ctx, log, db)
	if err != nil {
		return nil, err
	}

	return tariffService, nil
}

func NewBillingAPIServer(
	// Common
	ctx context.Context,
	log logger.Logger,
	tracer trace.TracerProvider,

	// Applications
	accountService *account_application.AccountService,
	orderService *order_application.OrderService,
	paymentService *payment_application.PaymentService,
	tariffService *tariff_application.TariffService,
) (*api.Server, error) {
	// Run API server
	API := api.Server{}

	apiService, err := API.Use(
		// Common
		ctx,
		log,
		tracer,

		// services
		accountService,
		orderService,
		paymentService,
		tariffService,
	)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewBillingService(
	// Common
	log logger.Logger,
	config *config.Config,

	// Observability
	monitoring *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	// Delivery
	httpAPIServer *api.Server,
) (*BillingService, error) {
	return &BillingService{
		// Common
		Log:    log,
		Config: config,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitoring,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		// Delivery
		httpAPIServer: httpAPIServer,
	}, nil
}

func InitializeBillingService() (*BillingService, func(), error) {
	panic(wire.Build(BillingSet))
}
