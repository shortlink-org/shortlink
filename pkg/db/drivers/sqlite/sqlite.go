package sqlite

import (
	"context"
	"time"

	_ "github.com/mattn/go-sqlite3" // Init SQLite-driver
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	"go.opentelemetry.io/otel/sdk/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
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

	// Set configuration
	s.setConfig()

	if s.config.Path == "" {
		return &StoreError{
			Op:      "init",
			Err:     ErrInvalidPath,
			Details: "sqlite path configuration is empty",
		}
	}

	options := []otelsql.Option{
		otelsql.WithAttributes(semconv.DBSystemNameSQLite),
		otelsql.WithDBName("SQLite"),
	}

	if s.metrics != nil {
		options = append(options, otelsql.WithMeterProvider(s.metrics))
	}

	if s.tracer != nil {
		options = append(options, otelsql.WithTracerProvider(s.tracer))
	}

	client, err := otelsql.Open("sqlite3", s.config.Path, options...)
	if err != nil {
		return &StoreError{
			Op:      "init",
			Err:     ErrClientConnection,
			Details: err.Error(),
		}
	}

	s.client = client

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
	err := s.client.Close()
	if err != nil {
		return &StoreError{
			Op:      "close",
			Err:     err,
			Details: "failed to close connection",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_SQLITE_PATH", "/tmp/links.sqlite") // SQLite URI

	s.config = Config{
		Path: viper.GetString("STORE_SQLITE_PATH"),
	}
}
