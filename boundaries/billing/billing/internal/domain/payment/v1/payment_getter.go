package v1

import (
	"github.com/google/uuid"
)

// GetId returns the id field value
func (m *Payment) GetId() uuid.UUID {
	return m.id
}

// GetName returns the name field value
func (m *Payment) GetName() string {
	return m.name
}

// GetStatus returns the status field value
func (m *Payment) GetStatus() StatusPayment {
	return m.status
}

// GetUserId returns the userId field value
func (m *Payment) GetUserId() uuid.UUID {
	return m.userId
}

// GetAmount returns the amount field value
func (m *Payment) GetAmount() int64 {
	return m.amount
}
