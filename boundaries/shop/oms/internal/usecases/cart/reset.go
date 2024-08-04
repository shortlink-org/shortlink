package cart

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/cart/v1"
)

// Reset resets the cart.
func (uc *UC) Reset(ctx context.Context, customerId uuid.UUID) error {
	workflowId := fmt.Sprintf("cart-%s", customerId)

	err := uc.temporalClient.SignalWorkflow(ctx, workflowId, "", v1.Event_EVENT_RESET.String(), customerId)
	if err != nil {
		return err
	}

	return nil
}
