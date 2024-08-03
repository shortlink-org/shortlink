package v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (c *CartRPC) Add(ctx context.Context, in *AddRequest) (*emptypb.Empty, error) {
	request, err := in.ToDomain()
	if err != nil {
		return nil, err
	}

	err = c.cartService.Add(ctx, request)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
