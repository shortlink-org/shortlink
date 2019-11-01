package logger

import (
	"context"
	"errors"
)

// NewLogger - return new instance logger
func NewLogger(config Configuration, loggerInstance int) (Logger, error) { // nolint unused
	var log Logger

	switch loggerInstance {
	case Zap:
		log = &zapLogger{}
	case Logrus:
		return &logrusLogger{}, nil
	default:
		return nil, errors.New("Invalid logger instance")
	}

	// Init logger
	if err := log.init(config); err != nil {
		return nil, err
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
