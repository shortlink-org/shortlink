package mysql

import (
	"database/sql"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"
)

// Config - configuration
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *sql.DB

	tracer  trace.TracerProvider
	metrics *metric.MeterProvider

	config Config
}
