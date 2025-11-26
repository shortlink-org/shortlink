package link_cqrs

import "errors"

var (
	ErrCreateLink = errors.New("error create a new link")
	ErrNotFound   = errors.New("error not found link")
)
