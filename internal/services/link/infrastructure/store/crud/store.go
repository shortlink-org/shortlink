/*
Store Endpoint
*/
package crud

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/cache/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/notify"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/badger"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/dgraph"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/leveldb"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/mongo"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/mysql"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/postgres"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/ram"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/redis"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/sqlite"
)

// Use return implementation of db
func (s *Store) Use(ctx context.Context, log logger.Logger, db *db.Store, cache *cache.Cache) (*Store, error) { // nolint unused
	// Set configuration
	s.setConfig()

	// Subscribe to Event
	//notify.Subscribe(api_domain.METHOD_ADD, store)
	//notify.Subscribe(api_domain.METHOD_GET, store)
	//notify.Subscribe(api_domain.METHOD_LIST, store)
	//notify.Subscribe(api_domain.METHOD_UPDATE, store)
	//notify.Subscribe(api_domain.METHOD_DELETE, store)

	var err error

	switch s.typeStore {
	case "postgres":
		s.Repository, err = postgres.New(log, cache)
	case "mongo":
		s.Repository = &mongo.Store{}
	case "mysql":
		s.Repository = &mysql.Store{}
	case "redis":
		s.Repository = &redis.Store{}
	case "dgraph":
		s.Repository = dgraph.New(log)
	case "sqlite":
		s.Repository = &sqlite.Store{}
	case "leveldb":
		s.Repository = &leveldb.Store{}
	case "badger":
		s.Repository = &badger.Store{}
	case "ram":
		s.Repository = &ram.Store{}
	default:
		s.Repository = &ram.Store{}
	}
	if err != nil {
		return nil, err
	}

	// Init store
	err = s.Init(ctx, db)
	if err != nil {
		return nil, err
	}

	log.Info("init linkStore", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

// Notify ...
func (s *Store) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response { // nolint unused
	switch event {
	case v1.METHOD_ADD:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store add new link")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

		if addLink, ok := payload.(*v1.Link); ok {
			payload, err := s.Add(newCtx, addLink)
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
	case v1.METHOD_GET:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store get link")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

		link, err := s.Get(newCtx, payload.(string))
		return notify.Response{
			Name:    "RESPONSE_STORE_GET",
			Payload: link,
			Error:   err,
		}
	case v1.METHOD_LIST:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store get links")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

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

		payload, err := s.List(newCtx, &filter)
		return notify.Response{
			Name:    "RESPONSE_STORE_LIST",
			Payload: payload,
			Error:   err,
		}
	case v1.METHOD_UPDATE:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store update link")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

		if linkUpdate, ok := payload.(*v1.Link); ok {
			payload, err := s.Update(newCtx, linkUpdate)
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
	case v1.METHOD_DELETE:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store delete link")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

		err := s.Delete(newCtx, payload.(string))
		return notify.Response{
			Name:    "RESPONSE_STORE_DELETE",
			Payload: nil,
			Error:   err,
		}
	}

	return notify.Response{}
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram
	s.typeStore = viper.GetString("STORE_TYPE")
}
