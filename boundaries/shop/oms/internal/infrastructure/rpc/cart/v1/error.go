package v1

import (
	"errors"
	"fmt"
)

var ErrInvalidCustomerId = errors.New("invalid customer id")

type ParseItemError struct {
	Err error

	item string
}

func (e ParseItemError) Error() string {
	return fmt.Sprintf("failed to parse item: %s: %v", e.item, e.Err)
}
