package logger

import (
	"fmt"
	"go.uber.org/zap"
)

type zapLogger struct { // nolint unused
	logger *zap.Logger
}

func (log *zapLogger) Init() error {
	var err error
	if log.logger, err = zap.NewProduction(); err != nil {
		return err
	}
	return nil
}

func (log *zapLogger) Info(msg string, fields ...Fields) {
	var err error
	zapFields, err := log.converter(fields...)
	if err != nil {
		log.Error(err.Error(), nil)
		return
	}

	log.logger.Info(msg, zapFields...)
}

func (log *zapLogger) Warn(msg string, fields ...Fields) {
	log.logger.Warn(msg)
}

func (log *zapLogger) Error(msg string, fields ...Fields) {
	log.logger.Error(msg)
}

func (log *zapLogger) Fatal(msg string, fields ...Fields) {
	log.logger.Fatal(msg)
}

func (log *zapLogger) Close() {
	_ = log.logger.Sync() // nolint errcheck
}

func (log *zapLogger) converter(fields ...Fields) ([]zap.Field, error) {
	var zapFields []zap.Field

	for _, field := range fields {
		for k, v := range field {
			switch v := v.(type) {
			case string:
				zapFields = append(zapFields, zap.String(k, v))
			case int:
				zapFields = append(zapFields, zap.Int(k, v))
			default:
				return nil, fmt.Errorf("Don't support type field: %T", v)
			}
		}
	}

	return zapFields, nil
}
