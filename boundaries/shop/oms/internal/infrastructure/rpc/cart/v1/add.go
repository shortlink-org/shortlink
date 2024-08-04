package v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/dto"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

// Add adds an item to the cart
func (c *CartRPC) Add(ctx context.Context, in *v1.AddRequest) (*emptypb.Empty, error) {
	request, err := dto.AddRequestToDomain(in)
	if err != nil {
		return nil, err
	}

	err = c.cartService.Add(ctx, request)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
