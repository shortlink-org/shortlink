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
	"golang.org/x/text/message"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	shortctx "github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/flags"
	logger_di "github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	mq_di "github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/permission"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/store"
	traicing_di "github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/internal/pkg/cache"
	"github.com/shortlink-org/shortlink/internal/pkg/db"
	short_i18n "github.com/shortlink-org/shortlink/internal/pkg/i18n"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/rpc"
)

// Service - heplers
type Service struct {
	// Common
	Ctx  context.Context
	Cfg  *config.Config
	Log  logger.Logger
	I18N *message.Printer

	// Security
	Auth *authzed.Client

	// Delivery
	DB        *db.Store
	Cache     redisCache.UniversalClient
	MQ        *mq.DataBus
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
	short_i18n.New,
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
	i18n *message.Printer,

	// Delivery
	serverRPC *rpc.Server,
	clientRPC *grpc.ClientConn,
	dataBus *mq.DataBus,
	store_db *db.Store,
	shortcache redisCache.UniversalClient,

	// Observability
	monitor *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
) (*Service, error) {
	return &Service{
		// Common
		Ctx:  ctx,
		Cfg:  cfg,
		Log:  log,
		I18N: i18n,

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
