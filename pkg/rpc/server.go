package rpc

import (
	"context"
	"fmt"
	"net"
	"runtime/debug"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
	grpc_logger "github.com/shortlink-org/shortlink/pkg/rpc/middleware/logger"
	pprof_interceptor "github.com/shortlink-org/shortlink/pkg/rpc/middleware/pprof"
	session_interceptor "github.com/shortlink-org/shortlink/pkg/rpc/middleware/session"
)

type Server struct {
	Run      func()
	Server   *grpc.Server
	Endpoint string
}

type server struct {
	interceptorStreamServerList []grpc.StreamServerInterceptor
	interceptorUnaryServerList  []grpc.UnaryServerInterceptor
	optionsNewServer            []grpc.ServerOption

	port int
	host string

	log           logger.Logger
	serverMetrics *grpc_prometheus.ServerMetrics
}

// InitServer - initialize gRPC server
func InitServer(ctx context.Context, log logger.Logger, tracer trace.TracerProvider, monitor *metrics.Monitoring) (*Server, error) {
	viper.SetDefault("GRPC_SERVER_ENABLED", true) // gRPC grpServer enable

	if !viper.GetBool("GRPC_SERVER_ENABLED") {
		//nolint:nilnil // it's correct logic
		return nil, nil
	}

	config, err := setServerConfig(log, tracer, monitor) //nolint:contextcheck // false positive
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s:%d", config.host, config.port)

	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	// Initialize the gRPC grpServer.
	rpc := grpc.NewServer(config.optionsNewServer...)

	grpServer := &Server{
		Server: rpc,
		Run: func() {
			// Register reflection service on gRPC grpServer.
			reflection.Register(rpc)

			// After all your registrations, make sure all of the Prometheus metrics are initialized.
			config.serverMetrics.InitializeMetrics(rpc)

			log.Info("Run gRPC grpServer", field.Fields{"port": config.port, "host": config.host})
			err = rpc.Serve(lis)
			if err != nil {
				log.Error(err.Error())
			}
		},
		Endpoint: endpoint,
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		log.Info("Shutdown gRPC grpServer")
		rpc.GracefulStop()
	}()

	if err != nil {
		return nil, fmt.Errorf("failed to serve: %w", err)
	}

	return grpServer, nil
}

// setConfig - set configuration
func setServerConfig(log logger.Logger, tracer trace.TracerProvider, monitor *metrics.Monitoring) (*server, error) {
	viper.SetDefault("GRPC_SERVER_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_SERVER_PORT")

	viper.SetDefault("GRPC_SERVER_HOST", "0.0.0.0") // gRPC host
	grpc_host := viper.GetString("GRPC_SERVER_HOST")

	config := &server{
		port: grpc_port,
		host: grpc_host,

		log: log,
	}

	config.WithLogger(log)
	config.WithTracer(tracer)
	config.withSession()
	config.withPprofLabels()

	if monitor != nil {
		config.WithMetrics(monitor)
		config.WithRecovery(monitor)
	}

	config.optionsNewServer = append(config.optionsNewServer,
		// Initialize your gRPC server's interceptor.
		grpc.ChainUnaryInterceptor(config.interceptorUnaryServerList...),
		grpc.ChainStreamInterceptor(config.interceptorStreamServerList...),
	)

	// NOTE: made after initialize your gRPC server's interceptor.
	err := config.WithTLS()
	if err != nil {
		return nil, err
	}

	return config, nil
}

// WithMetrics - setup metrics.
func (s *server) WithMetrics(monitor *metrics.Monitoring) {
	s.serverMetrics = grpc_prometheus.NewServerMetrics(
		grpc_prometheus.WithServerHandlingTimeHistogram(
			grpc_prometheus.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)
	monitor.Prometheus.MustRegister(s.serverMetrics)

	exemplarFromCtx := grpc_prometheus.WithExemplarFromContext(exemplarFromContext)

	s.interceptorUnaryServerList = append(s.interceptorUnaryServerList, s.serverMetrics.UnaryServerInterceptor(exemplarFromCtx))
	s.interceptorStreamServerList = append(s.interceptorStreamServerList, s.serverMetrics.StreamServerInterceptor(exemplarFromCtx))
}

// WithTracer - setup tracing
func (s *server) WithTracer(tracer trace.TracerProvider) {
	if tracer == nil {
		return
	}

	s.optionsNewServer = append(s.optionsNewServer, grpc.StatsHandler(
		otelgrpc.NewServerHandler(otelgrpc.WithTracerProvider(tracer))),
	)
}

// WithRecovery - setup recovery
func (s *server) WithRecovery(monitor *metrics.Monitoring) {
	// Setup metric for panic recoveries.
	panicsTotal := promauto.With(monitor.Prometheus).NewCounter(prometheus.CounterOpts{
		Name: "grpc_req_panics_recovered_total",
		Help: "Total number of gRPC requests recovered from internal panic.",
	})
	grpcPanicRecoveryHandler := func(in any) error {
		panicsTotal.Inc()
		s.log.Error("recovered from panic", field.Fields{
			"panic": in,
			"stack": debug.Stack(),
		})

		return status.Errorf(codes.Internal, "%s", in)
	}

	recoveryHandler := grpc_recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)

	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g., logging) can operate in the recovered state instead of being directly affected by any panic
	s.interceptorUnaryServerList = append(s.interceptorUnaryServerList, grpc_recovery.UnaryServerInterceptor(recoveryHandler))

	// Create a server. Recovery handlers should typically be last in the chain so that other middleware
	// (e.g., logging) can operate in the recovered state instead of being directly affected by any panic
	s.interceptorStreamServerList = append(s.interceptorStreamServerList, grpc_recovery.StreamServerInterceptor(recoveryHandler))
}

// WithLogger - setup logger
func (s *server) WithLogger(log logger.Logger) {
	viper.SetDefault("GRPC_SERVER_LOGGER_ENABLED", true) // Enable logging for gRPC-Client
	isEnableLogger := viper.GetBool("GRPC_SERVER_LOGGER_ENABLED")

	if isEnableLogger {
		s.interceptorStreamServerList = append(s.interceptorStreamServerList, grpc_logger.StreamServerInterceptor(log))
		s.interceptorUnaryServerList = append(s.interceptorUnaryServerList, grpc_logger.UnaryServerInterceptor(log))
	}
}

// WithTLS - setup TLS
func (s *server) WithTLS() error {
	viper.SetDefault("GRPC_SERVER_TLS_ENABLED", false) // gRPC tls
	isEnableTLS := viper.GetBool("GRPC_SERVER_TLS_ENABLED")

	viper.SetDefault("GRPC_SERVER_CERT_PATH", "ops/cert/shortlink-server.pem") // gRPC server cert
	certFile := viper.GetString("GRPC_SERVER_CERT_PATH")

	viper.SetDefault("GRPC_SERVER_KEY_PATH", "ops/cert/shortlink-server-key.pem") // gRPC server key
	keyFile := viper.GetString("GRPC_SERVER_KEY_PATH")

	if isEnableTLS {
		creds, errTLSFromFile := credentials.NewServerTLSFromFile(certFile, keyFile)
		if errTLSFromFile != nil {
			return fmt.Errorf("failed to setup TLS: %w", errTLSFromFile)
		}

		s.optionsNewServer = append(s.optionsNewServer, grpc.Creds(creds))
	}

	return nil
}

// withSession - setup session
func (s *server) withSession() {
	s.interceptorUnaryServerList = append(s.interceptorUnaryServerList, session_interceptor.SessionUnaryServerInterceptor())
	s.interceptorStreamServerList = append(s.interceptorStreamServerList, session_interceptor.SessionStreamServerInterceptor())
}

// withPprofLabels - setup pprof labels
func (s *server) withPprofLabels() {
	s.interceptorUnaryServerList = append(s.interceptorUnaryServerList, pprof_interceptor.UnaryServerInterceptor())
	s.interceptorStreamServerList = append(s.interceptorStreamServerList, pprof_interceptor.StreamServerInterceptor())
}
