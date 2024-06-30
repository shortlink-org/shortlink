package sqlite

import (
	"context"
	"time"

	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/sdk/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
	"go.opentelemetry.io/otel/trace"
)

// New - create new instance of Store
func New(tracer trace.TracerProvider, metrics *metric.MeterProvider) *Store {
	return &Store{
		metrics: metrics,
		tracer:  tracer,
	}
}

// Init - init connection
func (s *Store) Init(ctx context.Context) error {
	const SET_MAX_OPEN_CONNS = 25
	const SET_MAX_IDLE_CONNS = 2

	var err error

	// Set configuration
	s.setConfig()

	options := []otelsql.Option{
		otelsql.WithAttributes(semconv.DBSystemSqlite),
		otelsql.WithDBName("SQLite"),
	}

	if s.metrics != nil {
		options = append(options, otelsql.WithMeterProvider(s.metrics))
	}

	if s.tracer != nil {
		options = append(options, otelsql.WithTracerProvider(s.tracer))
	}

	s.client, err = otelsql.Open("sqlite3", s.config.Path, options...)
	if err != nil {
		return err
	}

	s.client.SetMaxOpenConns(SET_MAX_OPEN_CONNS)
	s.client.SetMaxIdleConns(SET_MAX_IDLE_CONNS)
	s.client.SetConnMaxLifetime(time.Minute)

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		_ = s.close()
	}()

	return nil
}

// GetConn - return connection
func (s *Store) GetConn() any {
	return s.client
}

// close - close connection
func (s *Store) close() error {
	return s.client.Close()
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_SQLITE_PATH", "/tmp/links.sqlite") // SQLite URI

	s.config = Config{
		Path: viper.GetString("STORE_SQLITE_PATH"),
	}
}
