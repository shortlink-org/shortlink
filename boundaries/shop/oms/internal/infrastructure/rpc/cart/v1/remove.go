package v1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

func (c *CartRPC) Remove(ctx context.Context, in *v1.RemoveRequest) (*emptypb.Empty, error) {
	// TODO implement me
	panic("implement me")
}
