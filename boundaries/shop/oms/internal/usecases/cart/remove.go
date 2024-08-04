package cart

import (
	"context"
	"fmt"

	domain "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/cart/dto"
)

// Remove removes items from the cart.
func (uc *UC) Remove(ctx context.Context, in *domain.CartState) error {
	workflowId := fmt.Sprintf("cart-%s", in.GetCustomerId().String())

	request := dto.CartStateToCartEvent(in, v1.Event_EVENT_REMOVE)
	err := uc.temporalClient.SignalWorkflow(ctx, workflowId, "", v1.Event_EVENT_REMOVE.String(), request)
	if err != nil {
		return err
	}

	return nil
}
