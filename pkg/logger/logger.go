package logger

import (
	"context"
	"log/slog"

	"github.com/shortlink-org/shortlink/pkg/logger/tracer"
)

type SlogLogger struct {
	logger *slog.Logger
}

func New(cfg Configuration) (*SlogLogger, error) {
	// Validate config and set defaults
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	// JSON handler with source and formatted timestamp (from record, not time.Now)
	handler := slog.NewJSONHandler(cfg.Writer, &slog.HandlerOptions{
		Level:     convertLevel(cfg.Level),
		AddSource: true,
		ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.TimeKey && attr.Value.Kind() == slog.KindTime {
				return slog.String(slog.TimeKey, attr.Value.Time().Format(cfg.TimeFormat))
			}

			return attr
		},
	})

	return &SlogLogger{logger: slog.New(handler)}, nil
}

func (log *SlogLogger) Close() error {
	// slog has nothing to close
	return nil
}

// convertLevel converts our int level to slog.Level.
func convertLevel(level int) slog.Level {
	switch level {
	case ERROR_LEVEL:
		return slog.LevelError
	case WARN_LEVEL:
		return slog.LevelWarn
	case INFO_LEVEL:
		return slog.LevelInfo
	case DEBUG_LEVEL:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

// levelString maps slog.Level to a severity string for tracer.
func levelString(level slog.Level) string {
	switch {
	case level <= slog.LevelError:
		return "ERROR"
	case level == slog.LevelWarn:
		return "WARN"
	case level == slog.LevelInfo:
		return "INFO"
	default:
		return "DEBUG"
	}
}

// logWithContext enriches fields with trace correlation (if ctx carries a span) and logs.
func (log *SlogLogger) logWithContext(ctx context.Context, level slog.Level, msg string, fields ...any) {
	// Enrich with OTel span event + traceID/spanID if a span exists.
	if ctx != nil && ctx != context.Background() {
		enriched, err := tracer.NewTraceFromContext(ctx, levelString(level), msg, nil, fields...)
		if err == nil {
			fields = enriched
		} else {
			// Log enrichment failure once; avoid recursion (log directly).
			log.logger.Log(ctx, slog.LevelError,
				"OpenTelemetry trace enrichment failed",
				"err", err.Error(),
			)
		}
	}

	log.logger.Log(ctx, level, msg, fields...)
}