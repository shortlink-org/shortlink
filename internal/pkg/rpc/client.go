package rpc

import (
	"fmt"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	grpc_logger "github.com/shortlink-org/shortlink/internal/pkg/rpc/middleware/logger"
	session_interceptor "github.com/shortlink-org/shortlink/internal/pkg/rpc/middleware/session"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

type client struct {
	interceptorUnaryClientList  []grpc.UnaryClientInterceptor
	interceptorStreamClientList []grpc.StreamClientInterceptor
	optionsNewClient            []grpc.DialOption

	port int
	host string
}

// InitClient - set up a connection to the server.
func InitClient(log logger.Logger, tracer trace.TracerProvider, monitoring *monitoring.Monitoring) (*grpc.ClientConn, func(), error) {
	config, err := SetClientConfig(tracer, monitoring, log)
	if err != nil {
		return nil, nil, err
	}

	// Set up a connection to the server peer
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", config.host, config.port),
		config.optionsNewClient...,
	)
	if err != nil {
		return nil, nil, err
	}

	log.Info("Run gRPC client", field.Fields{"port": config.port, "host": config.host})

	cleanup := func() {
		_ = conn.Close()
	}

	return conn, cleanup, nil
}

// setConfig - set configuration
func SetClientConfig(tracer trace.TracerProvider, monitoring *monitoring.Monitoring, log logger.Logger) (*client, error) {
	viper.SetDefault("GRPC_CLIENT_PORT", "50051") // gRPC port
	grpc_port := viper.GetInt("GRPC_CLIENT_PORT")

	viper.SetDefault("GRPC_CLIENT_HOST", "0.0.0.0") // gRPC host
	grpc_host := viper.GetString("GRPC_CLIENT_HOST")

	config := &client{
		port: grpc_port,
		host: grpc_host,
	}

	// Initialize gRPC client's interceptor.
	config.withSession()
	config.withMetrics(monitoring)
	config.withTracer(tracer)
	config.withLogger(log)
	config.withTimeout()

	// Initialize your gRPC server's interceptor.
	config.optionsNewClient = append(
		config.optionsNewClient,
		grpc.WithChainUnaryInterceptor(config.interceptorUnaryClientList...),
		grpc.WithChainStreamInterceptor(config.interceptorStreamClientList...),
	)

	// NOTE: made after initialize your gRPC client's interceptor.
	err := config.withTLS()
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GetOptions - return options for gRPC client.
func (c *client) GetOptions() []grpc.DialOption {
	return c.optionsNewClient
}

// withTimeout - setup timeout
func (c *client) withTimeout() {
	viper.SetDefault("GRPC_CLIENT_TIMEOUT", 10000) // Set timeout for gRPC-client
	timeoutClient := viper.GetDuration("GRPC_CLIENT_TIMEOUT")

	c.interceptorUnaryClientList = append(c.interceptorUnaryClientList, timeout.UnaryClientInterceptor(timeoutClient*time.Millisecond))
}

// withLogger - setup logger
func (c *client) withLogger(log logger.Logger) {
	viper.SetDefault("GRPC_CLIENT_LOGGER_ENABLED", true) // Enable logging for gRPC-client
	isEnableLogger := viper.GetBool("GRPC_CLIENT_LOGGER_ENABLED")

	if isEnableLogger {
		c.interceptorUnaryClientList = append(c.interceptorUnaryClientList, grpc_logger.UnaryClientInterceptor(log))
		c.interceptorStreamClientList = append(c.interceptorStreamClientList, grpc_logger.StreamClientInterceptor(log))
	}
}

// withTLS - setup TLS
func (c *client) withTLS() error {
	viper.SetDefault("GRPC_CLIENT_TLS_ENABLED", false) // gRPC TLS
	isEnableTLS := viper.GetBool("GRPC_CLIENT_TLS_ENABLED")

	viper.SetDefault("GRPC_CLIENT_CERT_PATH", "ops/cert/intermediate_ca.pem") // gRPC client cert
	certFile := viper.GetString("GRPC_CLIENT_CERT_PATH")

	if isEnableTLS {
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			return err
		}

		c.optionsNewClient = append(c.optionsNewClient, grpc.WithTransportCredentials(creds))
		return nil
	}

	c.optionsNewClient = append(c.optionsNewClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return nil
}

// withTracer - setup tracing
func (c *client) withTracer(tracer trace.TracerProvider) {
	if tracer == nil {
		return
	}

	c.interceptorUnaryClientList = append(c.interceptorUnaryClientList, otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(otel.GetTracerProvider())))
	c.interceptorStreamClientList = append(c.interceptorStreamClientList, otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(otel.GetTracerProvider())))
}

// withMetrics - setup metrics.
func (c *client) withMetrics(monitoring *monitoring.Monitoring) {
	if monitoring == nil {
		return
	}

	clientMetrics := grpc_prometheus.NewClientMetrics(
		grpc_prometheus.WithClientHandlingTimeHistogram(
			grpc_prometheus.WithHistogramBuckets([]float64{0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120}),
		),
	)

	c.interceptorUnaryClientList = append(c.interceptorUnaryClientList, clientMetrics.UnaryClientInterceptor(grpc_prometheus.WithExemplarFromContext(exemplarFromContext)))
	c.interceptorStreamClientList = append(c.interceptorStreamClientList, clientMetrics.StreamClientInterceptor(grpc_prometheus.WithExemplarFromContext(exemplarFromContext)))

	monitoring.Prometheus.MustRegister(clientMetrics)
}

// withSession - setup session
func (c *client) withSession() {
	c.interceptorUnaryClientList = append(c.interceptorUnaryClientList, session_interceptor.SessionUnaryInterceptor())
	c.interceptorStreamClientList = append(c.interceptorStreamClientList, session_interceptor.SessionStreamInterceptor())
}
