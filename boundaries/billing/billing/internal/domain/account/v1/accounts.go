package v1

// Accounts is a list of billing accounts
type Accounts struct {
	// accounts
	List []*Account `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}
