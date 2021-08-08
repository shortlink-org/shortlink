/*
Store Endpoint
*/
package cqs

import (
	"context"

	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/logger/field"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/cqrs/cqs/postgres"
)

// Use return implementation of db
func (s *Store) Use(ctx context.Context, log logger.Logger, db *db.Store) (*Store, error) { // nolint unused
	// Set configuration
	s.setConfig()

	switch s.typeStore {
	case "postgres":
		s.Repository = &postgres.Store{}
	default:
		s.Repository = &postgres.Store{}
	}

	// Init store
	if err := s.Init(ctx, db); err != nil {
		return nil, err
	}

	log.Info("init linkStore", field.Fields{
		"db": s.typeStore,
	})

	return s, nil
}

// setConfig - set configuration
func (s *Store) setConfig() { // nolint unused
	viper.AutomaticEnv()
	viper.SetDefault("STORE_TYPE", "postgres") // Select: postgres
	s.typeStore = viper.GetString("STORE_TYPE")
}
