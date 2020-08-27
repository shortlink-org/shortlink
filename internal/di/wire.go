//go:generate wire
//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"fmt"
	"net"
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
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/internal/traicing"
)

// Service - heplers
type Service struct {
	Log    logger.Logger
	Tracer opentracing.Tracer
	// TracerClose func()
	Sentry        *sentryhttp.Handler
	DB            store.DB
	MQ            mq.MQ
	ServerRPC     *RPCServer
	ClientRPC     *grpc.ClientConn
	Monitoring    *http.ServeMux
	PprofEndpoint PprofEndpoint
}

type PprofEndpoint *http.ServeMux

type diAutoMaxPro *string

type RPCServer struct {
	Run      func()
	Server   *grpc.Server
	Endpoint string
}

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

// InitStore return store
func InitStore(ctx context.Context, log logger.Logger) (store.DB, func(), error) {
	var st store.Store
	db, err := st.Use(ctx, log)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db, cleanup, nil
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

func InitTracer(ctx context.Context, log logger.Logger) (opentracing.Tracer, func(), error) {
	viper.SetDefault("TRACER_SERVICE_NAME", "ShortLink") // Service Name
	viper.SetDefault("TRACER_URI", "localhost:6831")     // Tracing addr:host

	config := traicing.Config{
		ServiceName: viper.GetString("TRACER_SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	tracer, tracerClose, err := traicing.Init(config)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := tracerClose.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return tracer, cleanup, nil
}

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
	commonMux.Handle("/metrics", sentryHandler.Handle(promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	})))

	// Expose a liveness check on /live
	commonMux.HandleFunc("/live", sentryHandler.HandleFunc(health.LiveEndpoint))

	// Expose a readiness check on /ready
	commonMux.HandleFunc("/ready", sentryHandler.HandleFunc(health.ReadyEndpoint))

	return commonMux
}

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

func InitSentry() (*sentryhttp.Handler, func(), error) {
	viper.SetDefault("SENTRY_DSN", "__DSN__") // key for sentry
	DSN := viper.GetString("SENTRY_DSN")

	if DSN != "" {
		return nil, func() {}, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("DSN"),
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

// runGRPCServer ...
func runGRPCServer(log logger.Logger) (*RPCServer, func(), error) {
	viper.SetDefault("GRPC_SERVER_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_SERVER_PORT")

	endpoint := fmt.Sprintf("localhost:%d", grpc_port)
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, nil, err
	}

	rpc := grpc.NewServer()

	r := &RPCServer{
		Server: rpc,
		Run: func() {
			go rpc.Serve(lis)
			log.Info("Run gRPC server", logger.Fields{"port": grpc_port})
		},
		Endpoint: endpoint,
	}

	cleanup := func() {
		rpc.GracefulStop()
	}

	return r, cleanup, err
}

// runGRPCClient - set up a connection to the server.
func runGRPCClient() (*grpc.ClientConn, func(), error) {
	viper.SetDefault("GRPC_CLIENT_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_CLIENT_PORT")

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", grpc_port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		conn.Close()
	}

	return conn, cleanup, nil
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(InitAutoMaxProcs, InitLogger, InitTracer)

// FullService =========================================================================================================
var FullSet = wire.NewSet(DefaultSet, NewFullService, InitStore, InitMonitoring, InitProfiling, InitMQ, InitSentry, runGRPCServer, runGRPCClient)

func NewFullService(log logger.Logger, mq mq.MQ, monitoring *http.ServeMux, tracer opentracing.Tracer, db store.DB, pprofHTTP PprofEndpoint, sentryHandler *sentryhttp.Handler, autoMaxProcsOption diAutoMaxPro, serverRPC *RPCServer, clientRPC *grpc.ClientConn) (*Service, error) {
	return &Service{
		Log:    log,
		MQ:     mq,
		Tracer: tracer,
		// TracerClose: cleanup,
		Monitoring:    monitoring,
		DB:            db,
		PprofEndpoint: pprofHTTP,
		Sentry:        sentryHandler,
		ClientRPC:     clientRPC,
		ServerRPC:     serverRPC,
	}, nil
}

func InitializeFullService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(FullSet))
}

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(DefaultSet, NewLoggerService, InitMQ)

func NewLoggerService(log logger.Logger, mq mq.MQ, autoMaxProcsOption diAutoMaxPro) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq,
	}, nil
}

func InitializeLoggerService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(LoggerSet))
}

// BotService ==========================================================================================================
var BotSet = wire.NewSet(DefaultSet, NewBotService, InitMQ)

func NewBotService(log logger.Logger, mq mq.MQ, autoMaxProcsOption diAutoMaxPro) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq,
	}, nil
}

func InitializeBotService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(BotSet))
}

// MetadataService =====================================================================================================
var MetadataSet = wire.NewSet(DefaultSet, NewMetadataService, runGRPCServer)

func NewMetadataService(log logger.Logger, autoMaxProcsOption diAutoMaxPro, serverRPC *RPCServer) (*Service, error) {
	return &Service{
		Log:       log,
		ServerRPC: serverRPC,
	}, nil
}

func InitializeMetadataService(ctx context.Context) (*Service, func(), error) {
	panic(wire.Build(MetadataSet))
}
