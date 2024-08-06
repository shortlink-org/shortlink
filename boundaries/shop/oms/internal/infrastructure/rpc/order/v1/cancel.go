package v1

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
)

func (o *OrderRPC) Cancel(ctx context.Context, in *v1.CancelRequest) (*emptypb.Empty, error) {
	// parse order ID to UUID
	orderId, err := uuid.Parse(in.GetId())
	if err != nil {
		return nil, err
	}

	err = o.orderService.Cancel(ctx, orderId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
