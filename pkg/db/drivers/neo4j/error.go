package neo4j

import (
	"errors"
	"fmt"
)

// Error variables for wrapping underlying errors.
var (
	// ErrInvalidURI indicates an invalid Neo4j URI.
	ErrInvalidURI = errors.New("invalid Neo4j URI")
	// ErrClientConnection indicates a failure to connect to Neo4j.
	ErrClientConnection = errors.New("failed to connect to Neo4j server")
	// ErrInvalidCredentials indicates invalid authentication credentials.
	ErrInvalidCredentials = errors.New("invalid Neo4j credentials")
	// ErrCypherQuery indicates an invalid Cypher query.
	ErrCypherQuery = errors.New("invalid Cypher query")
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
