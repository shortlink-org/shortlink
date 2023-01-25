package redis

import (
	"context"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	Host     []string
	Username string
	Password string
}

// Store implementation of db interface
type Store struct {
	client redis.UniversalClient
	config Config
}

// Init ...
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	s.setConfig()

	// Connect to Redis
	s.client = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    s.config.Host,
		Username: s.config.Username,
		Password: s.config.Password,
		DB:       0, // use default DB
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(s.client); err != nil {
		return err
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(s.client); err != nil {
		return err
	}

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
func (s *Store) migrate() error {
	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_REDIS_URI", "localhost:6379") // Redis Hosts
	viper.SetDefault("STORE_REDIS_USERNAME", "")          // Redis Username
	viper.SetDefault("STORE_REDIS_PASSWORD", "")          // Redis Password

	s.config = Config{
		Host:     viper.GetStringSlice("STORE_REDIS_URI"),
		Username: viper.GetString("STORE_REDIS_USERNAME"),
		Password: viper.GetString("STORE_REDIS_PASSWORD"),
	}
}
