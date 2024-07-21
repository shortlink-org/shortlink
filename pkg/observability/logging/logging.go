package logging

import (
	"context"
	"sync"

	otellog "go.opentelemetry.io/otel/log"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

// ScopeRecords represents the records for a single instrumentation scope.
type ScopeRecords struct {
	// Name is the name of the instrumentation scope.
	Name string
	// Version is the version of the instrumentation scope.
	Version string
	// SchemaURL of the telemetry emitted by the scope.
	SchemaURL string

	// Records are the log records this instrumentation scope recorded.
	Records []otellog.Record
}

type Logger struct {
	otellog.Logger

	mu          sync.Mutex
	scopeRecord *ScopeRecords
}

// New - creates a new Logger.
func New(ctx context.Context, log logger.Logger) *Logger {
	return &Logger{
		scopeRecord: &ScopeRecords{
			Name:      "shortlink",
			Version:   "0.1.0",
			SchemaURL: "",
			Records:   make([]otellog.Record, 0),
		},
	}
}

// Enabled indicates whether a specific record should be stored.
//
//nolint:gocritic // This is a temporary implementation.
func (l *Logger) Enabled(ctx context.Context, record otellog.Record) bool {
	return true
}

// Emit stores the log record.
//
//nolint:gocritic // This is a temporary implementation.
func (l *Logger) Emit(_ context.Context, record otellog.Record) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.scopeRecord.Records = append(l.scopeRecord.Records, record)
}
