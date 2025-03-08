package postgres

import (
	"context"
	"fmt"

	"github.com/exaring/otelpgx"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/pkg/db/options"
)

// New return new instance of Store
func New(tracer trace.TracerProvider, metrics *metric.MeterProvider) *Store {
	return &Store{
		tracer: Tracer{
			TracerProvider: tracer,
		},
		metrics: metrics,
	}
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.config, err = getConfig(&s.tracer)
	if err != nil {
		return &StoreError{
			Op:      "init",
			Err:     err,
			Details: "failed to get postgres connection config",
		}
	}

	// Connect to Postgres
	s.client, err = pgxpool.NewWithConfig(ctx, s.config.config)
	if err != nil {
		return &StoreError{
			Op:      "init",
			Err:     err,
			Details: "failed to open the database",
		}
	}

	// Check connecting
	err = s.client.Ping(ctx)
	if err != nil {
		s.client.Close()

		return &PingConnectionError{err}
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		s.client.Close()
	}()

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// setConfig - set configuration
func getConfig(tracer *Tracer) (*Config, error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", "postgres", "shortlink", "shortlink")

	viper.AutomaticEnv()
	viper.SetDefault("STORE_POSTGRES_URI", dbinfo)                  // Postgres URI
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db

	// Create pool config
	cnfPool, err := pgxpool.ParseConfig(viper.GetString("STORE_POSTGRES_URI"))
	if err != nil {
		return nil, &StoreError{
			Op:      "ParseConfig",
			Err:     err,
			Details: "failed to parse postgres connection config",
		}
	}

	// Instrument the pgxpool config with OpenTelemetry.
	params := []otelpgx.Option{
		otelpgx.WithIncludeQueryParameters(),
	}
	if tracer.TracerProvider != nil {
		params = append(params, otelpgx.WithTracerProvider(tracer))
	}

	cnfPool.ConnConfig.Tracer = otelpgx.NewTracer(params...)

	return &Config{
		config: cnfPool,
		mode:   viper.GetInt("STORE_MODE_WRITE"),
	}, nil
}
