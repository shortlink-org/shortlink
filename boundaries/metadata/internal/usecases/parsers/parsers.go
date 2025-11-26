// Package parsers contains metadata extraction application logic.
package parsers

import (
	"github.com/shortlink-org/go-sdk/cqrs/bus"
	"github.com/shortlink-org/go-sdk/logger"

	meta_store "github.com/shortlink-org/shortlink/boundaries/metadata/internal/infrastructure/repository/store"
)

type UC struct {
	MetaStore *meta_store.MetaStore
	EventBus  *bus.EventBus
	log       logger.Logger
}

func New(store *meta_store.MetaStore, eventBus *bus.EventBus, log logger.Logger) (*UC, error) {
	return &UC{
		MetaStore: store,
		EventBus:  eventBus,
		log:       log,
	}, nil
}
