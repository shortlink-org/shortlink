/*
Store Endpoint
*/
package query

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	v12 "github.com/batazor/shortlink/internal/services/link/domain/link_cqrs/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/query/postgres"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
)

// New return implementation of db
func New(ctx context.Context, log logger.Logger, db *db.Store, cache *cache.Cache) (*Store, error) { // nolint unused
	s := &Store{
		log:   log,
		cache: cache,
	}

	// Set configuration
	s.setConfig()

	var err error

	switch s.typeStore {
	case "postgres":
		fallthrough
	default:
		s.store, err = postgres.New(ctx, db)
		if err != nil {
			return nil, err
		}
	}

	log.Info("init queryStore", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

func (s *Store) Get(ctx context.Context, id string) (*v12.LinkView, error) {
	return s.store.Get(ctx, id)
}

func (s *Store) List(ctx context.Context, filter *query.Filter) (*v12.LinksView, error) {
	if filter.Pagination == nil {
		filter.Pagination = &query.Pagination{
			Page:  0,
			Limit: 10,
		}
	}

	return s.store.List(ctx, filter)
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}
