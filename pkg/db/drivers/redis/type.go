package redis

import (
	"github.com/redis/rueidis"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"
)

// Config - config
type Config struct {
	Username string
	Password string
	Host     []string
}

// Store implementation of db interface
type Store struct {
	client rueidis.Client

	tracer  trace.TracerProvider
	metrics *metric.MeterProvider

	config Config
}
