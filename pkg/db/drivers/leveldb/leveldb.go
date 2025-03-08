package leveldb

import (
	"context"

	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb"
)

// Config - config
type Config struct {
	Path string
}

// Store implementation of db interface
type Store struct {
	client *leveldb.DB
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	s.client, err = leveldb.OpenFile(s.config.Path, nil)
	if err != nil {
		return &StoreError{
			Op:      "init",
			Err:     ErrDatabaseOpen,
			Details: err.Error(),
		}
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if err := s.close(); err != nil {
			// We can't return the error here since we're in a goroutine,
			// but in a real application you might want to log this
			_ = err
		}
	}()

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// Close - close
func (s *Store) close() error {
	if err := s.client.Close(); err != nil {
		return &StoreError{
			Op:      "close",
			Err:     err,
			Details: "failed to close leveldb database",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_LEVELDB_PATH", "/tmp/links.db") // LevelDB path to file

	s.config = Config{
		Path: viper.GetString("STORE_LEVELDB_PATH"),
	}
}
