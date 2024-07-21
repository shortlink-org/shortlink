package v1

import (
	"github.com/google/uuid"
)

// Order service
type Order struct {
	// order id
	id uuid.UUID
	// user id
	userId uuid.UUID
	// tariff id
	tariffId uuid.UUID
	// status order
	status StatusOrder
}

// GetId returns the id field value
func (m *Order) GetId() uuid.UUID {
	return m.id
}

// GetUserId returns the userId field value
func (m *Order) GetUserId() uuid.UUID {
	return m.userId
}

// GetTariffId returns the tariffId field value
func (m *Order) GetTariffId() uuid.UUID {
	return m.tariffId
}

// GetStatus returns the status field value
func (m *Order) GetStatus() StatusOrder {
	return m.status
}
