package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func New(config config.Configuration) (*LogrusLogger, error) {
	log := &LogrusLogger{
		logger: logrus.New(),
	}

	// Logging =================================================================
	// Setup the logger backend using sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/sirupsen/logrus
	log.logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: config.TimeFormat,
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
	log.logger.SetOutput(config.Writer)
	log.setLogLevel(config.Level)

	return log, nil
}

func (log *LogrusLogger) Close() error {
	return nil
}

func (log *LogrusLogger) Get() any {
	return log.logger
}

func (log *LogrusLogger) converter(fields ...field.Fields) *logrus.Entry {
	logrusFields := logrus.Fields{}

	for _, field := range fields {
		for k, v := range field {
			logrusFields[k] = v
		}
	}

	entryLog := log.logger.WithFields(logrusFields)

	return entryLog
}

func (log *LogrusLogger) setLogLevel(logLevel int) {
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
