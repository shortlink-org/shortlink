package rules

import (
	"github.com/shortlink-org/go-sdk/specification"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/vo/email"
)

// EmailDuplicateSpecification validates that an email is not a duplicate in the allowlist.
type EmailDuplicateSpecification struct {
	email string
}

// NewEmailDuplicateSpecification creates a new EmailDuplicateSpecification.
func NewEmailDuplicateSpecification(email string) specification.Specification[EmailValidationData] {
	return &EmailDuplicateSpecification{email: email}
}

// IsSatisfiedBy checks if the email is not a duplicate.
func (s *EmailDuplicateSpecification) IsSatisfiedBy(item *EmailValidationData) error {
	if item == nil {
		return email.ErrInvalidEmail(s.email)
	}

	normalized := NormalizeEmail(s.email)
	if normalized == "" {
		return email.ErrInvalidEmail(s.email)
	}

	if item.SeenEmails == nil {
		return nil // No seen emails, so no duplicate
	}

	if item.SeenEmails[normalized] {
		return email.ErrDuplicateEmail(s.email)
	}

	return nil
}
