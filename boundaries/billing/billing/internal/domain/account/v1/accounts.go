package v1

// Accounts is a list of billing accounts
type Accounts struct {
	// accounts
	list []*Account `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

// GetList returns the list field value
func (m *Accounts) GetList() []*Account {
	return m.list
}
