package v1

import (
	"errors"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

// SetStatus sets the status of the order
func (b *OrderStateBuilder) SetStatus(status OrderStatus) *OrderStateBuilder {
	if status == OrderStatus_ORDER_STATUS_UNSPECIFIED {
		b.errors = errors.Join(b.errors, ErrInvalidOrderStatus)
		return b
	}
	b.orderState.fsm.SetState(status.String())
	return b
}

// Build finalizes the building process and returns the built OrderState
func (b *OrderStateBuilder) Build() (*OrderState, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	return b.orderState, nil
}
