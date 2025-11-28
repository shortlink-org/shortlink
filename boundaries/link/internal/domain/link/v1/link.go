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

	// Allowed emails for private link access
	// Empty slice means public link, non-empty means private link
	allowedEmails []string
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

// GetAllowedEmails returns the list of allowed emails for private link access.
func (m *Link) GetAllowedEmails() []string {
	if m == nil {
		return nil
	}

	// Return a copy to prevent external modification
	result := make([]string, len(m.allowedEmails))
	copy(result, m.allowedEmails)
	return result
}

// IsPublic returns true if the link is public (no allowed emails).
func (m *Link) IsPublic() bool {
	if m == nil {
		return true
	}

	return len(m.allowedEmails) == 0
}

// CanBeViewedByEmail checks if the link can be viewed by the given email.
// For public links, always returns true.
// For private links, checks if the email is in the allowlist (after normalization).
func (m *Link) CanBeViewedByEmail(email string) bool {
	if m == nil {
		return false
	}

	// Public links are accessible to everyone
	if m.IsPublic() {
		return true
	}

	// Normalize email for comparison using Email VO
	viewerEmail, err := NewEmail(email)
	if err != nil || viewerEmail.IsEmpty() {
		return false
	}

	// Check if email is in allowlist
	for _, allowedEmailStr := range m.allowedEmails {
		allowedEmail, err := NewEmail(allowedEmailStr)
		if err != nil {
			continue // Skip invalid emails in allowlist
		}

		if viewerEmail == allowedEmail {
			return true
		}
	}

	return false
}

