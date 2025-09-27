package traicing

import (
	"time"
)

// Config - config
type Config struct {
	ServiceName         string
	ServiceVersion      string
	URI                 string
	PyroscopeURI        string
	FlightRecorder      *FlightRecorderConfig
}

// FlightRecorderConfig configures the Go 1.25 trace.FlightRecorder
type FlightRecorderConfig struct {
	Enabled  bool
	MinAge   time.Duration
	MaxBytes int64
}
