package mongo

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/johejo/golang-migrate-extra/source/file"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"

	storeOptions "github.com/shortlink-org/shortlink/pkg/db/options"
)

// Init - initialize
func (s *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	s.setConfig()

	// Connect to MongoDB
	opts := options.Client().
		ApplyURI(s.config.URI).
		SetCompressors([]string{"snappy", "zlib", "zstd"}).
		SetAppName(viper.GetString("SERVICE_NAME")).
		// TODO: wait new version
		// link: https://github.com/open-telemetry/opentelemetry-go-contrib/issues/6419
		// SetMonitor(otelmongo.NewMonitor()).
		SetRetryReads(true).
		SetRetryWrites(true)

	s.client, err = mongo.Connect(opts)
	if err != nil {
		return &StoreError{
			Op:      "init",
			Err:     ErrClientConnection,
			Details: err.Error(),
		}
	}

	// Check connecting
	err = s.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return &PingConnectionError{
			Err: err,
		}
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if err := s.close(ctx); err != nil {
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
func (s *Store) close(ctx context.Context) error {
	if err := s.client.Disconnect(ctx); err != nil {
		return &StoreError{
			Op:      "close",
			Err:     err,
			Details: "failed to disconnect mongodb client",
		}
	}

	return nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MONGODB_URI", "mongodb://shortlink:password@localhost:27017/shortlink") // MongoDB URI
	viper.SetDefault("STORE_MODE_WRITE", storeOptions.MODE_SINGLE_WRITE)                            // mode write to db

	s.config = Config{
		URI:  viper.GetString("STORE_MONGODB_URI"),
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
