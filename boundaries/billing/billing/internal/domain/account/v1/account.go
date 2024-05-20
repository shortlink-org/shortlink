package v1

// Account is a billing account
type Account struct {
	// account id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// user id
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// tariff id
	TariffId string `protobuf:"bytes,3,opt,name=tariff_id,json=tariffId,proto3" json:"tariff_id,omitempty"`
}

// GetId returns the Id field value
func (m *Account) GetId() string {
	return m.Id
}

// GetUserId returns the UserId field value
func (m *Account) GetUserId() string {
	return m.UserId
}

// GetTariffId returns the TariffId field value
func (m *Account) GetTariffId() string {
	return m.TariffId
}
