package clickhouse

import (
	"context"

	"github.com/spf13/viper"
	"github.com/uptrace/go-clickhouse/ch"
	"github.com/uptrace/go-clickhouse/chdebug"
	"github.com/uptrace/go-clickhouse/chotel"
)

// Config - config
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *ch.DB
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	s.setConfig()

	// Connect to Clickhouse
	s.client = ch.Connect(ch.WithDSN(s.config.URI))
	s.client.AddQueryHook(chdebug.NewQueryHook(chdebug.WithVerbose(true)))
	s.client.AddQueryHook(chotel.NewQueryHook())

	if err := s.client.Ping(ctx); err != nil {
		return err
	}

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// Close - close
func (s *Store) Close() error {
	return s.client.Close()
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_CLICKHOUSE_URI", "clickhouse://localhost:9000/default?sslmode=disable") // Clickhouse URI

	s.config = Config{
		URI: viper.GetString("STORE_CLICKHOUSE_URI"),
	}
}
