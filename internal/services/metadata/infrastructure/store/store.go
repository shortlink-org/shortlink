/*
Metadata Service. Infrastructure layer
*/
package meta_store

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/metadata/infrastructure/store/ram"
)

// Use return implementation of db
func (s *MetaStore) Use(_ context.Context, log logger.Logger, _ *db.Store) (*MetaStore, error) {
	// Set configuration
	s.setConfig()

	// Subscribe to Event
	notify.Subscribe(v1.METHOD_ADD, s)
	notify.Subscribe(v1.METHOD_GET, s)
	// notify.Subscribe(api_type.METHOD_LIST, store)
	// notify.Subscribe(api_domain.METHOD_UPDATE, store)
	// notify.Subscribe(api_domain.METHOD_DELETE, store)

	switch s.typeStore {
	case "ram":
		s.Store = &ram.Store{}
	default:
		s.Store = &ram.Store{}
	}

	log.Info("init metaStore", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

// Notify - ...
func (s *MetaStore) Notify(ctx context.Context, event uint32, payload any) notify.Response[any] {
	return notify.Response[any]{
		Name:    "RESPONSE_STORE_ADD",
		Payload: payload,
		Error:   nil,
	}
}

func (s *MetaStore) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra
	s.typeStore = viper.GetString("STORE_TYPE")
}
