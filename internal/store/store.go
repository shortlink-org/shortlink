package store

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/internal/store/badger"
	"github.com/batazor/shortlink/internal/store/cassandra"
	"github.com/batazor/shortlink/internal/store/dgraph"
	"github.com/batazor/shortlink/internal/store/leveldb"
	"github.com/batazor/shortlink/internal/store/mongo"
	"github.com/batazor/shortlink/internal/store/postgres"
	"github.com/batazor/shortlink/internal/store/ram"
	"github.com/batazor/shortlink/internal/store/redis"
	"github.com/batazor/shortlink/internal/store/sqlite"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

// Use return implementation of store
func (store *Store) Use(ctx context.Context, log logger.Logger) DB { // nolint unused
	// Set configuration
	store.setConfig()

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, store)
	notify.Subscribe(api_type.METHOD_GET, store)
	notify.Subscribe(api_type.METHOD_LIST, store)
	notify.Subscribe(api_type.METHOD_UPDATE, store)
	notify.Subscribe(api_type.METHOD_DELETE, store)

	switch store.typeStore {
	case "postgres":
		store.store = &postgres.PostgresLinkList{}
	case "mongo":
		store.store = &mongo.MongoLinkList{}
	case "redis":
		store.store = &redis.RedisLinkList{}
	case "dgraph":
		store.store = &dgraph.DGraphLinkList{}
	case "sqlite":
		store.store = &sqlite.SQLiteLinkList{}
	case "leveldb":
		store.store = &leveldb.LevelDBLinkList{}
	case "badger":
		store.store = &badger.BadgerLinkList{}
	case "cassandra":
		store.store = &cassandra.CassandraLinkList{}
	case "ram":
		store.store = &ram.RAMLinkList{}
	default:
		store.store = &ram.RAMLinkList{}
	}

	if err := store.store.Init(); err != nil {
		panic(err)
	}

	log.Info("run store", logger.Fields{
		"store": store.typeStore,
	})

	return store.store
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram")
	s.typeStore = viper.GetString("STORE_TYPE")
}

// Notify ...
func (s *Store) Notify(event int, payload interface{}) *notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		payload, err := s.store.Add(payload.(link.Link))
		return &notify.Response{
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_GET:
		payload, err := s.store.Get(payload.(string))
		return &notify.Response{
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_LIST:
		payload, err := s.store.List(nil)
		return &notify.Response{
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_UPDATE:
		payload, err := s.store.Update(payload.(link.Link))
		return &notify.Response{
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_DELETE:
		err := s.store.Delete(payload.(string))
		return &notify.Response{
			Payload: nil,
			Error:   err,
		}
	}

	return nil
}
