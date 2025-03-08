package leveldb

import (
	"errors"
	"fmt"
)

// Error variables for wrapping underlying errors.
var (
	// ErrInvalidPath indicates an invalid LevelDB path.
	ErrInvalidPath = errors.New("invalid LevelDB path")
	// ErrDatabaseOpen indicates a failure to open the database.
	ErrDatabaseOpen = errors.New("failed to open LevelDB database")
	// ErrDatabaseClosed indicates operations on a closed database.
	ErrDatabaseClosed = errors.New("database is closed")
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
