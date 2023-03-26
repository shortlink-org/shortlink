package edgedb

import (
	"context"

	"github.com/edgedb/edgedb-go"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *edgedb.Client
	config Config
}

// Init ...
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

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
	return s.client
}

// Close ...
func (s *Store) Close() error {
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
