package http_server

import (
	"time"
)

// Config - base configuration for API
type Config struct {
	Port    int
	Timeout time.Duration
}
