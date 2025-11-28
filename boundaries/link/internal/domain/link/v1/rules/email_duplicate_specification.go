package rules

import (
	"fmt"

	"github.com/shortlink-org/go-sdk/specification"
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
		return fmt.Errorf("invalid email: %s", s.email)
	}

	normalized := NormalizeEmail(s.email)
	if normalized == "" {
		return fmt.Errorf("invalid email: %s", s.email)
	}

	if item.SeenEmails == nil {
		return nil // No seen emails, so no duplicate
	}

	if item.SeenEmails[normalized] {
		return fmt.Errorf("duplicate email in allowlist: %s", s.email)
	}

	return nil
}

