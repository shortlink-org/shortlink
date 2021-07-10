//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package billing_di

import (
	"context"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	account_application "github.com/batazor/shortlink/internal/services/billing/application/account"
	"github.com/batazor/shortlink/internal/services/billing/application/payment"
	tariff_application "github.com/batazor/shortlink/internal/services/billing/application/tariff"
	api "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/balance/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/order/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/payment/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/tariff/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

type BillingService struct {
	Logger logger.Logger

	// Delivery
	httpAPIServer    *api.Server
	balanceRPCServer *balance_rpc.Balance
	orderRPCServer   *order_rpc.Order
	paymentRPCServer *payment_rpc.Payment
	tariffRPCServer  *tariff_rpc.Tariff

	// Application
	payment *payment.Payment

	// Repository
	accountRepository *billing_store.AccountRepository
	balanceRepository *billing_store.BalanceRepository
	orderRepository   *billing_store.OrderRepository
	paymentRepository *billing_store.PaymentRepository
	tariffRepository  *billing_store.TariffRepository
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

func InitializeBillingService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq mq.MQ, tracer *opentracing.Tracer) (*BillingService, func(), error) {
	panic(wire.Build(BillingSet))
}
