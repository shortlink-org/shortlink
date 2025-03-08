package cockroachdb

import (
	"errors"
	"fmt"
)

var (
	ErrCockroachConfig  = errors.New("failed to parse cockroachdb config")
	ErrCockroachConnect = errors.New("failed to connect to cockroachdb")
)

type StoreError struct {
	Op      string
	Err     error
	Details string
}

func (e *StoreError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("cockroachdb store error during %s: %s: %v", e.Op, e.Details, e.Err)
	}

	return fmt.Sprintf("cockroachdb store error during %s: %v", e.Op, e.Err)
}

func (e *StoreError) Unwrap() error {
	return e.Err
}
