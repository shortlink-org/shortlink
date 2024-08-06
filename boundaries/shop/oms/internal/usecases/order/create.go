package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/queue/v1"
	order_workflow "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/order/workflow"
)

func (uc *UC) Create(ctx context.Context, orderId uuid.UUID, customerId uuid.UUID, in v2.Items) error {
	workflowId := fmt.Sprintf("order-%s", orderId.String())

	_, err := uc.temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        workflowId,
		TaskQueue: v1.ORDER_TASK_QUEUE,
	}, order_workflow.Workflow, orderId, customerId, in)
	if err != nil {
		return err
	}

	return nil
}
