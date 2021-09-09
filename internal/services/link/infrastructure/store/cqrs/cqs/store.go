/*
Store Endpoint
*/
package cqs

import (
	"context"

	"github.com/go-redis/cache/v8"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	link "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/cqs/postgres"
	metadata "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"
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

	log.Info("init cqsStore", field.Fields{
		"db": s.typeStore,
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

func (s *Store) MetadataUpdate(ctx context.Context, data *metadata.Meta) (*metadata.Meta, error) {
	return s.store.MetadataUpdate(ctx, data)
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}
