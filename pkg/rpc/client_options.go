package rpc

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
	grpc_logger "github.com/shortlink-org/shortlink/pkg/rpc/middleware/logger"
	session_interceptor "github.com/shortlink-org/shortlink/pkg/rpc/middleware/session"
)

type Option func(*Client)

// Apply a batch of options
func (c *Client) apply(options ...Option) {
	for _, option := range options {
		option(c)
	}
}

// WithTimeout sets a unary timeout interceptor
func WithTimeout() Option {
	viper.SetDefault("GRPC_CLIENT_TIMEOUT", "10s") // Set timeout for gRPC-Client
	timeoutClient := viper.GetDuration("GRPC_CLIENT_TIMEOUT")

	return func(c *Client) {
		c.interceptorUnaryClientList = append(
			c.interceptorUnaryClientList,
			timeout.UnaryClientInterceptor(timeoutClient),
		)
	}
}

// WithLogger adds unary & stream logging interceptors
func WithLogger(log logger.Logger) Option {
	viper.SetDefault("GRPC_CLIENT_LOGGER_ENABLED", true) // Enable logging for gRPC-Client
	isEnableLogger := viper.GetBool("GRPC_CLIENT_LOGGER_ENABLED")

	if isEnableLogger == false {
		return func(_ *Client) {}
	}

	return func(c *Client) {
		c.interceptorUnaryClientList = append(c.interceptorUnaryClientList, grpc_logger.UnaryClientInterceptor(log))
		c.interceptorStreamClientList = append(c.interceptorStreamClientList, grpc_logger.StreamClientInterceptor(log))
	}
}

// WithTracer wires up otel handler
func WithTracer(tracer trace.TracerProvider, prom *metrics.Monitoring) Option {
	return func(c *Client) {
		if tracer == nil || prom == nil {
			return
		}

		c.optionsNewClient = append(c.optionsNewClient, grpc.WithStatsHandler(
			otelgrpc.NewClientHandler(
				otelgrpc.WithTracerProvider(tracer),
				otelgrpc.WithMeterProvider(prom.Metrics),
				otelgrpc.WithMessageEvents(otelgrpc.ReceivedEvents, otelgrpc.SentEvents),
			),
		))
	}
}

// WithMetrics registers Prom metrics + interceptors
func WithMetrics(prom *metrics.Monitoring) Option {
	return func(c *Client) {
		if prom == nil {
			return
		}

		clientMetrics := grpc_prometheus.NewClientMetrics(
			grpc_prometheus.WithClientHandlingTimeHistogram(
				grpc_prometheus.WithHistogramBuckets([]float64{
					0.001, 0.01, 0.1, 0.3, 0.6, 1, 3, 6, 9, 20, 30, 60, 90, 120,
				}),
			),
		)
		exemplarFromCtx := grpc_prometheus.WithExemplarFromContext(exemplarFromContext)

		c.interceptorUnaryClientList = append(c.interceptorUnaryClientList, clientMetrics.UnaryClientInterceptor(exemplarFromCtx))
		c.interceptorStreamClientList = append(c.interceptorStreamClientList, clientMetrics.StreamClientInterceptor(exemplarFromCtx))

		defer func() {
			// ignore duplicate-registration panic
			_ = recover()
		}()
		prom.Prometheus.MustRegister(clientMetrics)
	}
}

// WithSession adds session interceptors with optional ignore rules
func WithSession() Option {
	return func(c *Client) {
		c.interceptorUnaryClientList = append(c.interceptorUnaryClientList, session_interceptor.SessionUnaryClientInterceptor())
		c.interceptorStreamClientList = append(c.interceptorStreamClientList, session_interceptor.SessionStreamClientInterceptor())
	}
}
