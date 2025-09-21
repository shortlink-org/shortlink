/*
Package event_store - implementation of event store
*/
package eventsourcing

import (
	"context"

	"github.com/spf13/viper"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/db"
	es_postgres "github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing/store/postgres"
)

// New - create new EventStore
func New(ctx context.Context, log logger.Logger, store db.DB) (EventSourcing, error) {
	var err error

	// Initialize EventStore
	e := &eventSourcing{}

	// Set configuration
	e.setConfig()

	switch e.typeStore {
	case "postgres":
		e.repository, err = es_postgres.New(ctx, store)
		if err != nil {
			return nil, err
		}
	default:
		e.repository, err = es_postgres.New(ctx, store)
		if err != nil {
			return nil, err
		}
	}

	log.Info("run db", field.Fields{
		"db": e.typeStore,
	})

	return e.repository, nil
}

// setConfig - set configuration
func (e *eventSourcing) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "ram") // Select: postgres

	e.typeStore = viper.GetString("STORE_TYPE")
}
