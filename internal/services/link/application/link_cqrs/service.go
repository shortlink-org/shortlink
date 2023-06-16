package link_cqrs

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/notify"
	link "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/cqrs/cqs"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/cqrs/query"
)

type Service struct {
	// Observer interface for subscribe on system event
	notify.Subscriber[link.Link]

	// Repository
	cqsStore   *cqs.Store
	queryStore *query.Store

	logger logger.Logger
}

func New(logger logger.Logger, cqsStore *cqs.Store, queryStore *query.Store) (*Service, error) {
	service := &Service{
		cqsStore:   cqsStore,
		queryStore: queryStore,

		logger: logger,
	}

	// Subscribe to event
	service.EventHandlers()

	return service, nil
}
