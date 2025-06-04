/*
Store Endpoint
*/
package cqs

import (
	"context"

	"github.com/go-redis/cache/v9"
	"github.com/spf13/viper"

	link "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/cqrs/cqs/postgres"
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
		s.store, err = postgres.New(ctx, store)
		if err != nil {
			return nil, err
		}
	default:
		return nil, db.UnknownStoreTypeError{StoreType: s.typeStore}
	}

	log.Info("init cqsStore", field.Fields{
		"store": s.typeStore,
	})

	return s, nil
}

func (s *Store) LinkAdd(ctx context.Context, data *link.Link) (*link.Link, error) {
	return s.store.LinkAdd(ctx, data)
}

func (s *Store) LinkUpdate(ctx context.Context, data *link.Link) (*link.Link, error) {
	return s.store.LinkUpdate(ctx, data)
}

func (s *Store) LinkDelete(ctx context.Context, id string) error {
	return s.store.LinkDelete(ctx, id)
}

// func (s *Store) MetadataUpdate(ctx context.Context, data *metadata.Meta) (*metadata.Meta, error) {
// 	return s.store.MetadataUpdate(ctx, data)
// }

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}
