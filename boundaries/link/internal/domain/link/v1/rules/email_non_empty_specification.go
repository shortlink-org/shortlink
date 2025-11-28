package rules

import (
	"fmt"

	"github.com/shortlink-org/go-sdk/specification"
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
		return fmt.Errorf("invalid email: %s", s.email)
	}

	normalized := NormalizeEmail(item.Email)
	if normalized == "" {
		return fmt.Errorf("invalid email: %s", s.email)
	}

	return nil
}

