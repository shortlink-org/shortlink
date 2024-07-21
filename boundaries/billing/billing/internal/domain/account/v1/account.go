package v1

import (
	"github.com/google/uuid"
)

// Account is a billing account
type Account struct {
	// account id
	id uuid.UUID
	// user id
	userId uuid.UUID
	// tariff id
	tariffId uuid.UUID
}

// GetId returns the id field value
func (m *Account) GetId() uuid.UUID {
	return m.id
}

// GetUserId returns the userId field value
func (m *Account) GetUserId() uuid.UUID {
	return m.userId
}

// GetTariffId returns the tariffId field value
func (m *Account) GetTariffId() uuid.UUID {
	return m.tariffId
}
