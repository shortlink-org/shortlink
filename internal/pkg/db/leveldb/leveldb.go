package leveldb

import (
	"context"

	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb"
)

// Config ...
type Config struct {
	Path string
}

// Store implementation of db interface
type Store struct { // nolint:decorder
	client *leveldb.DB
	config Config
}

// Init ...
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	s.client, err = leveldb.OpenFile("/tmp/links.db", nil)
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
func (l *Store) Close() error {
	return l.client.Close()
}

// setConfig - set configuration
func (l *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_LEVELDB_PATH", "/tmp/links.db") // LevelDB path to file

	l.config = Config{
		Path: viper.GetString("STORE_LEVELDB_PATH"),
	}
}
