package couchbase

import (
	"context"
	"fmt"

	"github.com/couchbase/gocb/v2"
	"github.com/spf13/viper"
)

type config struct {
	uri     string
	options gocb.ClusterOptions
}

// Store implementation of db interface
type Store struct {
	client *gocb.Cluster
	config *config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	err := s.setConfig()
	if err != nil {
		return &StoreError{
			Op:      "setConfig",
			Err:     err,
			Details: "failed to set configuration for Couchbase",
		}
	}

	s.client, err = gocb.Connect(s.config.uri, s.config.options)
	if err != nil {
		return &StoreError{
			Op:      "Connect",
			Err:     fmt.Errorf("%w: %w", ErrCouchbaseConnect, err),
			Details: "failed to connect to Couchbase cluster",
		}
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		//nolint:errcheck // ignore
		_ = s.client.Close(nil)
	}()

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() error {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_COUCHBASE_URI", "couchbase://localhost") // Couchbase URI (e.g. couchbase://localhost)

	s.config = &config{
		uri:     viper.GetString("STORE_COUCHBASE_URI"),
		options: gocb.ClusterOptions{},
	}

	return nil
}
