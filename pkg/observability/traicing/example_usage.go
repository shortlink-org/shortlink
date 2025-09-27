/*
Example usage of Go 1.25 FlightRecorder integration

This file demonstrates how to use the FlightRecorder for perfect tracing
in your application. Remove this file if not needed in production.
*/
package traicing

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shortlink-org/go-sdk/logger"
)

// ExampleFlightRecorderUsage demonstrates how to use the FlightRecorder
func ExampleFlightRecorderUsage(log *logger.SlogLogger) {
	// Example 1: Basic FlightRecorder setup with the provided configuration
	cfg := FlightRecorderConfig{
		Enabled:  true,
		MinAge:   5 * time.Second,
		MaxBytes: 3 << 20, // 3MB
	}

	fr, err := NewFlightRecorder(cfg, log)
	if err != nil {
		log.Error("Failed to create flight recorder", slog.Any("err", err))
		return
	}

	if err := fr.Start(); err != nil {
		log.Error("Failed to start flight recorder", slog.Any("err", err))
		return
	}
	defer fr.Stop()

	// Example 2: Save trace on specific conditions
	someBusinessLogic := func() error {
		// Simulate some work that might fail
		time.Sleep(100 * time.Millisecond)
		
		// Simulate an error condition
		if time.Now().Unix()%2 == 0 {
			return fmt.Errorf("simulated business logic error")
		}
		return nil
	}

	// Example 3: Using middleware for automatic error tracing
	middleware := NewRecorderMiddleware(log)
	
	wrappedFunction := middleware.WrapWithErrorTracking(someBusinessLogic)
	if err := wrappedFunction(); err != nil {
		log.Info("Business logic returned error, trace automatically saved", slog.Any("err", err))
	}

	// Example 4: Manual trace saving with context
	SaveTraceWithContext(context.Background(), "manual_trace", map[string]interface{}{
		"user_id":    "12345",
		"request_id": "req-789",
		"endpoint":   "/api/v1/links",
	}, log)

	// Example 5: Signal-based trace saving
	setupSignalHandling(log)

	log.Info("FlightRecorder example completed")
}

// setupSignalHandling sets up signal handling for trace saving
func setupSignalHandling(log *logger.SlogLogger) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1, syscall.SIGUSR2)

	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGUSR1:
				SaveTraceOnSignal("USR1", log)
			case syscall.SIGUSR2:
				SaveTraceOnSignal("USR2", log)
			}
		}
	}()

	log.Info("Signal handlers set up. Send SIGUSR1 or SIGUSR2 to save traces")
}

// ExamplePanicRecovery demonstrates panic recovery with trace saving
func ExamplePanicRecovery(log *logger.SlogLogger) {
	middleware := NewRecorderMiddleware(log)
	
	riskyFunction := func() {
		// Simulate some work
		time.Sleep(50 * time.Millisecond)
		
		// Simulate a panic
		panic("something went terribly wrong!")
	}

	// Wrap the risky function to automatically save trace on panic
	safeFunction := middleware.WrapWithPanicRecovery(riskyFunction)
	
	// This will recover from panic and save trace
	defer func() {
		if r := recover(); r != nil {
			log.Error("Recovered from panic, trace saved", slog.Any("panic", r))
		}
	}()
	
	safeFunction()
}

// ExampleHealthCheck demonstrates how to check FlightRecorder status
func ExampleHealthCheck(log *logger.SlogLogger) {
	status := HealthCheck()
	log.Info("FlightRecorder health check", slog.Any("status", status))
}

// ExampleEnvironmentVariables shows the environment variables that can be used
func ExampleEnvironmentVariables() {
	fmt.Println(`
FlightRecorder Environment Variables:

FLIGHT_RECORDER_ENABLED=true               # Enable/disable flight recorder
FLIGHT_RECORDER_MIN_AGE=5s                 # Minimum age of trace data to retain
FLIGHT_RECORDER_MAX_BYTES=3145728          # Maximum buffer size (3MB)

Example usage:
export FLIGHT_RECORDER_ENABLED=true
export FLIGHT_RECORDER_MIN_AGE=10s
export FLIGHT_RECORDER_MAX_BYTES=5242880   # 5MB

You can also configure it programmatically:
config := traicing.Config{
    ServiceName:    "my-service",
    ServiceVersion: "1.0.0",
    URI:            "localhost:4317",
    FlightRecorder: &traicing.FlightRecorderConfig{
        Enabled:  true,
        MinAge:   5 * time.Second,
        MaxBytes: 3 << 20, // 3MB
    },
}
	`)
}