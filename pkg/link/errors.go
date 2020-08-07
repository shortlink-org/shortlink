package link

import (
	"fmt"
)

// NotFoundError - not found link
type NotFoundError struct { // nolint unused
	Link *Link
	Err  error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found link: %s", e.Link.Url)
}

type NotUniqError struct { // nolint unused
	Link *Link
	Err  error
}

func (e *NotUniqError) Error() string {
	return fmt.Sprintf("Not uniq link: %s", e.Link.Url)
}
