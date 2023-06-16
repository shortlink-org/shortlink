package cockroachdb

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	host string
	port int
}

// Store implementation of db interface
type Store struct {
	client *pgx.Conn
	config *pgx.ConnConfig
}

// Init ...
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

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (s *Store) Close() error {
	err := s.client.Close(context.Background())
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
