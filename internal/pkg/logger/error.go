package logger

import (
	"errors"
)

// ErrInvalidLoggerInstance is an error when logger instance is invalid
var ErrInvalidLoggerInstance = errors.New("invalid logger instance")
