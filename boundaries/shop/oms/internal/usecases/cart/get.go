package cart

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/cart/dto"
)

// Get gets the cart.
func (uc *UC) Get(ctx context.Context, customerId uuid.UUID) (*v1.CartState, error) {
	workflowId := fmt.Sprintf("cart-%s", customerId)

	resp, err := uc.temporalClient.QueryWorkflow(ctx, workflowId, "", v1.Event_EVENT_GET.String(), nil)
	if err != nil {
		return nil, err
	}

	var cartState v3.CartState
	err = resp.Get(&cartState)
	if err != nil {
		return nil, err
	}

	state, err := dto.CartStateToCartState(&cartState)
	if err != nil {
		return nil, err
	}

	return state, nil
}
