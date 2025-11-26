package v1

import (
	v1 "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1"
)

// NotFoundError - not found link
type NotFoundError struct {
	Link v1.Link
	Hash string
}

func (e *NotFoundError) Error() string {
	if e.Hash != "" {
		return "Not found link: " + e.Hash
	}

	if hash := e.Link.GetHash(); hash != "" {
		return "Not found link: " + hash
	}

	return "Not found link"
}

// NotFoundByHashError - not found link by hash
type NotFoundByHashError struct {
	Hash string
}

func (e *NotFoundByHashError) Error() string {
	return "Not found link by hash: " + e.Hash
}

// CreateLinkError - create link error
type CreateLinkError struct {
	Link v1.Link
}

func (e *CreateLinkError) Error() string {
	return "Create link error: " + e.Link.GetHash()
}
