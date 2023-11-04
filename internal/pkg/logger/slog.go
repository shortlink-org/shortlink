package logger

import (
	"context"
	"log/slog"
)

type StructLogger struct {
	Logger
	slog.Handler
}

func NewStructLogger(log Logger) (*slog.Logger, error) {
	structLogger := &StructLogger{
		Logger: log,
	}

	return slog.New(structLogger), nil
}

func (sl *StructLogger) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

//nolint:gocritic // This is a wrapper for the logger
func (sl *StructLogger) Handle(ctx context.Context, record slog.Record) error {
	switch record.Level {
	case slog.LevelDebug:
		sl.DebugWithContext(ctx, record.Message)
	case slog.LevelInfo:
		sl.InfoWithContext(ctx, record.Message)
	case slog.LevelWarn:
		sl.WarnWithContext(ctx, record.Message)
	case slog.LevelError:
		sl.ErrorWithContext(ctx, record.Message)
	default:
		sl.DebugWithContext(ctx, record.Message)
	}

	return nil
}

func (sl *StructLogger) WithAttrs(_ []slog.Attr) slog.Handler {
	return sl.Handler
}

func (sl *StructLogger) WithGroup(_ string) slog.Handler {
	return sl.Handler
}
