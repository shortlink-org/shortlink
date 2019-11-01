package logger

import (
	"github.com/sirupsen/logrus"
)

type logrusLogger struct { // nolint unused
	logger *logrus.Logger
}

type logrusLogEntry struct { // nolint unused
	entry *logrus.Entry
}

func (log *logrusLogger) init(config Configuration) error {
	log.logger = logrus.New()

	// Logging =================================================================
	// Setup the logger backend using sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/sirupsen/logrus
	log.logger.Formatter = new(logrus.JSONFormatter)

	log.setLogLevel(config.Level)

	return nil
}

func (log *logrusLogger) Info(msg string, fields ...Fields) {
	log.converter(fields...).logger.Info(msg)
}

func (log *logrusLogger) Warn(msg string, fields ...Fields) {
	log.logger.Warn(msg)
}

func (log *logrusLogger) Error(msg string, fields ...Fields) {
	log.logger.Error(msg)
}

func (log *logrusLogger) Fatal(msg string, fields ...Fields) {
	log.logger.Fatal(msg)
}

func (log *logrusLogger) SetConfig(config Configuration) error {
	log.setLogLevel(config.Level)

	return nil
}

func (log *logrusLogger) Close() {}

func (log *logrusLogger) converter(fields ...Fields) *logrusLogger {
	var logrusFields logrus.Fields

	for _, field := range fields {
		for k, v := range field {
			logrusFields[k] = v
		}
	}

	log.logger.WithFields(logrusFields)

	return log
}

func (log *logrusLogger) setLogLevel(logLevel int) {
	switch logLevel {
	case LOG_LEVEL_DEBUG:
		log.logger.SetLevel(logrus.DebugLevel)
	case LOG_LEVEL_INFO:
		log.logger.SetLevel(logrus.InfoLevel)
	default:
		log.logger.SetLevel(logrus.InfoLevel)
	}
}
