package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"github.com/shortlink-org/shortlink/pkg/fsm"
)

// Error definitions
var (
	ErrInvalidOrderID     = errors.New("invalid order id")
	ErrInvalidProductID   = errors.New("invalid product id")
	ErrInvalidOrderStatus = errors.New("invalid order status")
)

// OrderStateBuilder is used to build a new OrderState
type OrderStateBuilder struct {
	orderState *OrderState
	errors     error
}

// NewOrderStateBuilder returns a new instance of OrderStateBuilder
func NewOrderStateBuilder(customerId uuid.UUID) *OrderStateBuilder {
	return &OrderStateBuilder{orderState: NewOrderState(customerId)}
}

// SetId sets the id of the order
func (b *OrderStateBuilder) SetId(id uuid.UUID) *OrderStateBuilder {
	if id == uuid.Nil {
		b.errors = errors.Join(b.errors, ErrInvalidOrderID)
		return b
	}

	b.orderState.id = id
	return b
}

// AddItem adds an item to the order
func (b *OrderStateBuilder) AddItem(productId uuid.UUID, quantity int32, price decimal.Decimal) *OrderStateBuilder {
	if productId == uuid.Nil {
		b.errors = errors.Join(b.errors, ErrInvalidProductID)
		return b
	}
	item := NewItem(productId, quantity, price)
	b.orderState.items = append(b.orderState.items, item)
	return b
}

// SetStatus sets the status of the order using a context for the FSM transition
func (b *OrderStateBuilder) SetStatus(ctx context.Context, status OrderStatus) *OrderStateBuilder {
	if status == OrderStatus_ORDER_STATUS_UNSPECIFIED {
		b.errors = errors.Join(b.errors, ErrInvalidOrderStatus)
		return b
	}

	// Map the desired status to the corresponding FSM event
	event, err := mapOrderStatusToEvent(status)
	if err != nil {
		b.errors = errors.Join(b.errors, err)
		return b
	}

	// Trigger the event with the provided context
	err = b.orderState.fsm.TriggerEvent(ctx, fsm.Event(event))
	if err != nil {
		b.errors = errors.Join(b.errors, err)
		return b
	}

	return b
}

// mapOrderStatusToEvent maps OrderStatus to corresponding FSM events
func mapOrderStatusToEvent(status OrderStatus) (string, error) {
	switch status {
	case OrderStatus_ORDER_STATUS_PROCESSING:
		// To reach "Processing", trigger the "Pending" event
		return OrderStatus_ORDER_STATUS_PENDING.String(), nil
	case OrderStatus_ORDER_STATUS_CANCELLED:
		// To reach "Cancelled", trigger the "Cancelled" event
		return OrderStatus_ORDER_STATUS_CANCELLED.String(), nil
	case OrderStatus_ORDER_STATUS_COMPLETED:
		// To reach "Completed", trigger the "Completed" event
		return OrderStatus_ORDER_STATUS_COMPLETED.String(), nil
	default:
		return "", fmt.Errorf("no event mapping for status '%s'", status)
	}
}

// Build finalizes the building process and returns the built OrderState
func (b *OrderStateBuilder) Build() (*OrderState, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	return b.orderState, nil
}
