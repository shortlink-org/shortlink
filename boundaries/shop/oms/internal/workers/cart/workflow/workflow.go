package cart_workflow

import (
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"

	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/cart/v1/model/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/cart/workflow/dto"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/cart/workflow/model/cart/v1"
)

// Workflow is a Temporal workflow that manages the cart state.
func Workflow(ctx workflow.Context, customerId uuid.UUID) error {
	state := v2.NewCartState(customerId)

	// Set up query handler for getting cart state
	err := workflow.SetQueryHandler(ctx, v2.Event_EVENT_GET.String(), func() (*v3.CartState, error) {
		return dto.CartStateToDomain(state), nil
	})
	if err != nil {
		return err
	}

	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	addToCartChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_ADD.String())
	removeFromCartChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_REMOVE.String())
	resetCartChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_RESET.String())

	selector := workflow.NewSelector(ctx)

	selector.AddReceive(addToCartChannel, func(c workflow.ReceiveChannel, _ bool) {
		var request v1.CartEvent
		c.Receive(ctx, &request)

		for _, item := range request.Items {
			productId, err := uuid.Parse(item.ProductId)
			if err != nil {
				logger.Error("Invalid product ID %v", err)
			}

			state.AddItem(v2.NewCartItem(productId, item.Quantity))
		}
	})

	selector.AddReceive(removeFromCartChannel, func(c workflow.ReceiveChannel, _ bool) {
		var request v1.CartEvent
		c.Receive(ctx, &request)

		for _, item := range request.Items {
			productId, err := uuid.Parse(item.ProductId)
			if err != nil {
				logger.Error("Invalid product ID %v", err)
			}

			state.RemoveItem(v2.NewCartItem(productId, item.Quantity))
		}
	})

	selector.AddReceive(resetCartChannel, func(c workflow.ReceiveChannel, _ bool) {
		var customerId string
		c.Receive(ctx, &customerId)

		state.Reset()
	})

	for {
		selector.Select(ctx)
	}

	return nil
}
