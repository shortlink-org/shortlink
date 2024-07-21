package v1

import (
	"errors"

	"github.com/google/uuid"
)

// OrderBuilder is used to build a new Order
type OrderBuilder struct {
	order  *Order
	errors error
}

// NewOrderBuilder returns a new instance of OrderBuilder
func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{order: &Order{}}
}

// SetId sets the id of the order
func (b *OrderBuilder) SetId(id uuid.UUID) *OrderBuilder {
	if id == uuid.Nil {
		b.errors = errors.Join(b.errors, ErrInvalidOrderId)
		return b
	}

	b.order.id = id

	return b
}

// SetUserId sets the userId of the order
func (b *OrderBuilder) SetUserId(userId uuid.UUID) *OrderBuilder {
	if userId == uuid.Nil {
		b.errors = errors.Join(b.errors, ErrInvalidOrderUserId)
		return b
	}

	b.order.userId = userId

	return b
}

// SetTariffId sets the tariffId of the order
func (b *OrderBuilder) SetTariffId(tariffId uuid.UUID) *OrderBuilder {
	if tariffId == uuid.Nil {
		b.errors = errors.Join(b.errors, ErrInvalidOrderTariffId)

		return b
	}

	b.order.tariffId = tariffId

	return b
}

// SetStatus sets the status of the order
func (b *OrderBuilder) SetStatus(status StatusOrder) *OrderBuilder {
	// Check for a valid status value if necessary
	if _, ok := StatusOrder_name[int32(status)]; !ok {
		b.errors = errors.Join(b.errors, ErrInvalidOrderStatus)

		return b
	}

	b.order.status = status

	return b
}

// Build finalizes the building process and returns the built Order
func (b *OrderBuilder) Build() (*Order, error) {
	if b.errors != nil {
		return nil, b.errors
	}

	// Generate a new id if it is not set
	if b.order.id == uuid.Nil {
		var err error
		b.order.id, err = uuid.NewV7()
		if err != nil {
			return nil, err
		}
	}

	return b.order, nil
}
