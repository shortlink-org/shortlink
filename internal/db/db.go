/*
Data Base package
*/

package db

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/db/badger"
	"github.com/batazor/shortlink/internal/db/cassandra"
	"github.com/batazor/shortlink/internal/db/dgraph"
	"github.com/batazor/shortlink/internal/db/leveldb"
	"github.com/batazor/shortlink/internal/db/mongo"
	"github.com/batazor/shortlink/internal/db/mysql"
	"github.com/batazor/shortlink/internal/db/postgres"
	"github.com/batazor/shortlink/internal/db/ram"
	"github.com/batazor/shortlink/internal/db/redis"
	"github.com/batazor/shortlink/internal/db/rethinkdb"
	"github.com/batazor/shortlink/internal/db/scylla"
	"github.com/batazor/shortlink/internal/db/sqlite"
	"github.com/batazor/shortlink/internal/logger"
)

// Use return implementation of db
func (store *Store) Use(ctx context.Context, log logger.Logger) (DB, error) { // nolint unused
	// Set configuration
	store.setConfig()

	switch store.typeStore {
	case "postgres":
		store.store = &postgres.Store{}
	case "mongo":
		store.store = &mongo.Store{}
	case "mysql":
		store.store = &mysql.Store{}
	case "redis":
		store.store = &redis.Store{}
	case "dgraph":
		store.store = &dgraph.Store{}
	case "sqlite":
		store.store = &sqlite.Store{}
	case "leveldb":
		store.store = &leveldb.Store{}
	case "badger":
		store.store = &badger.Store{}
	case "cassandra":
		store.store = &cassandra.Store{}
	case "scylla":
		store.store = &scylla.Store{}
	case "rethinkdb":
		store.store = &rethinkdb.Store{}
	case "ram":
		store.store = &ram.Store{}
	default:
		store.store = &ram.Store{}
	}

	if err := store.store.Init(ctx); err != nil {
		return nil, err
	}

	log.Info("run db", logger.Fields{
		"db": store.typeStore,
	})

	return store.store, nil
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra
	s.typeStore = viper.GetString("STORE_TYPE")
}
