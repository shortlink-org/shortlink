package permission

import (
	"context"

	"github.com/authzed/authzed-go/v1"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/auth"
	error_di "github.com/shortlink-org/shortlink/pkg/di/pkg/error"
	"github.com/shortlink-org/shortlink/pkg/observability/metrics"
	"github.com/shortlink-org/shortlink/pkg/rpc"
)

func New(_ context.Context, log logger.Logger, tracer trace.TracerProvider, monitor *metrics.Monitoring) (*authzed.Client, error) {
	// Initialize gRPC Client's interceptor.
	opts := []rpc.Option{
		rpc.WithSession(),
		rpc.WithMetrics(monitor),
		rpc.WithTracer(tracer, monitor),
		rpc.WithTimeout(),
		rpc.WithLogger(log),
	}

	permission, err := auth.New(opts...)
	if err != nil {
		return nil, &error_di.BaseError{Err: err}
	}

	return permission, nil
}
