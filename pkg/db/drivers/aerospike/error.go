package aerospike

import (
	"errors"
	"fmt"
)

// Error variables for wrapping underlying errors.
var (
	// ErrInvalidURI indicates an invalid Aerospike URI.
	ErrInvalidURI = errors.New("invalid Aerospike URI")
	// ErrInvalidPort indicates a failure during port conversion.
	ErrInvalidPort = errors.New("invalid port")
	// ErrClientConnection indicates a failure to connect to Aerospike.
	ErrClientConnection = errors.New("failed to connect to Aerospike client")
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
		return fmt.Sprintf("store error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("store error during %s: %v", e.Op, e.Err)
}

// Unwrap allows errors.Is and errors.As to work with StoreError.
func (e *StoreError) Unwrap() error {
	return e.Err
}
