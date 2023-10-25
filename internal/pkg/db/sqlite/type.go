package sqlite

import (
	"database/sql"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"
)

// Config - config
type Config struct {
	Path string
}

// Store implementation of db interface
type Store struct {
	client *sql.DB
	config Config

	tracer  trace.TracerProvider
	metrics *metric.MeterProvider
}
