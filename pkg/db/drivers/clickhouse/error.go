package clickhouse

import (
	"errors"
	"fmt"
)

// Common error variables for Clickhouse store operations.
var (
	// ErrPing indicates that a Ping to the Clickhouse database failed.
	ErrPing = errors.New("failed to ping Clickhouse database")
	// ErrClose indicates that closing the Clickhouse connection failed.
	ErrClose = errors.New("failed to close Clickhouse connection")
)

// StoreError is a custom error type for Clickhouse store operations with additional details.
type StoreError struct {
	Op      string
	Err     error
	Details string
}

// Error implements the error interface.
func (e *StoreError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("clickhouse store error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("clickhouse store error during %s: %v", e.Op, e.Err)
}

// Unwrap returns the underlying error, enabling errors.Is and errors.As to work with StoreError.
func (e *StoreError) Unwrap() error {
	return e.Err
}
