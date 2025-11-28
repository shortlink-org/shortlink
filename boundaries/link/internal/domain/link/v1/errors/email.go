package errors

import v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"

// MaxAllowlistSize is the maximum number of emails allowed in the allowlist.
const MaxAllowlistSize = 100

// ErrInvalidEmail indicates that an email address is invalid.
func ErrInvalidEmail(email string) *v1.LinkError {
	return v1.ErrInvalidEmail(email)
}

// ErrAllowlistTooLarge indicates that the allowlist exceeds the maximum size.
func ErrAllowlistTooLarge(currentSize, maxSize int) *v1.LinkError {
	return v1.ErrAllowlistTooLarge(currentSize, maxSize)
}

// ErrDuplicateEmail indicates that an email already exists in the allowlist.
func ErrDuplicateEmail(email string) *v1.LinkError {
	return v1.ErrDuplicateEmail(email)
}

