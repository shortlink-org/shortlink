/*
Billing Service. Infrastructure layer
*/
package event_store

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	es_postgres "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/store/postgres"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

// Use return implementation of db
func (s *Repository) Use(ctx context.Context, log logger.Logger, eventStore db.DB) (*Repository, error) {
	// Set configuration
	s.setConfig()

	switch s.typeStore {
	case "postgres":
		s.EventStore = &es_postgres.Store{}
	default:
		s.EventStore = &es_postgres.Store{}
	}

	if err := s.EventStore.Init(ctx, eventStore); err != nil {
		return nil, err
	}

	log.Info("run db", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

// setConfig - set configuration
func (s *Repository) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}
