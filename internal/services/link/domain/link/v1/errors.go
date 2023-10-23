package v1

import (
	"fmt"
)

// NotFoundError - not found link
type NotFoundError struct {
	Link *Link
	Err  error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found link: %s", e.Link.GetHash())
}

type NotUniqError struct { //nolint:decorder
	Link *Link
	Err  error
}

func (e *NotUniqError) Error() string {
	return fmt.Sprintf("Not uniq link: %s", e.Link.GetUrl())
}
