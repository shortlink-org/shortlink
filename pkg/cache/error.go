package cache

import (
	"fmt"
)

// InitCacheError is an error returned when cache initialization fails.
type InitCacheError struct {
	err error
}

func (e *InitCacheError) Error() string {
	return fmt.Sprintf("error init cache: %s", e.err.Error())
}
