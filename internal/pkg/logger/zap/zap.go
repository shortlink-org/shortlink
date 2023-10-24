package zap

import (
	"time"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

type Logger struct {
	Logger *otelzap.Logger
}

func New(cfg config.Configuration) (*Logger, error) {
	log := &Logger{}
	logLevel := log.setLogLevel(cfg.Level)

	// To keep the example deterministic, disable timestamps in the output.
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoder(log.timeEncoder(cfg.TimeFormat)),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Wrap zap logger to extend Zap with API that accepts a context.Context.
	log.Logger = otelzap.New(zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(zapcore.AddSync(cfg.Writer)),
		logLevel,
	), zap.AddCaller(), zap.AddCallerSkip(1)))

	return log, nil
}

func (log *Logger) Close() error {
	err := log.Logger.Sync()
	return err
}

func (log *Logger) Get() any {
	return log.Logger
}

func (log *Logger) converter(fields ...field.Fields) []zap.Field {
	var zapFields []zap.Field

	for _, items := range fields {
		for k, v := range items {
			zapFields = append(zapFields, zap.Any(k, v))
		}
	}

	return zapFields
}

func (log *Logger) setLogLevel(logLevel int) zap.AtomicLevel {
	atom := zap.NewAtomicLevel()

	switch logLevel {
	case config.FATAL_LEVEL:
		atom.SetLevel(zapcore.FatalLevel)
	case config.ERROR_LEVEL:
		atom.SetLevel(zapcore.ErrorLevel)
	case config.WARN_LEVEL:
		atom.SetLevel(zapcore.WarnLevel)
	case config.INFO_LEVEL:
		atom.SetLevel(zapcore.InfoLevel)
	case config.DEBUG_LEVEL:
		atom.SetLevel(zapcore.DebugLevel)
	default:
		atom.SetLevel(zapcore.InfoLevel)
	}

	return atom
}

func (log *Logger) timeEncoder(format string) func(time.Time, zapcore.PrimitiveArrayEncoder) {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(format))
	}
}
