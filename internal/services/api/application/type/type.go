package api_type

import (
	"time"
)

// Config - base configuration for API
type Config struct { // nolint unused
	Port    int
	Timeout time.Duration
}
