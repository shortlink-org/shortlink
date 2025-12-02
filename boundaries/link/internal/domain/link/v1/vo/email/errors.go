package email

import (
	"fmt"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// MaxAllowlistSize is the maximum number of emails allowed in the allowlist.
const MaxAllowlistSize = 100

// ErrInvalidEmail indicates that an email address is invalid.
func ErrInvalidEmail(email string) *v1.LinkError {
	message := "invalid email"
	if email != "" {
		message = fmt.Sprintf("invalid email: %s", email)
	}

	return v1.NewLinkError(v1.CodeInvalidInput, message, nil)
}

// ErrAllowlistTooLarge indicates that the allowlist exceeds the maximum size.
func ErrAllowlistTooLarge(currentSize, maxSize int) *v1.LinkError {
	message := fmt.Sprintf("allowlist too large: %d emails (max: %d)", currentSize, maxSize)
	return v1.NewLinkError(v1.CodeInvalidInput, message, nil)
}

// ErrDuplicateEmail indicates that an email already exists in the allowlist.
func ErrDuplicateEmail(email string) *v1.LinkError {
	message := "duplicate email in allowlist"
	if email != "" {
		message = fmt.Sprintf("duplicate email in allowlist: %s", email)
	}

	return v1.NewLinkError(v1.CodeConflict, message, nil)
}

