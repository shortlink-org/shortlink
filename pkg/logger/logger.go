package logger

import (
	"context"
)

func NewLogger(config Configuration, loggerInstance int) (Logger, error) { // nolint unused
	var log Logger

	switch loggerInstance {
	case Zap:
		log = &zapLogger{}
	case Logrus:
		return &logrusLogger{}, nil
	}

	return log, nil
}

// WithLogger set logger
func WithLogger(ctx context.Context, logger Logger) context.Context { //nolint unused
	return context.WithValue(ctx, keyLogger, logger)
}

// GetLogger return logger
func GetLogger(ctx context.Context) Logger { //nolint unused
	return ctx.Value(keyLogger).(Logger)
}
