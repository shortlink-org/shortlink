package ram

import (
	"context"
	"sync"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/batch"
	"github.com/batazor/shortlink/internal/db/options"
)

// Config ...
type Config struct { // nolint unused
	mode int
	job  *batch.Config
}

// Store implementation of db interface
type Store struct { // nolint unused
	// sync.Map solver problem with cache contention
	links sync.Map

	config Config
}

// Init ...
func (s *Store) Init(ctx context.Context) error { // nolint unparam
	// Set configuration
	s.setConfig()

	return nil
}

// Close ...
func (ram *Store) Close() error {
	return nil
}

// Migrate ...
func (ram *Store) migrate() error { // nolint unused
	return nil
}

// setConfig - set configuration
func (ram *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_RAM_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db

	ram.config = Config{
		mode: viper.GetInt("STORE_RAM_MODE_WRITE"),
	}
}
