package edgedb

import (
	"context"

	"github.com/edgedb/edgedb-go"
	"github.com/spf13/viper"
)

// Config - config
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *edgedb.Client
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	// Connect to EdgeDB
	s.client, err = edgedb.CreateClientDSN(ctx, s.config.URI, edgedb.Options{
		Database: "shortlink",
	})
	if err != nil {
		return err
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		_ = s.close()
	}()

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// close - close
func (s *Store) close() error {
	return s.client.Close()
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_EDGEDB_URI", "edgedb://localhost:5656") // EdgeDB URI

	s.config = Config{
		URI: viper.GetString("STORE_EDGEDB_URI"),
	}
}
