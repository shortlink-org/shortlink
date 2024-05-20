package v1

// Order list
type Orders struct {
	// order list
	List []*Order `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}
