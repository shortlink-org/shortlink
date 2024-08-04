package v1

import (
	"context"

	"github.com/google/uuid"

	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/dto"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
)

func (c *CartRPC) Get(ctx context.Context, in *v1.GetRequest) (*v1.GetResponse, error) {
	// customerId to uuid
	customerId, err := uuid.Parse(in.CustomerId)
	if err != nil {
		return nil, err
	}

	response, err := c.cartService.Get(ctx, customerId)
	if err != nil {
		return nil, err
	}

	return dto.GetResponseFromDomain(response), nil
}
