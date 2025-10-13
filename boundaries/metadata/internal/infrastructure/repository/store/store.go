/*
Metadata Service. Infrastructure layer
*/
package storeRepository

import (
	"context"
	"log/slog"

	"github.com/spf13/viper"

	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/repository/store/ram"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/notify"
)

// Use return implementation of db
func (s *MetaStore) Use(_ context.Context, log logger.Logger, _ db.DB) (*MetaStore, error) {
	// Set configuration
	s.setConfig()

	// Subscribe to Event
	// notify.Subscribe(v1.METHOD_ADD, s)
	// notify.Subscribe(v1.METHOD_GET, s)
	// notify.Subscribe(v1.METHOD_LIST, s)
	// notify.Subscribe(v1.METHOD_UPDATE, s)
	// notify.Subscribe(v1.METHOD_DELETE, s)

	switch s.typeStore {
	case "ram":
		s.Store = &ram.Store{}
	default:
		s.Store = &ram.Store{}
	}

	log.Info("init metaStore",
		slog.String("db", s.typeStore),
	)

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
