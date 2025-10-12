package mysql

import (
	"context"
	"net/url"

	"github.com/XSAM/otelsql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/sdk/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
	"go.opentelemetry.io/otel/trace"
)

func New(tracer trace.TracerProvider, metrics *metric.MeterProvider) *Store {
	return &Store{
		tracer:  tracer,
		metrics: metrics,
	}
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	err := s.setConfig()
	if err != nil {
		return &StoreError{
			Op:      "setConfig",
			Err:     err,
			Details: "failed to set mysql configuration",
		}
	}

	options := []otelsql.Option{
		otelsql.WithTracerProvider(s.tracer),
		otelsql.WithMeterProvider(s.metrics),
		otelsql.WithSQLCommenter(true),
	}

	// Connect to MySQL
	if s.client, err = otelsql.Open("mysql", s.config.URI, options...); err != nil {
		return &StoreError{
			Op:      "init",
			Err:     ErrClientConnection,
			Details: err.Error(),
		}
	}

	// Check connection
	if errPing := s.client.Ping(); errPing != nil {
		_ = s.client.Close()

		return &PingConnectionError{
			Err: errPing,
		}
	}

	// Register DB stats to meter
	err = otelsql.RegisterDBStatsMetrics(s.client, otelsql.WithAttributes(
		semconv.DBSystemNameMySQL,
	))
	if err != nil {
		return &StoreError{
			Op:      "init",
			Err:     err,
			Details: "failed to register DB stats metrics",
		}
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if err := s.close(); err != nil {
			// We can't return the error here since we're in a goroutine,
			// but in a real application you might want to log this
			_ = err
		}
	}()

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// Close - close
func (s *Store) close() error {
	if err := s.client.Close(); err != nil {
		return &StoreError{
			Op:      "close",
			Err:     err,
			Details: "failed to close mysql connection",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() error {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MYSQL_URI", "shortlink:shortlink@(localhost:3306)/shortlink") // MySQL URI

	// parse uri
	uri, err := url.Parse(viper.GetString("STORE_MYSQL_URI"))
	if err != nil {
		return &StoreError{
			Op:      "setConfig",
			Err:     ErrInvalidDSN,
			Details: "parsing MySQL URI from environment variable",
		}
	}

	values := uri.Query()
	values.Add("parseTime", "true")

	uri.RawQuery = values.Encode()

	s.config = Config{
		URI: uri.String(),
	}

	return nil
}
