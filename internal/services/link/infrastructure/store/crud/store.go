/*
Store Endpoint
*/
package crud

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

// New return implementation of db
func New(ctx context.Context, log logger.Logger, db *db.Store, cache *cache.Cache) (*Store, error) { // nolint unused
	s := &Store{
		log:   log,
		cache: cache,
	}

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
		s.store, err = postgres.New()
		if err != nil {
			return nil, err
		}
	case "mongo":
		s.store = &mongo.Store{}
	case "mysql":
		s.store = &mysql.Store{}
	case "redis":
		s.store = &redis.Store{}
	case "dgraph":
		s.store = dgraph.New(log)
	case "sqlite":
		s.store = &sqlite.Store{}
	case "leveldb":
		s.store = &leveldb.Store{}
	case "badger":
		s.store = &badger.Store{}
	case "ram":
		s.store = &ram.Store{}
	default:
		s.store = &ram.Store{}
	}

	// Init store
	err = s.store.Init(ctx, db)
	if err != nil {
		return nil, err
	}

	log.Info("init linkStore", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

func (s *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	// cache
	link := v1.Link{}
	err := s.cache.Get(ctx, fmt.Sprintf(`link:%s`, id), &link)
	if err != nil {
		s.log.ErrorWithContext(ctx, err.Error())
	}
	if err == nil {
		return &link, nil
	}

	response, err := s.store.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// save cache
	err = s.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf(`link:%s`, id),
		Value: &response,
		TTL:   5 * time.Minute,
	})
	if err != nil {
		s.log.ErrorWithContext(ctx, err.Error())
	}

	return response, err
}

func (s *Store) List(ctx context.Context, filter *query.Filter) (*v1.Links, error) {
	response, err := s.store.List(ctx, filter)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (s *Store) Add(ctx context.Context, in *v1.Link) (*v1.Link, error) {
	response, err := s.store.Add(ctx, in)
	if err != nil {
		return nil, err
	}

	return response, err
}

func (s *Store) Update(ctx context.Context, in *v1.Link) (*v1.Link, error) {
	response, err := s.store.Update(ctx, in)
	if err != nil {
		return nil, err
	}

	// update cache
	err = s.cache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   fmt.Sprintf(`link:%s`, in.Hash),
		Value: &response,
		TTL:   5 * time.Minute,
	})
	if err != nil {
		s.log.ErrorWithContext(ctx, err.Error())
	}

	return response, err
}

func (s *Store) Delete(ctx context.Context, id string) error {
	// drop from cache
	err := s.cache.Delete(ctx, fmt.Sprintf(`link:%s`, id))
	if err != nil {
		s.log.ErrorWithContext(ctx, err.Error())
	}

	err = s.store.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
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
			payload, err := s.store.Add(newCtx, addLink)
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

		payload, err := s.store.List(newCtx, &filter)
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
			payload, err := s.store.Update(newCtx, linkUpdate)
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

		err := s.store.Delete(newCtx, payload.(string))
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
