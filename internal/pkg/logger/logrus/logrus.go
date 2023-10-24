package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

type Logger struct {
	logger *logrus.Logger
}

func New(cfg config.Configuration) (*Logger, error) {
	log := &Logger{
		logger: logrus.New(),
	}

	// Logging =================================================================
	// Setup the logger backend using sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/sirupsen/logrus
	log.logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: cfg.TimeFormat,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
			logrus.FieldKeyFunc:  "caller",
		},
	}

	// Tracing
	log.logger.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	)))

	log.logger.SetReportCaller(false) // TODO: https://github.com/sirupsen/logrus/pull/973
	log.logger.SetOutput(cfg.Writer)
	log.setLogLevel(cfg.Level)

	return log, nil
}

func (log *Logger) Close() error {
	return nil
}

func (log *Logger) Get() any {
	return log.logger
}

func (log *Logger) converter(fields ...field.Fields) *logrus.Entry {
	logrusFields := logrus.Fields{}

	for _, items := range fields {
		for k, v := range items {
			logrusFields[k] = v
		}
	}

	entryLog := log.logger.WithFields(logrusFields)

	return entryLog
}

func (log *Logger) setLogLevel(logLevel int) {
	switch logLevel {
	case config.FATAL_LEVEL:
		log.logger.SetLevel(logrus.FatalLevel)
	case config.ERROR_LEVEL:
		log.logger.SetLevel(logrus.ErrorLevel)
	case config.WARN_LEVEL:
		log.logger.SetLevel(logrus.WarnLevel)
	case config.INFO_LEVEL:
		log.logger.SetLevel(logrus.InfoLevel)
	case config.DEBUG_LEVEL:
		log.logger.SetLevel(logrus.DebugLevel)
	default:
		log.logger.SetLevel(logrus.InfoLevel)
	}
}
