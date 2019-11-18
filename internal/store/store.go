package store

import (
	"context"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/pkg/link"
	"github.com/spf13/viper"
)

// DB - common interface of store
type DB interface { // nolint unused
	Init() error
	Close() error

	Get(id string) (*link.Link, error)
	List() ([]*link.Link, error)
	Add(data link.Link) (*link.Link, error)
	Update(data link.Link) (*link.Link, error)
	Delete(id string) error
}

// Store abstract type
type Store struct { // nolint unused
	typeStore string
}

// Use return implementation of store
func (s *Store) Use(ctx context.Context) DB {
	var store DB
	log := logger.GetLogger(ctx)

	// Set configuration
	s.setConfig()

	switch s.typeStore {
	case "postgres":
		store = &PostgresLinkList{}
	case "mongo":
		store = &MongoLinkList{}
	case "redis":
		store = &RedisLinkList{}
	case "dgraph":
		store = &DGraphLinkList{}
	case "sqlite":
		store = &SQLiteLinkList{}
	case "leveldb":
		store = &LevelDBLinkList{}
	case "badger":
		store = &BadgerLinkList{}
	case "ram":
		store = &RAMLinkList{}
	default:
		store = &RAMLinkList{}
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
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram")
	s.typeStore = viper.GetString("STORE_TYPE")
}
