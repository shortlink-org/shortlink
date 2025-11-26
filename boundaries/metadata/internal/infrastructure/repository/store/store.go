// Package storeRepository contains metadata persistence infrastructure logic.
package storeRepository

import (
	"context"
	"log/slog"

	"github.com/shortlink-org/go-sdk/db"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/repository/store/ram"
)

// Use return implementation of db
func (s *MetaStore) Use(ctx context.Context, log logger.Logger, _ db.DB) (*MetaStore, error) {
	// Set configuration
	s.setConfig()

	switch s.typeStore {
	case "ram":
		s.Store = &ram.Store{}
	default:
		s.Store = &ram.Store{}
	}

	log.InfoWithContext(ctx, "init metaStore",
		slog.String("db", s.typeStore),
	)

	return s, nil
}

func (s *MetaStore) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, redis, dgraph, sqlite, leveldb, badger, ram, scylla, cassandra

	s.typeStore = viper.GetString("STORE_TYPE")
}
