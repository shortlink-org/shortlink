package cockroachdb

import (
	"context"

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
		return err
	}

	s.client, err = pgx.ConnectConfig(ctx, s.config)
	if err != nil {
		return err
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
		return err
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
		return err
	}

	s.config.RuntimeParams["application_name"] = viper.GetString("SERVICE_NAME")

	return nil
}
