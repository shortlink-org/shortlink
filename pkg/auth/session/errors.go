package session

import "errors"

var (
	ErrSessionNotFound  = errors.New("session not found")
	ErrMetadataNotFound = errors.New("metadata not found")
	ErrUserIDNotFound   = errors.New("user-id not found")
)
