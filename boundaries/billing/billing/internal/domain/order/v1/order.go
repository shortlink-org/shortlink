package v1

// Order service
type Order struct {
	// order id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// user id
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// tariff id
	TariffId string `protobuf:"bytes,3,opt,name=tariff_id,json=tariffId,proto3" json:"tariff_id,omitempty"`
	// status order
	Status StatusOrder `protobuf:"varint,4,opt,name=status,proto3,enum=domain.billing.order.v1.StatusOrder" json:"status,omitempty"`
}
