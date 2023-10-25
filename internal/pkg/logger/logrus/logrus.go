package logrus

import (
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otellogrus"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

type Logger struct {
	log *logrus.Logger
}

func New(cfg config.Configuration) (*Logger, error) {
	log := &Logger{
		log: logrus.New(),
	}

	// Logging =================================================================
	// Setup the logger backend using sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/sirupsen/logrus
	log.log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: cfg.TimeFormat,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
			logrus.FieldKeyFunc:  "caller",
		},
	}

	// Tracing
	log.log.AddHook(otellogrus.NewHook(otellogrus.WithLevels(
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	)))

	log.log.SetReportCaller(false) // TODO: https://github.com/sirupsen/logrus/pull/973
	log.log.SetOutput(cfg.Writer)
	log.setLogLevel(cfg.Level)

	return log, nil
}

func (log *Logger) Close() error {
	return nil
}

func (log *Logger) Get() any {
	return log.log
}

func (log *Logger) converter(fields ...field.Fields) *logrus.Entry {
	logrusFields := logrus.Fields{}

	for _, items := range fields {
		for k, v := range items {
			logrusFields[k] = v
		}
	}

	entryLog := log.log.WithFields(logrusFields)

	return entryLog
}

func (log *Logger) setLogLevel(logLevel int) {
	switch logLevel {
	case config.FATAL_LEVEL:
		log.log.SetLevel(logrus.FatalLevel)
	case config.ERROR_LEVEL:
		log.log.SetLevel(logrus.ErrorLevel)
	case config.WARN_LEVEL:
		log.log.SetLevel(logrus.WarnLevel)
	case config.INFO_LEVEL:
		log.log.SetLevel(logrus.InfoLevel)
	case config.DEBUG_LEVEL:
		log.log.SetLevel(logrus.DebugLevel)
	default:
		log.log.SetLevel(logrus.InfoLevel)
	}
}
