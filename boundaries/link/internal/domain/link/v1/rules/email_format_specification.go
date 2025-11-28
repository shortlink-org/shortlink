package rules

import (
	"fmt"
	"net/mail"

	"github.com/shortlink-org/go-sdk/specification"
)

// EmailFormatSpecification validates that an email address has valid RFC 5322 format.
type EmailFormatSpecification struct {
	email string
}

// NewEmailFormatSpecification creates a new EmailFormatSpecification.
func NewEmailFormatSpecification(email string) specification.Specification[EmailValidationData] {
	return &EmailFormatSpecification{email: email}
}

// IsSatisfiedBy checks if the email has valid format.
func (s *EmailFormatSpecification) IsSatisfiedBy(item *EmailValidationData) error {
	if item == nil {
		return fmt.Errorf("invalid email: %s", s.email)
	}

	normalized := NormalizeEmail(item.Email)
	if normalized == "" {
		return fmt.Errorf("invalid email: %s", s.email)
	}

	_, err := mail.ParseAddress(normalized)
	if err != nil {
		return fmt.Errorf("invalid email: %s", s.email)
	}

	return nil
}

