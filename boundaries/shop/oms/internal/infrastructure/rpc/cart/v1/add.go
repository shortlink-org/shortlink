package v1

import (
	"context"

	"go.temporal.io/sdk/client"
	"google.golang.org/protobuf/types/known/emptypb"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/cart"
)

func (c *CartRPC) Add(ctx context.Context, in *AddRequest) (*emptypb.Empty, error) {
	workflow, err := c.client.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		// ID: "CART-1722217371",
		TaskQueue: "CART_TASK_QUEUE",
	}, cart.Workflow, &v1.CartState{})
	if err != nil {
		return nil, err
	}

	// wait for the workflow to complete
	if err := workflow.Get(ctx, nil); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
