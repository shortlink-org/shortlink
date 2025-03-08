package postgres

import (
	"errors"
	"fmt"
)

// Error variables for wrapping underlying errors.
var (
	// ErrInvalidDSN indicates an invalid PostgreSQL connection string.
	ErrInvalidDSN = errors.New("invalid PostgreSQL DSN")
	// ErrClientConnection indicates a failure to connect to PostgreSQL.
	ErrClientConnection = errors.New("failed to connect to PostgreSQL server")
	// ErrInvalidDatabase indicates an invalid database name.
	ErrInvalidDatabase = errors.New("invalid database name")
	// ErrInvalidCredentials indicates invalid authentication credentials.
	ErrInvalidCredentials = errors.New("invalid PostgreSQL credentials")
	// ErrInvalidSchema indicates an invalid schema name.
	ErrInvalidSchema = errors.New("invalid schema name")
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

// PingConnectionError - error ping connection
type PingConnectionError struct {
	Err error
}

func (e *PingConnectionError) Error() string {
	return "failed to ping the database: " + e.Err.Error()
}
