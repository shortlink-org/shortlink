package badger

import (
	"context"

	"github.com/dgraph-io/badger/v4"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Path string
}

// Store implementation of db interface
type Store struct {
	client *badger.DB
	config Config
}

// Init ...
func (b *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	b.setConfig()

	b.client, err = badger.Open(badger.DefaultOptions(b.config.Path))
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
func (b *Store) Close() error {
	return b.client.Close()
}

// setConfig - set configuration
func (b *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_BADGER_PATH", "/tmp/links.badger") // Badger path to file

	b.config = Config{
		Path: viper.GetString("STORE_BADGER_PATH"),
	}
}
