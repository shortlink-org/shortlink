package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq" // need for init PostgreSQL interface
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
)

// Init ...
func (p *Store) Init(ctx context.Context) error {
	var err error
	p.tracer = Tracer{}

	// Set configuration
	p.config, err = getConfig(&p.tracer)
	if err != nil {
		return err
	}

	// Apply migration
	err = p.migrate()
	if err != nil {
		return err
	}

	// Connect to Postgres
	p.client, err = pgxpool.NewWithConfig(ctx, p.config.poolConfig)
	if err != nil {
		return err
	}

	// Check connect
	err = p.client.Ping(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (p *Store) Close() error {
	p.client.Close()
	return nil
}

// setConfig - set configuration
func getConfig(tracer *Tracer) (*Config, error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", "postgres", "shortlink", "shortlink")

	viper.AutomaticEnv()
	viper.SetDefault("STORE_POSTGRES_URI", dbinfo)                  // Postgres URI
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db
	viper.SetDefault("STORE_POSTGRES_POOL_CONFIG", "")              // config for connect to db

	// Parse config for connect
	cnf, err := pgx.ParseConfig(viper.GetString("STORE_POSTGRES_URI"))
	if err != nil {
		return nil, err
	}

	// Instrument the pgxpool config with OpenTelemetry.
	cnf.Tracer = tracer

	// Create pool config
	cnfPool, err := pgxpool.ParseConfig(viper.GetString("STORE_POSTGRES_POOL_CONFIG"))
	if err != nil {
		return nil, err
	}

	cnfPool.ConnConfig = cnf

	return &Config{
		configConnect: cnf,
		mode:          viper.GetInt("STORE_MODE_WRITE"),
		poolConfig:    cnfPool,
	}, nil
}
