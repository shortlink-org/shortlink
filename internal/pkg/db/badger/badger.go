package badger

import (
	"context"

	"github.com/dgraph-io/badger/v4"
	"github.com/spf13/viper"
)

// Config - config
type Config struct {
	Path string
}

// Store implementation of db interface
type Store struct {
	client *badger.DB
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	s.client, err = badger.Open(badger.DefaultOptions(s.config.Path))
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

// Close - close
func (s *Store) close() error {
	return s.client.Close()
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_BADGER_PATH", "/tmp/links.badger") // Badger path to file

	s.config = Config{
		Path: viper.GetString("STORE_BADGER_PATH"),
	}
}
