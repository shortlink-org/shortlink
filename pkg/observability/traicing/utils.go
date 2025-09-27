/*
Utility functions for FlightRecorder integration
*/
package traicing

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/shortlink-org/go-sdk/logger"
)

var (
	globalFlightRecorder *FlightRecorder
	globalMutex          sync.RWMutex
)

// SetGlobalFlightRecorder sets the global flight recorder instance
func SetGlobalFlightRecorder(fr *FlightRecorder) {
	globalMutex.Lock()
	defer globalMutex.Unlock()
	globalFlightRecorder = fr
}

// GetGlobalFlightRecorder returns the global flight recorder instance
func GetGlobalFlightRecorder() *FlightRecorder {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	return globalFlightRecorder
}

// SaveTraceOnError saves trace data when an error occurs  
func SaveTraceOnError(err error, log logger.Logger) {
	if err == nil {
		return
	}

	fr := GetGlobalFlightRecorder()
	if fr == nil || !fr.IsRunning() {
		return
	}

	filename := fmt.Sprintf("error_trace_%d.out", time.Now().Unix())
	if saveErr := fr.WriteToFile(filename); saveErr != nil {
		log.Error("Failed to save error trace", 
			slog.Any("save_err", saveErr),
			slog.String("filename", filename),
			slog.Any("error", err))
	} else {
		log.Info("Error trace saved", 
			slog.String("filename", filename),
			slog.Any("error", err))
	}
}

// SaveTraceOnSignal saves trace data when receiving a specific signal (like SIGUSR1)
func SaveTraceOnSignal(signal string, log logger.Logger) {
	fr := GetGlobalFlightRecorder()
	if fr == nil || !fr.IsRunning() {
		if log != nil {
			log.Warn("Flight recorder not available for signal trace", slog.String("signal", signal))
		}
		return
	}

	filename := fmt.Sprintf("signal_%s_trace_%d.out", signal, time.Now().Unix())
	if err := fr.WriteToFile(filename); err != nil {
		if log != nil {
			log.Error("Failed to save signal trace", 
				slog.Any("err", err),
				slog.String("filename", filename),
				slog.String("signal", signal))
		}
	} else {
		if log != nil {
			log.Info("Signal trace saved", 
				slog.String("filename", filename),
				slog.String("signal", signal))
		}
	}
}

// SaveTraceWithContext saves trace data with additional context information
func SaveTraceWithContext(ctx context.Context, reason string, metadata map[string]interface{}, log logger.Logger) {
	fr := GetGlobalFlightRecorder()
	if fr == nil || !fr.IsRunning() {
		if log != nil {
			log.Warn("Flight recorder not available", slog.String("reason", reason))
		}
		return
	}

	filename := fmt.Sprintf("%s_trace_%d.out", reason, time.Now().Unix())
	if err := fr.WriteToFile(filename); err != nil {
		if log != nil {
			log.Error("Failed to save trace with context", 
				slog.Any("err", err),
				slog.String("filename", filename),
				slog.String("reason", reason),
				slog.Any("metadata", metadata))
		}
	} else {
		// Also save metadata to a separate file
		metadataFile := fmt.Sprintf("%s_metadata_%d.txt", reason, time.Now().Unix())
		if file, err := os.Create(metadataFile); err == nil {
			defer file.Close()
			
			fmt.Fprintf(file, "Trace Reason: %s\n", reason)
			fmt.Fprintf(file, "Timestamp: %s\n", time.Now().Format(time.RFC3339))
			fmt.Fprintf(file, "Stack Trace:\n%s\n", debug.Stack())
			
			if metadata != nil {
				fmt.Fprintf(file, "\nMetadata:\n")
				for k, v := range metadata {
					fmt.Fprintf(file, "  %s: %v\n", k, v)
				}
			}
		}

		if log != nil {
			log.Info("Trace with context saved", 
				slog.String("filename", filename),
				slog.String("metadata_file", metadataFile),
				slog.String("reason", reason),
				slog.Any("metadata", metadata))
		}
	}
}

// RecorderMiddleware provides middleware functions for easy integration
type RecorderMiddleware struct {
	log logger.Logger
}

// NewRecorderMiddleware creates a new middleware instance
func NewRecorderMiddleware(log logger.Logger) *RecorderMiddleware {
	return &RecorderMiddleware{log: log}
}

// WrapWithPanicRecovery wraps a function to automatically save trace on panic
func (rm *RecorderMiddleware) WrapWithPanicRecovery(fn func()) func() {
	return func() {
		defer func() {
			if r := recover(); r != nil {
				fr := GetGlobalFlightRecorder()
				if fr != nil {
					fr.SaveTraceOnPanic()
				}
				// Re-panic to preserve original behavior
				panic(r)
			}
		}()
		fn()
	}
}

// WrapWithErrorTracking wraps a function to automatically save trace on error
func (rm *RecorderMiddleware) WrapWithErrorTracking(fn func() error) func() error {
	return func() error {
		err := fn()
		if err != nil {
			SaveTraceOnError(err, rm.log)
		}
		return err
	}
}

// HealthCheck returns the status of the flight recorder
func HealthCheck() map[string]interface{} {
	fr := GetGlobalFlightRecorder()
	if fr == nil {
		return map[string]interface{}{
			"flight_recorder": "not_initialized",
			"status":          "disabled",
		}
	}

	config := fr.GetConfig()
	return map[string]interface{}{
		"flight_recorder": "initialized",
		"status":          map[string]interface{}{
			"running":   fr.IsRunning(),
			"enabled":   config.Enabled,
			"min_age":   config.MinAge.String(),
			"max_bytes": config.MaxBytes,
		},
	}
}