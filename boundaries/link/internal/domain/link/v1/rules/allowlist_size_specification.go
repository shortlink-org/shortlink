package rules

import (
	"fmt"

	"github.com/shortlink-org/go-sdk/specification"
)

// MaxAllowlistSize is the maximum number of emails allowed in the allowlist.
const MaxAllowlistSize = 100

// AllowlistSizeSpecification validates that the allowlist does not exceed the maximum size.
type AllowlistSizeSpecification struct{}

// NewAllowlistSizeSpecification creates a new AllowlistSizeSpecification.
func NewAllowlistSizeSpecification() specification.Specification[[]string] {
	return &AllowlistSizeSpecification{}
}

// IsSatisfiedBy checks if the allowlist size is within limits.
func (s *AllowlistSizeSpecification) IsSatisfiedBy(item *[]string) error {
	if item == nil {
		return fmt.Errorf("allowlist too large: 0 emails (max: %d)", MaxAllowlistSize)
	}

	if len(*item) > MaxAllowlistSize {
		return fmt.Errorf("allowlist too large: %d emails (max: %d)", len(*item), MaxAllowlistSize)
	}

	return nil
}

