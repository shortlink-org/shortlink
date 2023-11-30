package mongo

import (
	"context"

	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/johejo/golang-migrate-extra/source/file"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

	storeOptions "github.com/shortlink-org/shortlink/internal/pkg/db/options"
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
		SetMonitor(otelmongo.NewMonitor()).
		SetRetryReads(true).
		SetRetryWrites(true)

	s.client, err = mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	// Check connect
	err = s.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		_ = s.close(ctx)
	}()

	return nil
}

// GetConn - get connect
func (s *Store) GetConn() any {
	return s.client
}

// Close - close
func (s *Store) close(ctx context.Context) error {
	return s.client.Disconnect(ctx)
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
