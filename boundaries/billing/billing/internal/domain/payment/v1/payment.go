package v1

// Payment - information about payment
type Payment struct {
	// ID payment
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name payment
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Status payment
	Status StatusPayment `protobuf:"varint,3,opt,name=status,proto3,enum=domain.billing.payment.v1.StatusPayment" json:"status,omitempty"`
	// User ID
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Amount payment
	Amount int64 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
}
