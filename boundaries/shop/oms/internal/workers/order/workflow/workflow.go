package order_workflow

import (
	"github.com/google/uuid"
	"go.temporal.io/sdk/workflow"

	v2 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/workers/order/workflow/dto"
)

// Workflow is a Temporal workflow that manages the order state.
func Workflow(ctx workflow.Context, orderId, customerId uuid.UUID, in v2.Items) error {
	state := v2.NewOrderState(customerId)

	// Set up query handler for getting order state
	err := workflow.SetQueryHandler(ctx, v2.Event_EVENT_GET.String(), func() (*v3.OrderState, error) {
		return dto.OrderStateToDomain(state), nil
	})
	if err != nil {
		return err
	}

	// https://docs.temporal.io/docs/concepts/workflows/#workflows-have-options
	logger := workflow.GetLogger(ctx)

	cancelOrderChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_CANCEL.String())
	completeOrderChannel := workflow.GetSignalChannel(ctx, v2.Event_EVENT_COMPLETE.String())

	selector := workflow.NewSelector(ctx)

	selector.AddReceive(cancelOrderChannel, func(c workflow.ReceiveChannel, _ bool) {
		logger.Info("Order cancellation started.")
	})

	selector.AddReceive(completeOrderChannel, func(c workflow.ReceiveChannel, _ bool) {
		logger.Info("Order completion started.")
	})

	for {
		selector.Select(ctx)
	}
}
