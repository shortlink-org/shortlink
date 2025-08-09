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

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/badger"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/dgraph"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/leveldb"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mongo"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mysql"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/postgres"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/ram"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/redis"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/sqlite"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

// New return implementation of db
func New(ctx context.Context, log logger.Logger, store db.DB, c *cache.Cache) (*Store, error) { //nolint:gocognit // ignore
	s := &Store{
		log:   log,
		cache: c,
	}

	// Set configuration
	s.setConfig()

	var err error

	switch s.typeStore {
	case "postgres":
		s.store, err = postgres.New(ctx, store)
		if err != nil {
			return nil, err
		}
	case "mysql":
		s.store, err = mysql.New(ctx, store)
		if err != nil {
			return nil, err
		}
	case "mongo":
		s.store, err = mongo.New(ctx, store)
		if err != nil {
			return nil, err
		}
	case "redis":
		s.store, err = redis.New(ctx, store)
		if err != nil {
			return nil, err
		}
	case "dgraph":
		s.store, err = dgraph.New(ctx, store, log)
		if err != nil {
			return nil, err
		}
	case "leveldb":
		s.store, err = leveldb.New(ctx, store, log)
		if err != nil {
			return nil, err
		}
	case "badger":
		s.store, err = badger.New(ctx)
		if err != nil {
			return nil, err
		}
	case "sqlite":
		s.store, err = sqlite.New(ctx, store)
		if err != nil {
			return nil, err
		}
	case "ram":
		s.store, err = ram.New(ctx)
		if err != nil {
			return nil, err
		}
	default:
		return nil, db.UnknownStoreTypeError{StoreType: s.typeStore}
	}

	log.Info("init linkStore", field.Fields{
		"store": s.typeStore,
	})

	return s, nil
}

func (s *Store) Get(ctx context.Context, id string) (*v1.Link, error) {
	link := &v1.Link{}

	err := s.cache.Get(ctx, fmt.Sprintf(`link:%s`, id), &link)
	if err != nil {
		s.log.ErrorWithContext(ctx, err.Error())
	}
	if err == nil && link.GetHash() != "" {
		return link, nil
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
		TTL:   5 * time.Minute, //nolint:mnd // ignore
	})
	if err != nil {
		s.log.ErrorWithContext(ctx, err.Error())
	}

	return response, err
}

func (s *Store) List(ctx context.Context, filter *types.FilterLink) (*v1.Links, error) {
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
		Key:   fmt.Sprintf(`link:%s`, in.GetHash()),
		Value: &response,
		TTL:   5 * time.Minute, //nolint:mnd // ignore
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
