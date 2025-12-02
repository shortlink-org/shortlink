package email

import "fmt"

// MaxAllowlistSize is the maximum number of emails allowed in the allowlist.
const MaxAllowlistSize = 100

// InvalidEmailError indicates that an email address is invalid.
type InvalidEmailError struct {
	Email string
}

func (e *InvalidEmailError) Error() string {
	if e == nil || e.Email == "" {
		return "invalid email"
	}

	return fmt.Sprintf("invalid email: %s", e.Email)
}

// ErrInvalidEmail creates InvalidEmailError without depending on link domain.
func ErrInvalidEmail(email string) error {
	return &InvalidEmailError{Email: email}
}

// AllowlistTooLargeError indicates that the allowlist exceeds the maximum size.
type AllowlistTooLargeError struct {
	CurrentSize int
	MaxSize     int
}

func (e *AllowlistTooLargeError) Error() string {
	if e == nil {
		return "allowlist too large"
	}

	return fmt.Sprintf("allowlist too large: %d emails (max: %d)", e.CurrentSize, e.MaxSize)
}

// ErrAllowlistTooLarge creates AllowlistTooLargeError with context.
func ErrAllowlistTooLarge(currentSize, maxSize int) error {
	return &AllowlistTooLargeError{
		CurrentSize: currentSize,
		MaxSize:     maxSize,
	}
}

// DuplicateEmailError indicates that an email already exists in the allowlist.
type DuplicateEmailError struct {
	Email string
}

func (e *DuplicateEmailError) Error() string {
	if e == nil || e.Email == "" {
		return "duplicate email in allowlist"
	}

	return fmt.Sprintf("duplicate email in allowlist: %s", e.Email)
}

// ErrDuplicateEmail creates DuplicateEmailError for the provided address.
func ErrDuplicateEmail(email string) error {
	return &DuplicateEmailError{Email: email}
}
