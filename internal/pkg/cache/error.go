package cache

import (
	"fmt"
)

// ErrorInitCache is an error returned when cache initialization fails.
type ErrorInitCache struct {
	err error
}

func (e *ErrorInitCache) Error() string {
	return fmt.Sprintf("error init cache: %s", e.err.Error())
}
