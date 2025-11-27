package link_cqrs

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	cqrsmessage "github.com/shortlink-org/go-sdk/cqrs/message"
	"github.com/shortlink-org/go-sdk/logger"

	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/cqrs/query"
)

type Service struct {
	// CQRS
	subscriber message.Subscriber
	marshaler  *cqrsmessage.ProtoMarshaler

	// Repository
	cqsStore   cqs.Repository
	queryStore query.Repository

	log logger.Logger
}

func New(
	log logger.Logger,
	subscriber message.Subscriber,
	marshaler *cqrsmessage.ProtoMarshaler,
	cqsStore cqs.Repository,
	queryStore query.Repository,
) (*Service, error) {
	service := &Service{
		subscriber: subscriber,
		marshaler:  marshaler,
		cqsStore:   cqsStore,
		queryStore: queryStore,
		log:        log,
	}

	// Subscribe to events
	err := service.EventHandlers(context.Background())
	if err != nil {
		return nil, err
	}

	return service, nil
}
