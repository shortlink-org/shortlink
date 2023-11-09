package session

import "errors"

var ErrSessionNotFound = errors.New("session not found")
var ErrMetadataNotFound = errors.New("metadata not found")
var ErrUserIDNotFound = errors.New("user-id not found")
