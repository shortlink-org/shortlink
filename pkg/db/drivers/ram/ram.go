package ram

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/db/options"
)

// Config - config
type Config struct {
	mode int // Type write mode. single or batch
}

// Store implementation of db interface
type Store struct {
	config Config
}

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	// Set configuration
	s.setConfig()

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		// Nothing to do
	}()

	return nil
}

// GetConn - get connect
func (*Store) GetConn() any {
	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MODE_WRITE", options.MODE_SINGLE_WRITE) // mode writes to db. Select: 0 (MODE_SINGLE_WRITE), 1 (MODE_BATCH_WRITE)

	s.config = Config{
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
