package order

import (
	"context"

	"github.com/google/uuid"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
)

func (uc *UC) Get(ctx context.Context, orderId uuid.UUID) (*v1.OrderState, error) {
	return nil, nil
}
