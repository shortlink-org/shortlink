/*
Store Endpoint
*/
package store

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/badger"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/cassandra"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/dgraph"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/leveldb"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/mongo"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/mysql"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/postgres"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/ram"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/redis"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/rethinkdb"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/scylla"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/sqlite"
	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/notify"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// Use return implementation of db
func (store *LinkStore) Use(_ context.Context, log logger.Logger, _ *db.Store) (*LinkStore, error) { // nolint unused
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
		store.Store = &postgres.Store{}
	case "mongo":
		store.Store = &mongo.Store{}
	case "mysql":
		store.Store = &mysql.Store{}
	case "redis":
		store.Store = &redis.Store{}
	case "dgraph":
		store.Store = &dgraph.Store{}
	case "sqlite":
		store.Store = &sqlite.Store{}
	case "leveldb":
		store.Store = &leveldb.Store{}
	case "badger":
		store.Store = &badger.Store{}
	case "cassandra":
		store.Store = &cassandra.CassandraLinkList{}
	case "scylla":
		store.Store = &scylla.Store{}
	case "rethinkdb":
		store.Store = &rethinkdb.Store{}
	case "ram":
		store.Store = &ram.Store{}
	default:
		store.Store = &ram.Store{}
	}

	log.Info("init linkStore", logger.Fields{
		"db": store.typeStore,
	})

	return store, nil
}

// Notify ...
func (s *LinkStore) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response { // nolint unused
	switch event {
	case api_type.METHOD_ADD:
		if addLink, ok := payload.(*link.Link); ok {
			payload, err := s.Store.Add(ctx, addLink)
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
		link, err := s.Store.Get(ctx, payload.(string))
		return notify.Response{
			Name:    "RESPONSE_STORE_GET",
			Payload: link,
			Error:   err,
		}
	case api_type.METHOD_LIST:
		filterRaw := ""
		if payload != nil {
			filterRaw = payload.(string)
		}

		// Parse filter
		var filter query.Filter
		if filterRaw != "" {
			if err := json.Unmarshal([]byte(filterRaw), &filter); err != nil {
				return notify.Response{
					Name:    "RESPONSE_STORE_LIST",
					Payload: payload,
					Error:   err,
				}
			}
		}

		payload, err := s.Store.List(ctx, &filter)
		return notify.Response{
			Name:    "RESPONSE_STORE_LIST",
			Payload: payload,
			Error:   err,
		}
	case api_type.METHOD_UPDATE:
		if linkUpdate, ok := payload.(*link.Link); ok {
			payload, err := s.Store.Update(ctx, linkUpdate)
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
		err := s.Store.Delete(ctx, payload.(string))
		return notify.Response{
			Name:    "RESPONSE_STORE_DELETE",
			Payload: nil,
			Error:   err,
		}
	}

	return notify.Response{}
}

// setConfig - set configuration
func (s *LinkStore) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra
	s.typeStore = viper.GetString("STORE_TYPE")
}
