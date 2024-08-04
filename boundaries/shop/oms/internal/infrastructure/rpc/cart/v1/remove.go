package v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/dto"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

func (c *CartRPC) Remove(ctx context.Context, in *v1.RemoveRequest) (*emptypb.Empty, error) {
	request, err := dto.RemoveRequestToDomain(in)
	if err != nil {
		return nil, err
	}

	// Signal the Temporal workflow to remove the items
	err = c.cartService.Remove(ctx, request)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
