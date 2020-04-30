package api_type

import (
	"time"
)

const (
	METHOD_ADD    = iota // nolint unused
	METHOD_GET           // nolint unused
	METHOD_LIST          // nolint unused
	METHOD_UPDATE        // nolint unused
	METHOD_DELETE        // nolint unused
)

// Config - base configuration for API
type Config struct { // nolint unused
	Port    int
	Timeout time.Duration
}
