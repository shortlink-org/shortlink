package v1

// Payments - list payments
type Payments struct {
	// list of payments
	List []*Payment `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}
