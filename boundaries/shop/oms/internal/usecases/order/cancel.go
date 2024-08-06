package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	domain "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
)

func (uc *UC) Cancel(ctx context.Context, orderId uuid.UUID) error {
	workflowId := fmt.Sprintf("order-%s", orderId.String())

	err := uc.temporalClient.SignalWorkflow(ctx, workflowId, "", domain.Event_EVENT_CANCEL.String(), nil)
	if err != nil {
		return err
	}

	return nil
}
