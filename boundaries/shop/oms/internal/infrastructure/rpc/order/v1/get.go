package v1

import (
	"context"

	"github.com/google/uuid"

	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/dto"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
)

func (o *OrderRPC) Get(ctx context.Context, in *v1.GetRequest) (*v1.GetResponse, error) {
	// parse order ID to UUID
	orderId, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, err
	}

	orderState, err := o.orderService.Get(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return &v1.GetResponse{
		Order: dto.DomainToOrderState(orderState),
	}, nil
}
