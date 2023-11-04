package redis

import (
	"context"

	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidisotel"
	"github.com/spf13/viper"
)

// Config - config
type Config struct {
	Username string
	Password string
	Host     []string
}

// Store implementation of db interface
type Store struct {
	client rueidis.Client
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	// Connect to Redis
	s.client, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress: s.config.Host,
		Username:    s.config.Username,
		Password:    s.config.Password,
		SelectDB:    0, // use default DB
	})
	if err != nil {
		return err
	}

	// Enable tracing instrumentation.
	s.client = rueidisotel.WithClient(s.client)

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

// close - close connection
//
//nolint:unparam // ignore
func (s *Store) close() error {
	s.client.Close()
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
