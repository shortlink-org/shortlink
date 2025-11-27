/*
Store Endpoint
*/
package query

import (
	"context"
	"log/slog"

	"github.com/go-redis/cache/v9"
	"github.com/shortlink-org/go-sdk/db"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/spf13/viper"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
	v12 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link_cqrs/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/cqrs/query/postgres"
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

	log.InfoWithContext(ctx, "init queryStore",
		slog.String("store", s.typeStore),
	)

	return s, nil
}

func (s *Store) Get(ctx context.Context, id string) (*v12.LinkView, error) {
	return s.store.Get(ctx, id)
}

func (s *Store) List(ctx context.Context, filter *v1.FilterLink) (*v12.LinksView, error) {
	return s.store.List(ctx, filter)
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres

	s.typeStore = viper.GetString("STORE_TYPE")
}
