package v1

import (
	"fmt"
)

// NotFoundError - not found link
type NotFoundError struct {
	Link *Link
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Not found link: %s", e.Link.GetHash())
}

type NotUniqError struct {
	Link *Link
}

func (e *NotUniqError) Error() string {
	return fmt.Sprintf("Not uniq link: %s", e.Link.GetUrl())
}

type CreateLinkError struct {
	Link *Link
}

func (e *CreateLinkError) Error() string {
	return fmt.Sprintf("Failed create link: %s", e.Link.GetUrl())
}

type PermissionDeniedError struct {
	Err error
}

func (e *PermissionDeniedError) Error() string {
	return fmt.Sprintf("Permission denied: %w", e.Err)
}
