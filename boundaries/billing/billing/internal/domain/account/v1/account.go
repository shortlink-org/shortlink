package v1

import (
	"github.com/google/uuid"
)

// Account is a billing account
type Account struct {
	// account id
	id uuid.UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// user id
	userId uuid.UUID `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// tariff id
	tariffId uuid.UUID `protobuf:"bytes,3,opt,name=tariff_id,json=tariffId,proto3" json:"tariff_id,omitempty"`
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
