package link

import (
	"errors"
)

var ErrCreateLink = errors.New("error create a new link")

// NotFoundByHash is an error when the link is not found by hash
type NotFoundByHash struct {
	Hash string
}

func (e NotFoundByHash) Error() string {
	return "link not found by hash: " + e.Hash
}
