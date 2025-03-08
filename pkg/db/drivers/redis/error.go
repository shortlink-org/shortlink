package redis

import (
	"errors"
	"fmt"
)

// Error variables for wrapping underlying errors.
var (
	// ErrInvalidURI indicates an invalid Redis URI.
	ErrInvalidURI = errors.New("invalid Redis URI")
	// ErrInvalidCredentials indicates invalid Redis credentials.
	ErrInvalidCredentials = errors.New("invalid Redis credentials")
	// ErrClientConnection indicates a failure to connect to Redis.
	ErrClientConnection = errors.New("failed to connect to Redis client")
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
		return fmt.Sprintf("redis error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("redis error during %s: %v", e.Op, e.Err)
}

// Unwrap allows errors.Is and errors.As to work with StoreError.
func (e *StoreError) Unwrap() error {
	return e.Err
}
