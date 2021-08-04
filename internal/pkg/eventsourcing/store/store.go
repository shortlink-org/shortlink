/*
Billing Service. Infrastructure layer
*/
package event_store

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db"
	es_postgres "github.com/batazor/shortlink/internal/pkg/eventsourcing/store/postgres"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
)

//gocyclo:ignore
// Use return implementation of db
func (s *Repository) Use(ctx context.Context, log logger.Logger, db *db.Store) (*Repository, error) {
	// Set configuration
	s.setConfig()

	switch s.typeStore {
	case "postgres":
		s.EventStore = &es_postgres.Store{}
	default:
		s.EventStore = &es_postgres.Store{}
	}

	if err := s.EventStore.Init(ctx, db); err != nil {
		return nil, err
	}

	log.Info("run db", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

// setConfig - set configuration
func (s *Repository) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}
