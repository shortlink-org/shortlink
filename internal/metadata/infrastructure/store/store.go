package meta_store

import (
	"context"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/metadata/infrastructure/store/ram"
)

// Use return implementation of db
func (store *MetaStore) Use(ctx context.Context, log logger.Logger) error {
	// Set configuration
	store.setConfig()

	switch store.typeStore {
	case "ram":
		store.Store = &ram.Store{}
	default:
		store.Store = &ram.Store{}
	}

	return nil
}

func (store *MetaStore) setConfig() {}
