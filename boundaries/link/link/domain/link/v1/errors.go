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
	return fmt.Sprintf("Not uniq link: %s", e.Link.GetUrl())
}

// CreateLinkError - failed create link
type CreateLinkError struct {
	Link *Link
}

func (e *CreateLinkError) Error() string {
	return fmt.Sprintf("Failed create link: %s", e.Link.GetUrl())
}

// PermissionDeniedError - permission denied
type PermissionDeniedError struct {
	Err error
}

func (e *PermissionDeniedError) Error() string {
	return "Permission denied"
}
