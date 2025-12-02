package rules

import (
	"github.com/shortlink-org/go-sdk/specification"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1/vo/email"
)

// AllowlistSizeSpecification validates that the allowlist does not exceed the maximum size.
type AllowlistSizeSpecification struct{}

// NewAllowlistSizeSpecification creates a new AllowlistSizeSpecification.
func NewAllowlistSizeSpecification() specification.Specification[[]string] {
	return &AllowlistSizeSpecification{}
}

// IsSatisfiedBy checks if the allowlist size is within limits.
func (s *AllowlistSizeSpecification) IsSatisfiedBy(item *[]string) error {
	if item == nil {
		return email.ErrAllowlistTooLarge(0, email.MaxAllowlistSize)
	}

	if len(*item) > email.MaxAllowlistSize {
		return email.ErrAllowlistTooLarge(len(*item), email.MaxAllowlistSize)
	}

	return nil
}
