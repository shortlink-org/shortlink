//go:generate go-bindata -prefix migrations -pkg migrations -ignore migrations.go -o migrations/migrations.go migrations
package mongo

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/batazor/shortlink/internal/db/mongo/migrations"
	storeOptions "github.com/batazor/shortlink/internal/db/options"
)

// Init ...
func (m *Store) Init(ctx context.Context) error {
	var err error

	// Set configuration
	m.setConfig()

	// Connect to MongoDB
	m.client, err = mongo.NewClient(options.Client().ApplyURI(m.config.URI))
	if err != nil {
		return err
	}

	err = m.client.Connect(ctx)
	if err != nil {
		return err
	}

	// TODO: check correct ping
	// Check connect
	//ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	//err = m.client.Ping(ctx, readpref.Primary())
	//if err != nil {
	//	return err
	//}

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
func (m *Store) migrate() error { // nolint unused
	// wrap assets into Resource
	s := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})

	driver, err := bindata.WithInstance(s)
	if err != nil {
		return err
	}

	ms, err := migrate.NewWithSourceInstance("go-bindata", driver, m.config.URI)
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
	viper.SetDefault("STORE_MONGODB_URI", "mongodb://localhost:27017/shortlink") // MongoDB URI
	viper.SetDefault("STORE_MODE_WRITE", storeOptions.MODE_SINGLE_WRITE)         // mode write to db

	m.config = Config{
		URI:  viper.GetString("STORE_MONGODB_URI"),
		mode: viper.GetInt("STORE_MODE_WRITE"),
	}
}
