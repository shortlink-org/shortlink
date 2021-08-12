//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package billing_di

import (
	"context"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/db"
	event_store "github.com/batazor/shortlink/internal/pkg/eventsourcing/store"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	account_application "github.com/batazor/shortlink/internal/services/billing/application/account"
	order_application "github.com/batazor/shortlink/internal/services/billing/application/order"
	payment_application "github.com/batazor/shortlink/internal/services/billing/application/payment"
	tariff_application "github.com/batazor/shortlink/internal/services/billing/application/tariff"
	api "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http"
	order_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/order/v1"
	payment_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/payment/v1"
	tariff_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/tariff/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

type BillingService struct {
	Logger logger.Logger

	// Delivery
	httpAPIServer    *api.Server
	orderRPCServer   *order_rpc.Order
	paymentRPCServer *payment_rpc.Payment
	tariffRPCServer  *tariff_rpc.Tariff

	// Repository
	accountRepository    *billing_store.AccountRepository
	tariffRepository     *billing_store.TariffRepository
	eventStoreRepository *event_store.EventStore
}

// BillingService ======================================================================================================
var BillingSet = wire.NewSet(
	// infrastructure
	NewBillingAPIServer,

	// repository
	NewBillingStore,

	// application
	NewTariffApplication,
	NewAccountApplication,
	NewOrderApplication,
	NewPaymentApplication,

	NewBillingService,
)

func NewBillingStore(ctx context.Context, logger logger.Logger, db *db.Store) (*billing_store.BillingStore, error) {
	store := &billing_store.BillingStore{}
	billingStore, err := store.Use(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	return billingStore, nil
}

func NewAccountApplication(logger logger.Logger, store *billing_store.BillingStore) (*account_application.AccountService, error) {
	accountService, err := account_application.New(logger, store.Account)
	if err != nil {
		return nil, err
	}

	return accountService, nil
}

func NewOrderApplication(logger logger.Logger, store *billing_store.BillingStore) (*order_application.OrderService, error) {
	orderService, err := order_application.New(logger, store.EventStore)
	if err != nil {
		return nil, err
	}

	return orderService, nil
}

func NewPaymentApplication(logger logger.Logger, store *billing_store.BillingStore) (*payment_application.PaymentService, error) {
	paymentService, err := payment_application.New(logger, store.EventStore)
	if err != nil {
		return nil, err
	}

	return paymentService, nil
}

func NewTariffApplication(logger logger.Logger, store *billing_store.BillingStore) (*tariff_application.TariffService, error) {
	tariffService, err := tariff_application.New(logger, store.Tariff)
	if err != nil {
		return nil, err
	}

	return tariffService, nil
}

func NewBillingAPIServer(
	ctx context.Context,
	logger logger.Logger,
	tracer *opentracing.Tracer,
	rpcServer *rpc.RPCServer,
	db *db.Store,

	// Applications
	accountService *account_application.AccountService,
	orderService *order_application.OrderService,
	paymentService *payment_application.PaymentService,
	tariffService *tariff_application.TariffService,
) (*api.Server, error) {
	// Run API server
	API := api.Server{}

	apiService, err := API.Use(
		ctx,
		db,
		logger,
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
	log logger.Logger,

	// Delivery
	httpAPIServer *api.Server,
) (*BillingService, error) {
	return &BillingService{
		Logger: log,

		// Delivery
		httpAPIServer: httpAPIServer,
	}, nil
}

func InitializeBillingService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq v1.MQ, tracer *opentracing.Tracer) (*BillingService, func(), error) {
	panic(wire.Build(BillingSet))
}
