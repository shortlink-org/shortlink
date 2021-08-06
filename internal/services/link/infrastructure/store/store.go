/*
Store Endpoint
*/
package store

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/pkg/notify"
	"github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/badger"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/dgraph"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/leveldb"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/mongo"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/mysql"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/postgres"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/query"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/ram"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/redis"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/rethinkdb" // nolint staticcheck
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/sqlite"
)

// Use return implementation of db
func (s *LinkStore) Use(ctx context.Context, log logger.Logger, db *db.Store) (*LinkStore, error) { // nolint unused
	// Set configuration
	s.setConfig()

	// Subscribe to Event
	//notify.Subscribe(uint32(link.LinkEvent_LINK_EVENT_ADD), store)
	//notify.Subscribe(uint32(link.LinkEvent_LINK_EVENT_GET), store)
	//notify.Subscribe(uint32(link.LinkEvent_LINK_EVENT_LIST), store)
	//notify.Subscribe(uint32(link.LinkEvent_LINK_EVENT_UPDATE), store)
	//notify.Subscribe(uint32(link.LinkEvent_LINK_EVENT_DELETE), store)

	switch s.typeStore {
	case "postgres":
		s.Store = &postgres.Store{}
	case "mongo":
		s.Store = &mongo.Store{}
	case "mysql":
		s.Store = &mysql.Store{}
	case "redis":
		s.Store = &redis.Store{}
	case "dgraph":
		s.Store = &dgraph.Store{}
	case "sqlite":
		s.Store = &sqlite.Store{}
	case "leveldb":
		s.Store = &leveldb.Store{}
	case "badger":
		s.Store = &badger.Store{}
	case "rethinkdb":
		s.Store = &rethinkdb.Store{}
	case "ram":
		s.Store = &ram.Store{}
	default:
		s.Store = &ram.Store{}
	}

	// Init store
	if err := s.Store.Init(ctx, db); err != nil {
		return nil, err
	}

	log.Info("init linkStore", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

// Notify ...
func (s *LinkStore) Notify(ctx context.Context, event uint32, payload interface{}) notify.Response { // nolint unused
	switch v1.LinkEvent(event) {
	case v1.LinkEvent_LINK_EVENT_ADD:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store add new link")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

		if addLink, ok := payload.(*v1.Link); ok {
			payload, err := s.Store.Add(newCtx, addLink)
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
	case v1.LinkEvent_LINK_EVENT_GET:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store get link")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

		link, err := s.Store.Get(newCtx, payload.(string))
		return notify.Response{
			Name:    "RESPONSE_STORE_GET",
			Payload: link,
			Error:   err,
		}
	case v1.LinkEvent_LINK_EVENT_LIST:
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

		payload, err := s.Store.List(newCtx, &filter)
		return notify.Response{
			Name:    "RESPONSE_STORE_LIST",
			Payload: payload,
			Error:   err,
		}
	case v1.LinkEvent_LINK_EVENT_UPDATE:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store update link")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

		if linkUpdate, ok := payload.(*v1.Link); ok {
			payload, err := s.Store.Update(newCtx, linkUpdate)
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
	case v1.LinkEvent_LINK_EVENT_DELETE:
		// start tracing
		span, newCtx := opentracing.StartSpanFromContext(ctx, "store delete link")
		span.SetTag("store", s.typeStore)
		defer span.Finish()

		err := s.Store.Delete(newCtx, payload.(string))
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
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram
	s.typeStore = viper.GetString("STORE_TYPE")
}
