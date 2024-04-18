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

// NotFoundByHashError - not found link by hash
type NotFoundByHashError struct {
	Hash string
}

func (e *NotFoundByHashError) Error() string {
	return fmt.Sprintf("Not found link by hash: %s", e.Hash)
}

// NotUniqError - not uniq link
type NotUniqError struct {
	Link *Link
}

func (e *NotUniqError) Error() string {
	link := e.Link.GetUrl()

	return fmt.Sprintf("Not uniq link: %s", link.String())
}

// CreateLinkError - failed create link
type CreateLinkError struct {
	Link *Link
}

func (e *CreateLinkError) Error() string {
	link := e.Link.GetUrl()

	return fmt.Sprintf("Failed create link: %s", link.String())
}

// PermissionDeniedError - permission denied
type PermissionDeniedError struct {
	Err error
}

func (e *PermissionDeniedError) Error() string {
	return "Permission denied"
}
