package v1

// Tariff is a domain model of tariff
type Tariff struct {
	// id of tariff
	id string
	// name of tariff
	name string
	// payload of tariff
	payload string
}

// GetId returns the id field value
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
