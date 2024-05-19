/*
Main DI-package
*/
package di

import (
	"context"

	"github.com/authzed/authzed-go/v1"
	"github.com/google/wire"
	redisCache "github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/pkg/cache"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	shortctx "github.com/shortlink-org/shortlink/pkg/di/pkg/context"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/flags"
	logger_di "github.com/shortlink-org/shortlink/pkg/di/pkg/logger"
	mq_di "github.com/shortlink-org/shortlink/pkg/di/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/permission"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/store"
	traicing_di "github.com/shortlink-org/shortlink/pkg/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

// Service - helpers
type Service struct {
	// Common
	Ctx context.Context
	Cfg *config.Config
	Log logger.Logger

	// Security
	Auth *authzed.Client

	// Delivery
	DB        db.DB
	Cache     redisCache.UniversalClient
	MQ        mq.MQ
	ServerRPC *rpc.Server
	ClientRPC *grpc.ClientConn

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro
}

// DefaultSet ==========================================================================================================
var DefaultSet = wire.NewSet(
	shortctx.New,
	autoMaxPro.New,
	flags.New,
	config.New,
	logger_di.New,
	traicing_di.New,
	monitoring.New,
	cache.New,
	profiling.New,
	permission.New,
)

// FullSet =============================================================================================================
var FullSet = wire.NewSet(
	DefaultSet,
	NewFullService,
	store.New,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
)

func NewFullService(
	// Common
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,

	// Delivery
	serverRPC *rpc.Server,
	clientRPC *grpc.ClientConn,
	dataBus mq.MQ,
	store_db db.DB,
	shortcache redisCache.UniversalClient,

	// Observability
	monitor *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		// Common
		Ctx: ctx,
		Cfg: cfg,
		Log: log,

		// Delivery
		MQ:        dataBus,
		DB:        store_db,
		Cache:     shortcache,
		ClientRPC: clientRPC,
		ServerRPC: serverRPC,

		// Observability
		Tracer:        tracer,
		Monitoring:    monitor,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,
	}, nil
}

func InitializeFullService() (*Service, func(), error) {
	panic(wire.Build(FullSet))
}
