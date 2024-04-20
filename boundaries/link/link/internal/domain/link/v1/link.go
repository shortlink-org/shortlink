package v1

// Link is a domain model.
type Link struct {
	// URL
	url Url
	// Hash by URL + salt
	hash string
	// Describe of a link
	describe string

	// Create at
	createdAt Time
	// Update at
	updatedAt Time
}

// GetUrl returns the value of the url field.
func (m *Link) GetUrl() *Url {
	return &m.url
}

// GetHash returns the value of the hash field.
func (m *Link) GetHash() string {
	return m.hash
}

// GetDescribe returns the value of the described field.
func (m *Link) GetDescribe() string {
	return m.describe
}

// GetCreatedAt returns the value of the createdAt field.
func (m *Link) GetCreatedAt() Time {
	return m.createdAt
}

// GetUpdatedAt returns the value of the updatedAt field.
func (m *Link) GetUpdatedAt() Time {
	return m.updatedAt
}
