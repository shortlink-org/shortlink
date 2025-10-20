package permission

import (
	"context"

	"github.com/authzed/authzed-go/v1"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/go-sdk/auth"
	rpc "github.com/shortlink-org/go-sdk/grpc"
	"github.com/shortlink-org/go-sdk/logger"

	"github.com/shortlink-org/go-sdk/observability/metrics"
	error_di "github.com/shortlink-org/shortlink/pkg/di/pkg/error"
)

func New(_ context.Context, log logger.Logger, tracer trace.TracerProvider, monitor *metrics.Monitoring) (*authzed.Client, error) {
	// Initialize gRPC Client's interceptor.
	opts := []rpc.Option{
		rpc.WithSession(),
		rpc.WithMetrics(monitor.Prometheus),
		rpc.WithTracer(tracer, monitor.Prometheus, monitor.Metrics),
		rpc.WithTimeout(),
		rpc.WithLogger(log),
	}

	permission, err := auth.New(opts...)
	if err != nil {
		return nil, &error_di.BaseError{Err: err}
	}

	return permission, nil
}
