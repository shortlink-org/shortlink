/*
Metadata Service. Infrastructure layer
*/
package meta_store

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/db"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/logger/field"
	"github.com/batazor/shortlink/internal/metadata/infrastructure/store/ram"
	"github.com/batazor/shortlink/internal/notify"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// Use return implementation of db
func (s *MetaStore) Use(_ context.Context, log logger.Logger, _ *db.Store) (*MetaStore, error) {
	// Set configuration
	s.setConfig()

	// Subscribe to Event
	notify.Subscribe(api_type.METHOD_ADD, s)
	notify.Subscribe(api_type.METHOD_GET, s)
	//notify.Subscribe(api_type.METHOD_LIST, store)
	//notify.Subscribe(api_type.METHOD_UPDATE, store)
	//notify.Subscribe(api_type.METHOD_DELETE, store)

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

func (s *MetaStore) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, mysql, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra
	s.typeStore = viper.GetString("STORE_TYPE")
}
