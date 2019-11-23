package store

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/store/badger"
	"github.com/batazor/shortlink/internal/store/dgraph"
	"github.com/batazor/shortlink/internal/store/leveldb"
	"github.com/batazor/shortlink/internal/store/mongo"
	"github.com/batazor/shortlink/internal/store/postgres"
	"github.com/batazor/shortlink/internal/store/ram"
	"github.com/batazor/shortlink/internal/store/redis"
	"github.com/batazor/shortlink/internal/store/sqlite"
)

// Use return implementation of store
func (s *Store) Use(ctx context.Context) DB { // nolint unused
	var store DB
	log := logger.GetLogger(ctx)

	// Set configuration
	s.setConfig()

	switch s.typeStore {
	case "postgres":
		store = &postgres.PostgresLinkList{}
	case "mongo":
		store = &mongo.MongoLinkList{}
	case "redis":
		store = &redis.RedisLinkList{}
	case "dgraph":
		store = &dgraph.DGraphLinkList{}
	case "sqlite":
		store = &sqlite.SQLiteLinkList{}
	case "leveldb":
		store = &leveldb.LevelDBLinkList{}
	case "badger":
		store = &badger.BadgerLinkList{}
	case "cassandra":
		store = &cassandra.CassandraLinkList{}
	case "ram":
		store = &ram.RAMLinkList{}
	default:
		store = &ram.RAMLinkList{}
	}

	if err := store.Init(); err != nil {
		panic(err)
	}

	log.Info("run store", logger.Fields{
		"store": s.typeStore,
	})

	return store
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram")
	s.typeStore = viper.GetString("STORE_TYPE")
}
