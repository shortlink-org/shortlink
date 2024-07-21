package v1

import (
	"errors"
)

var (
	ErrInvalidId      = errors.New("invalid id: id is empty")
	ErrInvalidName    = errors.New("invalid name: name is empty")
	ErrInvalidPayload = errors.New("invalid payload: payload is empty")
)
