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

type NotUniqError struct {
	Link *Link
	Err  error
}

func (e *NotUniqError) Error() string {
	return fmt.Sprintf("Not uniq link: %s", e.Link.GetUrl())
}

type ErrCreateLink struct {
	Err  error
	Link *Link
}

func (e *ErrCreateLink) Error() string {
	return fmt.Sprintf("Failed create link: %s", e.Link.GetUrl())
}
