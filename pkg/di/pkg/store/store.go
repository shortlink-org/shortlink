package store

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
)

// New - return implementation of db
func New(ctx context.Context, log logger.Logger, tracer trace.TracerProvider, monitor *monitoring.Monitoring) (db.DB, error) {
	store, err := db.New(ctx, log, tracer, monitor.Metrics)
	if err != nil {
		return nil, err
	}

	return store, nil
}
