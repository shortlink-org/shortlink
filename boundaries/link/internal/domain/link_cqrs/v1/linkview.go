package v1

// Link view
type LinkView struct {
	// URL
	url Url
	// Hash by URL + salt
	hash string
	// Describe of a link
	describe string
	// Metadata
	imageUrl string
	// Meta description
	metaDescription string
	// Meta keywords
	metaKeywords string

	// Create at
	createdAt Time
	// Update at
	updatedAt Time
}

// GetUrl returns the value of the url field.
func (m *LinkView) GetUrl() Url {
	return m.url
}

// GetHash returns the value of the hash field.
func (m *LinkView) GetHash() string {
	return m.hash
}

// GetDescribe returns the value of the describe field.
func (m *LinkView) GetDescribe() string {
	return m.describe
}

// GetImageUrl returns the value of the imageUrl field.
func (m *LinkView) GetImageUrl() string {
	return m.imageUrl
}

// GetMetaDescription returns the value of the metaDescription field.
func (m *LinkView) GetMetaDescription() string {
	return m.metaDescription
}

// GetMetaKeywords returns the value of the metaKeywords field.
func (m *LinkView) GetMetaKeywords() string {
	return m.metaKeywords
}

// GetCreatedAt returns the value of the createdAt field.
func (m *LinkView) GetCreatedAt() Time {
	return m.createdAt
}

// GetUpdatedAt returns the value of the updatedAt field.
func (m *LinkView) GetUpdatedAt() Time {
	return m.updatedAt
}
