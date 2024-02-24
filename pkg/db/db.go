/*
Data Base package
*/
package db

import (
	"context"

	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/pkg/db/badger"
	"github.com/shortlink-org/shortlink/pkg/db/cockroachdb"
	"github.com/shortlink-org/shortlink/pkg/db/dgraph"
	"github.com/shortlink-org/shortlink/pkg/db/leveldb"
	"github.com/shortlink-org/shortlink/pkg/db/mongo"
	"github.com/shortlink-org/shortlink/pkg/db/neo4j"
	"github.com/shortlink-org/shortlink/pkg/db/postgres"
	"github.com/shortlink-org/shortlink/pkg/db/ram"
	"github.com/shortlink-org/shortlink/pkg/db/redis"
	"github.com/shortlink-org/shortlink/pkg/db/sqlite"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
)

// Use return implementation of db
func New(ctx context.Context, log logger.Logger, tracer trace.TracerProvider, metrics *metric.MeterProvider) (*Store, error) {
	store := &Store{}

	// Set configuration
	store.setConfig()

	switch store.typeStore {
	case "cockroachdb":
		store.DB = &cockroachdb.Store{}
	case "postgres":
		store.DB = postgres.New(tracer, metrics)
	case "mongo":
		store.DB = &mongo.Store{}
	case "redis":
		store.DB = redis.New(tracer, metrics)
	case "dgraph":
		store.DB = dgraph.New(log)
	case "leveldb":
		store.DB = &leveldb.Store{}
	case "badger":
		store.DB = &badger.Store{}
	case "ram":
		store.DB = &ram.Store{}
	case "neo4j":
		store.DB = &neo4j.Store{}
	case "sqlite":
		store.DB = sqlite.New(tracer, metrics)
	default:
		store.DB = &ram.Store{}
	}

	if err := store.Init(ctx); err != nil {
		return nil, err
	}

	log.Info("run db", field.Fields{
		"db": store.typeStore,
	})

	return store, nil
}

// setConfig - set configuration
func (s *Store) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres, mongo, redis, dgraph, sqlite, leveldb, badger, neo4j, ram
	s.typeStore = viper.GetString("STORE_TYPE")
}
