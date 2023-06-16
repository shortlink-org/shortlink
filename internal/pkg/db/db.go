/*
Data Base package
*/
package db

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/db/badger"
	"github.com/shortlink-org/shortlink/internal/pkg/db/cockroachdb"
	"github.com/shortlink-org/shortlink/internal/pkg/db/dgraph"
	"github.com/shortlink-org/shortlink/internal/pkg/db/leveldb"
	"github.com/shortlink-org/shortlink/internal/pkg/db/mongo"
	"github.com/shortlink-org/shortlink/internal/pkg/db/neo4j"
	"github.com/shortlink-org/shortlink/internal/pkg/db/postgres"
	"github.com/shortlink-org/shortlink/internal/pkg/db/ram"
	"github.com/shortlink-org/shortlink/internal/pkg/db/redis"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

// Use return implementation of db
func (store *Store) Use(ctx context.Context, log logger.Logger) (*Store, error) {
	// Set configuration
	store.setConfig()

	switch store.typeStore {
	case "cockroachdb":
		store.Store = &cockroachdb.Store{}
	case "postgres":
		store.Store = &postgres.Store{}
	case "mongo":
		store.Store = &mongo.Store{}
	case "redis":
		store.Store = &redis.Store{}
	case "dgraph":
		store.Store = dgraph.New(log)
	case "leveldb":
		store.Store = &leveldb.Store{}
	case "badger":
		store.Store = &badger.Store{}
	case "ram":
		store.Store = &ram.Store{}
	case "neo4j":
		store.Store = &neo4j.Store{}
	case "sqlite":
		// disabled because it complicates cross-compilation
		// store.Store = &sqlite.Store{}
		fallthrough
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
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, redis, dgraph, sqlite, leveldb, badger, neo4j, ram
	s.typeStore = viper.GetString("STORE_TYPE")
}
