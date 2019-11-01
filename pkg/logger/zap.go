package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type zapLogger struct { // nolint unused
	logger *zap.Logger
}

func (log *zapLogger) init(config Configuration) error {
	logLevel := log.setLogLevel(config.Level)

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zap.NewProductionEncoderConfig()

	log.logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		logLevel,
	))

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

func (log *zapLogger) SetConfig(config Configuration) error {
	var err error
	logLevel := log.setLogLevel(config.Level)

	cfg := zap.Config{
		Level: logLevel,
	}
	if log.logger, err = cfg.Build(); err != nil {
		return err
	}

	return nil
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
			case time.Duration:
				zapFields = append(zapFields, zap.Duration(k, v))
			default:
				return nil, fmt.Errorf("Don't support type field: %T", v)
			}
		}
	}

	return zapFields, nil
}

func (log *zapLogger) setLogLevel(logLevel int) zap.AtomicLevel {
	atom := zap.NewAtomicLevel()

	switch logLevel {
	case LOG_LEVEL_DEBUG:
		atom.SetLevel(zap.DebugLevel)
	case LOG_LEVEL_INFO:
		atom.SetLevel(zap.InfoLevel)
	default:
		atom.SetLevel(zap.InfoLevel)
	}

	return atom
}
