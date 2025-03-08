package sqlite

import (
	"errors"
	"fmt"
)

// Error variables for wrapping underlying errors.
var (
	// ErrInvalidPath indicates an invalid SQLite database path.
	ErrInvalidPath = errors.New("invalid SQLite database path")
	// ErrClientConnection indicates a failure to connect to SQLite.
	ErrClientConnection = errors.New("failed to connect to SQLite database")
	// ErrInvalidConfiguration indicates invalid SQLite configuration.
	ErrInvalidConfiguration = errors.New("invalid SQLite configuration")
)

// StoreError is a custom error type for Store operations with added details.
type StoreError struct {
	Op      string
	Err     error
	Details string
}

// Error implements the error interface.
func (e *StoreError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("sqlite error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("sqlite error during %s: %v", e.Op, e.Err)
}

// Unwrap allows errors.Is and errors.As to work with StoreError.
func (e *StoreError) Unwrap() error {
	return e.Err
}
