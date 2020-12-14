//go:generate wire
//+build wireinject
// The build tag makes sure the stub is not built in the final build.

/*
Main DI-package
*/
package di

import (
	"context"
	"net/http"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	link_store "github.com/batazor/shortlink/internal/api/infrastructure/store"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/logger"
	meta_store "github.com/batazor/shortlink/internal/metadata/infrastructure/store"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/pkg/rpc"
)

// Service - heplers
type Service struct {
	Log    logger.Logger
	Tracer opentracing.Tracer
	// TracerClose func()
	Sentry        *sentryhttp.Handler
	DB            *db.Store
	LinkStore     *link_store.LinkStore
	MetaStore     *meta_store.MetaStore
	MQ            mq.MQ
	ServerRPC     *rpc.RPCServer
	ClientRPC     *grpc.ClientConn
	Monitoring    *http.ServeMux
	PprofEndpoint PprofEndpoint
}

// InitLinkStore
func InitLinkStore(ctx context.Context, log logger.Logger, conn *db.Store) (*link_store.LinkStore, error) {
	st := link_store.LinkStore{}
	linkStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return linkStore, nil
}

// InitMetaStore
func InitMetaStore(ctx context.Context, log logger.Logger, conn *db.Store) (*meta_store.MetaStore, error) {
	st := meta_store.MetaStore{}
	metaStore, err := st.Use(ctx, log, conn)
	if err != nil {
		return nil, err
	}

	return metaStore, nil
}

func InitLogger(ctx context.Context) (logger.Logger, func(), error) {
	viper.SetDefault("LOG_LEVEL", logger.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := logger.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	log, err := logger.NewLogger(logger.Zap, conf)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		// flushes buffer, if any
		_ = log.Close() // nolint errcheck
	}

	return log, cleanup, nil
}

func InitSentry() (*sentryhttp.Handler, func(), error) {
	viper.SetDefault("SENTRY_DSN", "") // key for sentry
	DSN := viper.GetString("SENTRY_DSN")

	if DSN == "" {
		return nil, func() {}, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: DSN,
	})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {

		// Since sentry emits events in the background we need to make sure
		// they are sent before we shut down
		sentry.Flush(time.Second * 5)
		sentry.Recover()
	}

	// Create an instance of sentryhttp
	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	return sentryHandler, cleanup, nil
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(InitAutoMaxProcs, InitLogger, InitTracer)

// FullService =========================================================================================================
var FullSet = wire.NewSet(
	DefaultSet,
	NewFullService,
	InitStore,
	InitSentry,
	InitMonitoring,
	InitProfiling,
	InitMQ,
	rpc.RunGRPCServer,
	rpc.RunGRPCClient,
	InitLinkStore,
)

func NewFullService(
	log logger.Logger,
	mq mq.MQ,
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer opentracing.Tracer,
	db *db.Store,
	linkStore *link_store.LinkStore,
	pprofHTTP PprofEndpoint,
	autoMaxProcsOption diAutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Log:    log,
		MQ:     mq,
		Tracer: tracer,
		// TracerClose: cleanup,
		Monitoring:    monitoring,
		Sentry:        sentryHandler,
		DB:            db,
		LinkStore:     linkStore,
		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

func InitializeFullService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(FullSet))
}
