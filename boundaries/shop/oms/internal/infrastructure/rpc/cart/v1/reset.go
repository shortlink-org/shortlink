package v1

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

// Reset resets the cart
func (c *CartRPC) Reset(ctx context.Context, in *v1.ResetRequest) (*emptypb.Empty, error) {
	// customerId to uuid
	customerId, err := uuid.Parse(in.CustomerId)
	if err != nil {
		return nil, err
	}

	err = c.cartService.Reset(ctx, customerId)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
