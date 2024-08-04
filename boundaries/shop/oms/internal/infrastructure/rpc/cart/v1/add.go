package v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Add adds an item to the cart
func (c *CartRPC) Add(ctx context.Context, in *AddRequest) (*emptypb.Empty, error) {
	request, err := AddRequestToDomain(in)
	if err != nil {
		return nil, err
	}

	err = c.cartService.Add(ctx, request)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
