package v1

// Payments - list payments
type Payments struct {
	// list of payments
	list []*Payment `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

// GetList
func (o *Payments) GetList() []*Payment {
	return o.list
}
