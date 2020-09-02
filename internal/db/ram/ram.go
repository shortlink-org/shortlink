package ram

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/db/options"
)

// Config ...
type Config struct { // nolint unused
	mode int
}

// Store implementation of db interface
type Store struct {
	config Config
}

// Init ...
func (s *Store) Init(_ context.Context) error {
	// Set configuration
	s.setConfig()

	return nil
}

// GetConn ...
func (s *Store) GetConn() interface{} {
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
