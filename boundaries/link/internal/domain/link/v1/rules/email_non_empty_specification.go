package rules

import (
	"github.com/shortlink-org/go-sdk/specification"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/vo/email"
)

// EmailNonEmptySpecification validates that an email is not empty after normalization.
type EmailNonEmptySpecification struct {
	email string
}

// NewEmailNonEmptySpecification creates a new EmailNonEmptySpecification.
func NewEmailNonEmptySpecification(email string) specification.Specification[EmailValidationData] {
	return &EmailNonEmptySpecification{email: email}
}

// IsSatisfiedBy checks if the email is not empty after normalization.
func (s *EmailNonEmptySpecification) IsSatisfiedBy(item *EmailValidationData) error {
	if item == nil {
		return email.ErrInvalidEmail(s.email)
	}

	normalized := NormalizeEmail(item.Email)
	if normalized == "" {
		return email.ErrInvalidEmail(s.email)
	}

	return nil
}
