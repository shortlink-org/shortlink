package store

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/monitoring"
)

// New - return implementation of db
func New(ctx context.Context, log logger.Logger, tracer trace.TracerProvider, monitor *monitoring.Monitoring) (*db.Store, func(), error) {
	db, err := db.New(ctx, log, tracer, monitor.Metrics)
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
