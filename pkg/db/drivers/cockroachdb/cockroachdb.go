package cockroachdb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

// Store implementation of db interface
type Store struct {
	client *pgx.Conn
	config *pgx.ConnConfig
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	err := s.setConfig()
	if err != nil {
		return &StoreError{
			Op:      "setConfig",
			Err:     err,
			Details: "failed to set CockroachDB configuration",
		}
	}

	s.client, err = pgx.ConnectConfig(ctx, s.config)
	if err != nil {
		return &StoreError{
			Op:      "ConnectConfig",
			Err:     fmt.Errorf("%w: %w", ErrCockroachConnect, err),
			Details: "failed to connect to CockroachDB with provided config",
		}
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		//nolint:errcheck // ignore
		_ = s.close(ctx)
	}()

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// close - close
func (s *Store) close(ctx context.Context) error {
	err := s.client.Close(ctx)
	if err != nil {
		return &StoreError{
			Op:      "close",
			Err:     err,
			Details: "failed to close CockroachDB connection",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() error {
	var err error

	viper.AutomaticEnv()
	viper.SetDefault("STORE_COCKROACHDB_URI", "postgresql://root@localhost:26257?sslmode=disable") // CockroachDB URI

	s.config, err = pgx.ParseConfig(viper.GetString("STORE_COCKROACHDB_URI"))
	if err != nil {
		return &StoreError{
			Op:      "ParseConfig",
			Err:     fmt.Errorf("%w: %w", ErrCockroachConfig, err),
			Details: "failed to parse CockroachDB URI from environment",
		}
	}

	s.config.RuntimeParams["application_name"] = viper.GetString("SERVICE_NAME")

	return nil
}
