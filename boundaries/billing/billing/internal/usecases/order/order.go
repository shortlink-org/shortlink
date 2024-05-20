package order_application

import (
	"context"

	billing "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/order/v1"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/pattern/eventsourcing"
)

type OrderService struct {
	log logger.Logger

	// Repositories
	orderRepository eventsourcing.EventSourcing
}

func New(log logger.Logger, orderRepository eventsourcing.EventSourcing) (*OrderService, error) {
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
	// return o.orderRepository.list(ctx, filter)
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
