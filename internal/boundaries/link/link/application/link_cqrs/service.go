package link_cqrs

import (
	link "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/repository/cqrs/query"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
)

type Service struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]

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
