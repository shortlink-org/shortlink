package link_cqrs

import (
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/cqrs/query"
)

type Service struct {
	// Observer interface for subscribe on system event
	// notify.Subscriber[link.Link]

	// Repository
	cqsStore   *cqs.Store
	queryStore *query.Store

	log logger.Logger
}

func New(log logger.Logger, cqsStore *cqs.Store, queryStore *query.Store) (*Service, error) {
	service := &Service{
		cqsStore:   cqsStore,
		queryStore: queryStore,

		log: log,
	}

	// Subscribe to event
	service.EventHandlers()

	return service, nil
}
