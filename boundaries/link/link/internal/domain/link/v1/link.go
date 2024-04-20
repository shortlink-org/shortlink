package v1

import (
	"net/url"
	"time"
)

// Link is a domain model.
type Link struct {
	// URL
	url url.URL
	// Hash by URL + salt
	hash string
	// Describe of a link
	describe string

	// Create at
	createdAt time.Time
	// Update at
	updatedAt time.Time
}

// GetUrl returns the value of the url field.
func (m *Link) GetUrl() url.URL {
	return m.url
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
func (m *Link) GetCreatedAt() time.Time {
	return m.createdAt
}

// GetUpdatedAt returns the value of the updatedAt field.
func (m *Link) GetUpdatedAt() time.Time {
	return m.updatedAt
}
