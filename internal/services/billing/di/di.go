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
	"github.com/batazor/shortlink/internal/services/billing/application"
	api "github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/balance/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/order/v1"
	payment_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/payment/v1"
	tariff_rpc "github.com/batazor/shortlink/internal/services/billing/infrastructure/rpc/tariff/v1"
	billing_store "github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
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
	//mq *api_mq.Event

	// Application
	payment *application.Payment

	// Repository
	billingStore *billing_store.BillingStore
}

// BillingService ======================================================================================================
var BillingSet = wire.NewSet(
	// infrastructure
	NewBillingAPIServer,
	//NewBalanceRPCServer,
	//NewOrderRPCServer,
	//NewPaymentRPCServer,
	//NewTariffRPCServer,

	// applications
	//NewPaymentApplication,

	// repository
	NewBillingStore,

	NewBillingService,
)

func NewBillingAPIServer(ctx context.Context, logger logger.Logger, tracer *opentracing.Tracer, rpcServer *rpc.RPCServer) (*api.Server, error) {
	// Run API server
	API := api.Server{}

	apiService, err := API.Use(ctx, logger, tracer)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

//func NewBalanceRPCServer() {}
//
//func NewOrderRPCServer() {}
//
//func NewPaymentRPCServer() {}
//
//func NewTariffRPCServer() {}
//
//func NewPaymentApplication() {}

func NewBillingStore(ctx context.Context, logger logger.Logger, db *db.Store) (*billing_store.BillingStore, error) {
	store := &billing_store.BillingStore{}
	billingStore, err := store.Use(ctx, logger, db)
	if err != nil {
		return nil, err
	}

	return billingStore, nil
}

func NewBillingService(
	log logger.Logger,

	// Delivery
	httpAPIServer *api.Server,
	//balanceRPCServer *balance_rpc.Balance,
	//orderRPCServer   *order_rpc.Order,
	//paymentRPCServer *payment_rpc.Payment,
	//tariffRPCServer  *tariff_rpc.Tariff,

	// Application
	//payment *application.Payment,
) (*BillingService, error) {
	return &BillingService{
		Logger: log,

		httpAPIServer: httpAPIServer,
		//balanceRPCServer: balanceRPCServer,
		//orderRPCServer:   orderRPCServer,
		//paymentRPCServer: paymentRPCServer,
		//tariffRPCServer:  tariffRPCServer,

		//payment:          payment,
	}, nil
}

func InitializeBillingService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq mq.MQ, tracer *opentracing.Tracer) (*BillingService, func(), error) {
	panic(wire.Build(BillingSet))
}
