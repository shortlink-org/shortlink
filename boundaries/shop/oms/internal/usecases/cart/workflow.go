package cart

import (
	"go.temporal.io/sdk/workflow"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
)

// Workflow is a Temporal workflow that manages the cart state.
func Workflow(ctx workflow.Context, state *v1.CartState) error {
	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	err := workflow.SetQueryHandler(ctx, "getCart", func(input []byte) (*v1.CartState, error) {
		return state, nil
	})
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
		return err
	}

	// addToCartChannel := workflow.GetSignalChannel(ctx, v1.ADD_TO_CART.String())
	// removeFromCartChannel := workflow.GetSignalChannel(ctx, v1.REMOVE_FROM_CART.String())
	//
	// selector := workflow.NewSelector(ctx)

	// for {
	// 	selector.AddReceive(addToCartChannel, func(c workflow.ReceiveChannel, _ bool) {
	// 		var signal interface{}
	// 		c.Receive(ctx, &signal)
	//
	// 		var message v1.AddToCartSignal
	// 		err := mapstructure.Decode(signal, &message)
	// 		if err != nil {
	// 			logger.Error("Invalid signal type %v", err)
	// 			return
	// 		}
	//
	// 		state.AddItem(message.Item)
	// 	})
	//
	// 	selector.AddReceive(removeFromCartChannel, func(c workflow.ReceiveChannel, _ bool) {
	// 		var signal interface{}
	// 		c.Receive(ctx, &signal)
	//
	// 		var message v1.RemoveFromCartSignal
	// 		err := mapstructure.Decode(signal, &message)
	// 		if err != nil {
	// 			logger.Error("Invalid signal type %v", err)
	// 			return
	// 		}
	//
	// 		state.RemoveItem(message.Item)
	// 	})
	//
	// 	// This ensures the workflow yields control periodically
	// 	selector.Select(ctx)
	//
	// 	// Add short sleep to yield control more explicitly
	// 	err = workflow.Sleep(ctx, 1*time.Second)
	// 	if err != nil {
	// 		logger.Error("Workflow sleep interrupted.", "Error", err)
	// 		return err
	// 	}
	// }

	return nil
}
