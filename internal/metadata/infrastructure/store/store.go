package meta_store

import (
	"context"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/metadata/domain"
	"github.com/batazor/shortlink/internal/metadata/infrastructure/store/ram"
)

type Store interface {
	Get(context.Context, string) (*rpc.Meta, error)
	Add(context.Context, *rpc.Meta) error
}

type MetaStore struct {
	store     Store
	typeStore string
}

// Use return implementation of db
func (store *MetaStore) Use(ctx context.Context, log logger.Logger) error {
	// Set configuration
	store.setConfig()

	switch store.typeStore {
	case "ram":
		store.store = &ram.Store{}
	default:
		store.store = &ram.Store{}
	}

	return nil
}

func (store *MetaStore) setConfig() {}
