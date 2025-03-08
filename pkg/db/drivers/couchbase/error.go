package couchbase

import (
	"errors"
	"fmt"
)

var ErrCouchbaseConnect = errors.New("failed to connect to Couchbase cluster")

type StoreError struct {
	Op      string
	Err     error
	Details string
}

func (e *StoreError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("couchbase store error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("couchbase store error during %s: %v", e.Op, e.Err)
}

func (e *StoreError) Unwrap() error {
	return e.Err
}
