package v1

// Tariff is a domain model of tariff
type Tariff struct {
	// id of tariff
	id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// name of tariff
	name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// payload of tariff
	payload string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

// GetId returns the Id field value
func (m *Tariff) GetId() string {
	return m.id
}

// GetName returns the Name field value
func (m *Tariff) GetName() string {
	return m.name
}

// GetPayload returns the Payload field value
func (m *Tariff) GetPayload() string {
	return m.payload
}
