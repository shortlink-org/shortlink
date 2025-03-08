package badger

import (
	"errors"
	"fmt"
)

// Error variables for common Badger errors.
var (
	// ErrBadgerOpen indicates a failure to open the Badger database.
	ErrBadgerOpen = errors.New("failed to open Badger DB")
	// ErrBadgerClose indicates a failure to close the Badger database.
	ErrBadgerClose = errors.New("failed to close Badger DB")
)

// StoreError is a custom error type for Badger store operations with additional details.
type StoreError struct {
	Op      string
	Err     error
	Details string
}

// Error implements the error interface.
func (e *StoreError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("badger store error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("badger store error during %s: %v", e.Op, e.Err)
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As.
func (e *StoreError) Unwrap() error {
	return e.Err
}
