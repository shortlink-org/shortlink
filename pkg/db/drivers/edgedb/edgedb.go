package edgedb

import (
	"context"
	"fmt"

	"github.com/geldata/gel-go"
	"github.com/geldata/gel-go/gelcfg"
	"github.com/spf13/viper"
)

// Config - config
type Config struct {
	URI string
}

// Store implementation of db interface
type Store struct {
	client *gel.Client
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	// Connect to EdgeDB
	s.client, err = gel.CreateClientDSN(s.config.URI, gelcfg.Options{
		Branch: "shortlink",
	})
	if err != nil {
		return &StoreError{
			Op:      "CreateClientDSN",
			Err:     fmt.Errorf("%w: %w", ErrConnect, err),
			Details: "failed to connect to EdgeDB at " + s.config.URI,
		}
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
	err := s.client.Close()
	if err != nil {
		return &StoreError{
			Op:      "close",
			Err:     fmt.Errorf("%w: %w", ErrClose, err),
			Details: "failed to close EdgeDB connection",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_EDGEDB_URI", "edgedb://localhost:5656") // EdgeDB URI

	s.config = Config{
		URI: viper.GetString("STORE_EDGEDB_URI"),
	}
}
