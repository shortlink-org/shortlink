package ram

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/db/options"
)

// Config ...
type Config struct { // nolint unused
	mode int // Type write mode. single or batch
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
func (_ *Store) GetConn() interface{} {
	return nil
}

// Close ...
func (_ *Store) Close() error {
	return nil
}

// Migrate ...
func (_ *Store) migrate() error { // nolint unused
	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode write to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
