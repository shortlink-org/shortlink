package order_application

import (
	"context"

	"github.com/batazor/shortlink/internal/pkg/logger"
	billing "github.com/batazor/shortlink/internal/services/billing/domain/billing/order/v1"
	billing_store "github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
)

type OrderService struct {
	logger logger.Logger

	// Repositories
	orderRepository billing_store.OrderRepository
}

func New(logger logger.Logger, orderRepository billing_store.OrderRepository) (*OrderService, error) {
	return &OrderService{
		logger: logger,

		// Repositories
		orderRepository: orderRepository,
	}, nil
}

func (o OrderService) Get(ctx context.Context, id string) (*billing.Order, error) {
	return o.orderRepository.Get(ctx, id)
}

func (o OrderService) List(ctx context.Context, filter interface{}) ([]*billing.Order, error) {
	return o.orderRepository.List(ctx, filter)
}

func (o OrderService) Add(ctx context.Context, in *billing.Order) (*billing.Order, error) {
	return o.orderRepository.Add(ctx, in)
}

func (o OrderService) Update(ctx context.Context, in *billing.Order) (*billing.Order, error) {
	return o.orderRepository.Update(ctx, in)
}

func (o OrderService) Delete(ctx context.Context, id string) error {
	return o.orderRepository.Delete(ctx, id)
}
