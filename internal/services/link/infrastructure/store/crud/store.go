/*
Store Endpoint
*/
package crud

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/badger"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/dgraph"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/leveldb"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/mongo"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/postgres"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/ram"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/redis"
)

// New return implementation of db
func New(ctx context.Context, log logger.Logger, db *db.Store, cache *cache.Cache) (*Store, error) { // nolint:gocognit
	s := &Store{
		log:   log,
		cache: cache,
	}

	// Set configuration
	s.setConfig()

	var err error

	switch s.typeStore {
	case "postgres":
		s.store, err = postgres.New(ctx, db)
		if err != nil {
			return nil, err
		}
	case "mongo":
		s.store, err = mongo.New(ctx, db)
		if err != nil {
			return nil, err
		}
	case "redis":
		s.store, err = redis.New(ctx, db)
		if err != nil {
			return nil, err
		}
	case "dgraph":
		s.store, err = dgraph.New(ctx, db, log)
		if err != nil {
			return nil, err
		}
	case "leveldb":
		s.store, err = leveldb.New(ctx, db)
		if err != nil {
			return nil, err
		}
	case "badger":
		s.store, err = badger.New(ctx)
		if err != nil {
			return nil, err
		}
	case "ram":
		fallthrough
	default:
		s.store, err = ram.New(ctx)
		if err != nil {
			return nil, err
		}
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
		TTL:   5 * time.Minute, // nolint:gomnd
	})
	if err != nil {
		s.log.ErrorWithContext(ctx, err.Error())
	}

	return response, err
}

func (s *Store) List(ctx context.Context, filter *query.Filter) (*v1.Links, error) {
	if filter.Pagination == nil {
		filter.Pagination = &query.Pagination{
			Page:  0,
			Limit: 10, // nolint:gomnd
		}
	}

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
		TTL:   5 * time.Minute, // nolint:gomnd
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

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, redis, dgraph, sqlite, leveldb, badger, ram
	s.typeStore = viper.GetString("STORE_TYPE")
}
