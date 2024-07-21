package v1

// Tariffs is a domain model of tariffs
type Tariffs struct {
	// List of tariffs
	List []*Tariff `json:"list,omitempty" protobuf:"bytes,1,rep,name=list,proto3"`
}

// GetList returns the List field value
func (m *Tariffs) GetList() []*Tariff {
	return m.List
}
