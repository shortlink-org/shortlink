package etcd

import (
	"context"
	"strings"
	"time"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Config ...
type Config struct { // nolint unused
	URI []string
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *clientv3.Client
	config Config
}

// Init ...
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	s.setConfig()

	// Connect to ETCD
	var err error
	s.client, err = clientv3.New(clientv3.Config{
		Endpoints:   s.config.URI,
		DialTimeout: 5 * time.Second,
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

// Migrate ...
func (s *Store) migrate() error { // nolint unused
	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_ETCD_URI", "localhost:2379") // ETCD URI

	etcd := strings.Split(viper.GetString("STORE_ETCD_URI"), ",")

	s.config = Config{
		URI: etcd,
	}
}
