package order

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	v1 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/domain/order/v1"
	v3 "github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/infrastructure/rpc/order/v1/model/v1"
	"github.com/shortlink-org/shortlink/boundaries/shop/oms/internal/usecases/order/dto"
)

func (uc *UC) Get(ctx context.Context, orderId uuid.UUID) (*v1.OrderState, error) {
	workflowId := fmt.Sprintf("order-%s", orderId.String())

	resp, err := uc.temporalClient.QueryWorkflow(ctx, workflowId, "", v1.Event_EVENT_GET.String(), nil)
	if err != nil {
		return nil, err
	}

	var orderState v3.OrderState
	err = resp.Get(&orderState)
	if err != nil {
		return nil, err
	}

	state, err := dto.OrderStateToDomain(&orderState)
	if err != nil {
		return nil, err
	}

	return state, nil
}
