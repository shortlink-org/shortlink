package mongo

import (
	"context"
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/johejo/golang-migrate-extra/source/file"
	"github.com/johejo/golang-migrate-extra/source/iofs"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"

	storeOptions "github.com/batazor/shortlink/internal/pkg/db/options"
)

//go:embed migrations/*.json
var migrations embed.FS

// Init ...
func (m *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	m.setConfig()

	// Connect to MongoDB
	opts := options.Client().ApplyURI(m.config.URI)
	opts.Monitor = otelmongo.NewMonitor()
	m.client, err = mongo.NewClient(opts)
	if err != nil {
		return err
	}

	err = m.client.Connect(ctx)
	if err != nil {
		return err
	}

	// Check connect
	err = m.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	// Apply migration
	err = m.migrate()
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
func (m *Store) Close() error {
	return m.client.Disconnect(context.Background())
}

// Migrate ...
func (m *Store) migrate() error {
	driver, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}

	ms, err := migrate.NewWithSourceInstance("iofs", driver, m.config.URI)
	if err != nil {
		return err
	}

	err = ms.Up()
	if err != nil && err.Error() != "no change" {
		return err
	}

	return nil
}

// setConfig - set configuration
func (m *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_MONGODB_URI", "mongodb://shortlink:password@localhost:27017/shortlink") // MongoDB URI
	viper.SetDefault("STORE_MODE_WRITE", storeOptions.MODE_SINGLE_WRITE)                            // mode write to db

	m.config = Config{
		URI:  viper.GetString("STORE_MONGODB_URI"),
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
