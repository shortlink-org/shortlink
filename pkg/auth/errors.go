package auth

import (
	"fmt"
)

// ClientInitError represents an error that occurs during client initialization
type ClientInitError struct {
	Cause error
}

func (e *ClientInitError) Error() string {
	return fmt.Sprintf("failed to initialize auth client: %v", e.Cause)
}

// ConfigurationError represents an error in the configuration setup
type ConfigurationError struct {
	Cause error
}

func (e *ConfigurationError) Error() string {
	return fmt.Sprintf("auth configuration error: %v", e.Cause)
}
