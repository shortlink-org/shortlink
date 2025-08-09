package v1

import (
	"fmt"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
)

// NotFoundError - not found link
type NotFoundError struct {
	Link v1.Link
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

// CreateLinkError - create link error
type CreateLinkError struct {
	Link v1.Link
}

func (e *CreateLinkError) Error() string {
	return fmt.Sprintf("Create link error: %s", e.Link.GetHash())
}
