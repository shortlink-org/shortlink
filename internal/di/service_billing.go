//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	billing_di "github.com/batazor/shortlink/internal/services/billing/di"
	"github.com/batazor/shortlink/pkg/rpc"
)

type ServiceBilling struct {
	Service

	BillingService *billing_di.BillingService
}

// InitMetaService =====================================================================================================
func InitBillingService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db *db.Store, mq v1.MQ, tracer *opentracing.Tracer) (*billing_di.BillingService, func(), error) {
	return billing_di.InitializeBillingService(ctx, runRPCClient, runRPCServer, log, db, mq, tracer)
}

// BillingService =======================================================================================================
var BillingSet = wire.NewSet(
	DefaultSet,
	store.New,
	rpc.InitServer,
	rpc.InitClient,
	mq_di.New,
	sentry.New,
	monitoring.New,
	InitBillingService,
	NewBillingService,
)

func NewBillingService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	mq v1.MQ,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	billingService *billing_di.BillingService,
) (*ServiceBilling, error) {
	return &ServiceBilling{
		Service: Service{
			Ctx:        ctx,
			Log:        log,
			MQ:         mq,
			Tracer:     tracer,
			Monitoring: monitoring,
		},

		BillingService: billingService,
	}, nil
}

func InitializeBillingService() (*ServiceBilling, func(), error) {
	panic(wire.Build(BillingSet))
}
