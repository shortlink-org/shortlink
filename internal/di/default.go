//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	ctx "github.com/batazor/shortlink/internal/di/internal/context"
	"github.com/batazor/shortlink/internal/di/internal/flags"
	logger_di "github.com/batazor/shortlink/internal/di/internal/logger"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/profiling"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/di/internal/traicing"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	meta_store "github.com/batazor/shortlink/internal/services/metadata/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
)

// Service - heplers
type Service struct {
	Ctx           context.Context
	Cfg           *config.Config
	Log           logger.Logger
	Tracer        *opentracing.Tracer
	Sentry        *sentryhttp.Handler
	DB            *db.Store
	MetaStore     *meta_store.MetaStore
	MQ            v1.MQ
	ServerRPC     *rpc.RPCServer
	ClientRPC     *grpc.ClientConn
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(ctx.New, autoMaxPro.New, flags.New, config.New, logger_di.New, traicing_di.New)

// FullService =========================================================================================================
var FullSet = wire.NewSet(
	DefaultSet,
	NewFullService,
	store.New,
	sentry.New,
	monitoring.New,
	profiling.New,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
)

func NewFullService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	mq v1.MQ,
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	db *db.Store,
	//linkStore *link_store.LinkStore,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Ctx:           ctx,
		Cfg:           cfg,
		Log:           log,
		MQ:            mq,
		Tracer:        tracer,
		Monitoring:    monitoring,
		Sentry:        sentryHandler,
		DB:            db,
		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

func InitializeFullService() (*Service, func(), error) {
	panic(wire.Build(FullSet))
}
