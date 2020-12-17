//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/heptiolabs/healthcheck"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"

	link_store "github.com/batazor/shortlink/internal/api/infrastructure/store"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/logger"
	meta_store "github.com/batazor/shortlink/internal/metadata/infrastructure/store"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/traicing"
	"github.com/batazor/shortlink/pkg/rpc"
)

// Service - heplers
type Service struct {
	Ctx    context.Context
	Log    logger.Logger
	Tracer *opentracing.Tracer
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

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

// Cobra/Flags =========================================================================================================
func InitFlags() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	return rootCmd, nil
}

// Monitoring ==========================================================================================================
func InitMonitoring(sentryHandler *sentryhttp.Handler) *http.ServeMux {
	// Create a new Prometheus registry
	registry := prometheus.NewRegistry()

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(registry, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	// Create an "common" listener
	commonMux := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	commonMux.Handle("/metrics", sentryHandler.Handle(promhttp.Handler()))

	// Expose a liveness check on /live
	commonMux.HandleFunc("/live", sentryHandler.HandleFunc(health.LiveEndpoint))

	// Expose a readiness check on /ready
	commonMux.HandleFunc("/ready", sentryHandler.HandleFunc(health.ReadyEndpoint))

	return commonMux
}

// Profiling ===========================================================================================================
type PprofEndpoint *http.ServeMux

func InitProfiling() PprofEndpoint {
	// Create an "common" listener
	pprofMux := http.NewServeMux()

	// Registration pprof-handlers
	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return pprofMux
}

// Health ==============================================================================================================
func NewHealthCheck() (healthcheck.Handler, error) {
	// create a new health instance
	endpoint := healthcheck.NewHandler()

	// Expose the /live and /ready endpoints over HTTP
	go http.ListenAndServe("0.0.0.0:9090", endpoint)

	return endpoint, nil
}

// AutoMaxProcs ========================================================================================================
type diAutoMaxPro *string

// InitAutoMaxProcs - Automatically set GOMAXPROCS to match Linux container CPU quota
func InitAutoMaxProcs(log logger.Logger) (diAutoMaxPro, func(), error) {
	undo, err := maxprocs.Set(maxprocs.Logger(func(s string, args ...interface{}) {
		log.Info(fmt.Sprintf(s, args))
	}))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		undo()
	}

	return nil, cleanup, nil
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

// Logger ==============================================================================================================
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

// Store ===============================================================================================================
// InitStore return db
func InitStore(ctx context.Context, log logger.Logger) (*db.Store, func(), error) {
	var st db.Store
	db, err := st.Use(ctx, log)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Store.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db, cleanup, nil
}

// MQ ==================================================================================================================
func InitMQ(ctx context.Context, log logger.Logger) (mq.MQ, func(), error) {
	viper.SetDefault("MQ_ENABLED", "false") // Enabled MQ-service

	if viper.GetBool("MQ_ENABLED") {
		var service mq.DataBus
		dataBus, err := service.Use(ctx, log)
		if err != nil {
			return nil, func() {}, err
		}

		cleanup := func() {
			if err := dataBus.Close(); err != nil {
				log.Error(err.Error())
			}
		}

		return dataBus, cleanup, nil
	}

	return nil, func() {}, nil
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(NewContext, InitAutoMaxProcs, InitLogger, traicing.NewTracer)

// FullService =========================================================================================================
var FullSet = wire.NewSet(
	DefaultSet,
	NewFullService,
	InitStore,
	InitSentry,
	InitMonitoring,
	InitProfiling,
	InitMQ,
	rpc.InitServer,
	rpc.InitClient,
)

func NewFullService(
	ctx context.Context,
	log logger.Logger,
	mq mq.MQ,
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	db *db.Store,
	//linkStore *link_store.LinkStore,
	pprofHTTP PprofEndpoint,
	autoMaxProcsOption diAutoMaxPro,
	serverRPC *rpc.RPCServer,
	clientRPC *grpc.ClientConn,
) (*Service, error) {
	return &Service{
		Ctx:    ctx,
		Log:    log,
		MQ:     mq,
		Tracer: tracer,
		// TracerClose: cleanup,
		Monitoring: monitoring,
		Sentry:     sentryHandler,
		DB:         db,
		//LinkStore:     linkStore,
		PprofEndpoint: pprofHTTP,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

func InitializeFullService() (*Service, func(), error) {
	panic(wire.Build(FullSet))
}
