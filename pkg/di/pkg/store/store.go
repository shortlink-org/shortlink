package store

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/observability/metrics"
	"github.com/shortlink-org/shortlink/pkg/db"
	error_di "github.com/shortlink-org/shortlink/pkg/di/pkg/error"
)

// New - return implementation of db
//
//nolint:ireturn // It's make by specification
func New(ctx context.Context, log logger.Logger, tracer trace.TracerProvider, monitor *metrics.Monitoring) (db.DB, error) {
	store, err := db.New(ctx, log, tracer, monitor.Metrics)
	if err != nil {
		return nil, &error_di.BaseError{Err: err}
	}

	return store, nil
}
