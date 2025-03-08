package etcd

import (
	"context"
	"strings"
	"time"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// Config - config
type Config struct {
	URI         []string
	DialTimeout time.Duration
}

// Store implementation of db interface
type Store struct {
	client *clientv3.Client
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	s.setConfig()

	// Connect to ETCD
	var err error

	s.client, err = clientv3.New(clientv3.Config{
		Endpoints:   s.config.URI,
		DialTimeout: s.config.DialTimeout,
	})
	if err != nil {
		return &StoreError{
			Op:      "init",
			Err:     ErrClientConnection,
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

// close - close
func (s *Store) close() error {
	if err := s.client.Close(); err != nil {
		return &StoreError{
			Op:      "close",
			Err:     err,
			Details: "failed to close etcd client",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_ETCD_URI", "localhost:2379") // ETCD URI
	viper.SetDefault("STORE_ETCD_TIMEOUT", "5s")         // ETCD timeout

	etcd := strings.Split(viper.GetString("STORE_ETCD_URI"), ",")

	s.config = Config{
		URI:         etcd,
		DialTimeout: viper.GetDuration("STORE_ETCD_TIMEOUT"),
	}
}
