package dgraph

import (
	"errors"
	"fmt"
)

var (
	ErrDgraphClient  = errors.New("failed to create Dgraph gRPC client")
	ErrDgraphMigrate = errors.New("failed to migrate Dgraph schema")
)

type StoreError struct {
	Op      string
	Err     error
	Details string
}

func (e *StoreError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("dgraph store error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("dgraph store error during %s: %v", e.Op, e.Err)
}

func (e *StoreError) Unwrap() error {
	return e.Err
}
