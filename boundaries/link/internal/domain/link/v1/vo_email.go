package v1

import (
	"strings"
)

// Email is a Value Object representing a normalized email address.
// It ensures that email addresses are always stored and compared in a normalized form.
type Email string

// NewEmail creates a new Email Value Object from a string.
// The email is automatically normalized (lowercase, trimmed).
// Returns an error if the email is empty after normalization.
func NewEmail(email string) (Email, error) {
	normalized := strings.ToLower(strings.TrimSpace(email))
	if normalized == "" {
		return "", ErrInvalidEmail(email)
	}

	return Email(normalized), nil
}

// String returns the normalized email address as a string.
func (e Email) String() string {
	return string(e)
}

// Value returns the normalized email address as a string.
// Alias for String() for consistency with other Value Objects.
func (e Email) Value() string {
	return e.String()
}

// Equals checks if two Email Value Objects are equal.
// Comparison is done on normalized values.
func (e Email) Equals(other Email) bool {
	return e == other
}

// IsEmpty checks if the email is empty.
func (e Email) IsEmpty() bool {
	return e == ""
}

