/*
Store package
*/

package store

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/notify"
	"github.com/batazor/shortlink/internal/store/badger"
	"github.com/batazor/shortlink/internal/store/cassandra"
	"github.com/batazor/shortlink/internal/store/dgraph"
	"github.com/batazor/shortlink/internal/store/leveldb"
	"github.com/batazor/shortlink/internal/store/mongo"
	"github.com/batazor/shortlink/internal/store/mysql"
	"github.com/batazor/shortlink/internal/store/postgres"
	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/internal/store/ram"
	"github.com/batazor/shortlink/internal/store/redis"
	"github.com/batazor/shortlink/internal/store/rethinkdb"
	"github.com/batazor/shortlink/internal/store/scylla"
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
	case "mysql":
		store.store = &mysql.MySQLLinkList{}
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
	case "scylla":
		store.store = &scylla.ScyllaLinkList{}
	case "rethinkdb":
		store.store = &rethinkdb.RethinkDBLinkList{}
	case "ram":
		store.store = &ram.RAMLinkList{}
	default:
		store.store = &ram.RAMLinkList{}
	}

	if err := store.store.Init(ctx); err != nil {
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
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra
	s.typeStore = viper.GetString("STORE_TYPE")
}

// Notify ...
func (s *Store) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		if addLink, ok := payload.(*link.Link); ok {
			payload, err := s.store.Add(ctx, addLink)
			return notify.Response{
				Name:    "RESPONSE_STORE_ADD",
				Payload: payload,
				Error:   err,
			}
		}

		return notify.Response{
			Name:    "RESPONSE_STORE_ADD",
			Payload: payload,
			Error:   errors.New("failed assert type"),
		}
	case api_type.METHOD_GET:
		payload, err := s.store.Get(ctx, payload.(string))
		return notify.Response{
			Name:    "RESPONSE_STORE_GET",
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_LIST:
		f := payload.(string)

		// Parse filter
		var filter query.Filter
		if err := json.Unmarshal([]byte(f), &filter); err != nil {
			return notify.Response{
				Name:    "RESPONSE_STORE_LIST",
				Payload: payload,
				Error:   err,
			}
		}

		payload, err := s.store.List(ctx, &filter)
		return notify.Response{
			Name:    "RESPONSE_STORE_LIST",
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_UPDATE:
		if linkUpdate, ok := payload.(*link.Link); ok {
			payload, err := s.store.Update(ctx, linkUpdate)
			return notify.Response{
				Name:    "RESPONSE_STORE_UPDATE",
				Payload: payload,
				Error:   err,
			}
		}

		return notify.Response{
			Name:    "RESPONSE_STORE_UPDATE",
			Payload: payload,
			Error:   errors.New("failed assert type"),
		}
	case api_type.METHOD_DELETE:
		err := s.store.Delete(ctx, payload.(string))
		return notify.Response{
			Name:    "RESPONSE_STORE_DELETE",
			Payload: nil,
			Error:   err,
		}
	}

	return notify.Response{}
}
