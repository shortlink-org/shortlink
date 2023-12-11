package order_application

import (
	"context"

	event_store "github.com/shortlink-org/shortlink/internal/pkg/eventsourcing/store"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	billing "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/order/v1"
)

type OrderService struct {
	log logger.Logger

	// Repositories
	orderRepository event_store.EventStore
}

func New(log logger.Logger, orderRepository event_store.EventStore) (*OrderService, error) {
	return &OrderService{
		log: log,

		// Repositories
		orderRepository: orderRepository,
	}, nil
}

func (o OrderService) Get(ctx context.Context, id string) (*billing.Order, error) {
	// return o.orderRepository.Get(ctx, id)
	panic("implement me")
}

func (o OrderService) List(ctx context.Context, filter any) ([]*billing.Order, error) {
	// return o.orderRepository.List(ctx, filter)
	panic("implement me")
}

func (o OrderService) Add(ctx context.Context, in *billing.Order) (*billing.Order, error) {
	// return o.orderRepository.Add(ctx, in)
	panic("implement me")
}

func (o OrderService) Update(ctx context.Context, in *billing.Order) (*billing.Order, error) {
	// return o.orderRepository.Update(ctx, in)
	panic("implement me")
}

func (o OrderService) Delete(ctx context.Context, id string) error {
	// return o.orderRepository.Delete(ctx, id)
	panic("implement me")
}
