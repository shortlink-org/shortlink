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

// FlightRecorderConfig defines the configuration parameters for Go 1.25 trace.FlightRecorder.
// This configuration enables continuous, low-overhead tracing by maintaining a rolling buffer
// of the most recent execution trace data in memory for post-mortem analysis.
type FlightRecorderConfig struct {
	// Enabled determines whether the flight recorder should be activated.
	// When false, no trace data will be collected or buffered.
	Enabled bool
	
	// MinAge specifies the minimum duration of trace data to retain in the buffer.
	// This is a hint to the Go runtime; actual retention may vary based on system conditions.
	// Typical values range from 30 seconds to several minutes depending on application needs.
	MinAge time.Duration
	
	// MaxBytes defines the maximum buffer size in bytes for storing trace data.
	// This is also a hint to the Go runtime; actual memory usage may exceed this value.
	// The buffer operates as a ring buffer, discarding oldest data when the limit is approached.
	// Recommended values are 1-10MB for most applications.
	MaxBytes int64
}
