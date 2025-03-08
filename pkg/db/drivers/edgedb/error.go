package edgedb

import (
	"errors"
	"fmt"
)

var (
	ErrConnect = errors.New("failed to connect to EdgeDB")
	ErrClose   = errors.New("failed to close EdgeDB connection")
)

type StoreError struct {
	Op      string
	Err     error
	Details string
}

func (e *StoreError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("edgedb store error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("edgedb store error during %s: %v", e.Op, e.Err)
}

func (e *StoreError) Unwrap() error {
	return e.Err
}
