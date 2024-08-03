package cart

import (
	"context"

	"go.temporal.io/sdk/client"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/cart/workflow"
)

// Add adds an item to the cart.
func (uc *UC) Add(ctx context.Context, in *v1.CartState) error {
	resp, err := uc.temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		// ID: "CART-1722217371",
		// TODO: use constants
		TaskQueue: "CART_TASK_QUEUE",
	}, workflow.Workflow, in)
	if err != nil {
		return err
	}

	// wait for the workflow to complete
	err = resp.Get(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
