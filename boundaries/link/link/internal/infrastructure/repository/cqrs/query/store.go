/*
Store Endpoint
*/
package query

import (
	"context"

	"github.com/go-redis/cache/v9"
	"github.com/spf13/viper"

	v12 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link_cqrs/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/cqrs/query/postgres"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

// New return implementation of db
func New(ctx context.Context, log logger.Logger, store db.DB, cacheStore *cache.Cache) (*Store, error) {
	s := &Store{
		log:   log,
		cache: cacheStore,
	}

	// Set configuration
	s.setConfig()

	var err error

	switch s.typeStore {
	case "postgres":
		fallthrough
	default:
		s.store, err = postgres.New(ctx, store)
		if err != nil {
			return nil, err
		}
	}

	log.Info("init queryStore", field.Fields{
		"store": s.typeStore,
	})

	return s, nil
}

func (s *Store) Get(ctx context.Context, id string) (*v12.LinkView, error) {
	return s.store.Get(ctx, id)
}

func (s *Store) List(ctx context.Context, filter *types.FilterLink) (*v12.LinksView, error) {
	return s.store.List(ctx, filter)
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}
