package cart

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/client"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/queue/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/cart/dto"
	cart_workflow "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/cart/workflow"
)

// Add adds an item to the cart.
func (uc *UC) Add(ctx context.Context, in *v1.CartState) error {
	workflowId := fmt.Sprintf("cart-%s", in.GetCustomerId().String())

	_, err := uc.temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        workflowId,
		TaskQueue: v2.CART_TASK_QUEUE,
	}, cart_workflow.Workflow, in.GetCustomerId())
	if err != nil {
		return err
	}

	request := dto.CartStateToCartEvent(in, v1.Event_EVENT_ADD)
	err = uc.temporalClient.SignalWorkflow(ctx, workflowId, "", v1.Event_EVENT_ADD.String(), request)
	if err != nil {
		return err
	}

	return nil
}
