package v1

// Order service
type Order struct {
	// order id
	id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// user id
	userId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// tariff id
	tariffId string `protobuf:"bytes,3,opt,name=tariff_id,json=tariffId,proto3" json:"tariff_id,omitempty"`
	// status order
	status StatusOrder `protobuf:"varint,4,opt,name=status,proto3,enum=domain.billing.order.v1.StatusOrder" json:"status,omitempty"`
}

// GetId returns the id field value
func (m *Order) GetId() string {
	return m.id
}

// GetUserId returns the userId field value
func (m *Order) GetUserId() string {
	return m.userId
}

// GetTariffId returns the tariffId field value
func (m *Order) GetTariffId() string {
	return m.tariffId
}

// GetStatus returns the status field value
func (m *Order) GetStatus() StatusOrder {
	return m.status
}
