package cart_workflow

import (
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"

	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/cart/workflow/model/cart/v1"
)

// Workflow is a Temporal workflow that manages the cart state.
func Workflow(ctx workflow.Context, customerId uuid.UUID) error {
	state := v2.NewCartState(customerId)

	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	addToCartChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_ADD.String())

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

	for {
		selector.Select(ctx)
	}

	return nil
}
