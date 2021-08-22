/*
Data Base package
*/
package db

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db/badger"
	"github.com/batazor/shortlink/internal/pkg/db/dgraph"
	"github.com/batazor/shortlink/internal/pkg/db/leveldb"
	"github.com/batazor/shortlink/internal/pkg/db/mongo"
	"github.com/batazor/shortlink/internal/pkg/db/mysql"
	"github.com/batazor/shortlink/internal/pkg/db/postgres"
	"github.com/batazor/shortlink/internal/pkg/db/ram"
	"github.com/batazor/shortlink/internal/pkg/db/redis"
	"github.com/batazor/shortlink/internal/pkg/db/sqlite"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
)

//gocyclo:ignore
// Use return implementation of db
func (store *Store) Use(ctx context.Context, log logger.Logger) (*Store, error) {
	// Set configuration
	store.setConfig()

	switch store.typeStore {
	case "postgres":
		store.Store = &postgres.Store{}
	case "mongo":
		store.Store = &mongo.Store{}
	case "mysql":
		store.Store = &mysql.Store{}
	case "redis":
		store.Store = &redis.Store{}
	case "dgraph":
		store.Store = dgraph.New(log)
	case "sqlite":
		store.Store = &sqlite.Store{}
	case "leveldb":
		store.Store = &leveldb.Store{}
	case "badger":
		store.Store = &badger.Store{}
	case "ram":
		store.Store = &ram.Store{}
	default:
		store.Store = &ram.Store{}
	}

	if err := store.Store.Init(ctx); err != nil {
		return nil, err
	}

	log.Info("run db", field.Fields{
		"db": store.typeStore,
	})

	return store, nil
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram
	s.typeStore = viper.GetString("STORE_TYPE")
}
