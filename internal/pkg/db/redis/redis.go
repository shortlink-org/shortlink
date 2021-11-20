package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// Config ...
type Config struct { // nolint unused
	URI string
}

// Store implementation of db interface
type Store struct { // nolint unused
	client *redis.Client
	config Config
}

// Init ...
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	s.setConfig()

	// Connect to Redis
	s.client = redis.NewClient(&redis.Options{
		Addr:     s.config.URI,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := s.client.Ping(ctx).Result(); err != nil {
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
	viper.SetDefault("STORE_REDIS_URI", "localhost:6379") // Redis URI

	s.config = Config{
		URI: viper.GetString("STORE_REDIS_URI"),
	}
}
