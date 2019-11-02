package logger

import (
	"github.com/sirupsen/logrus"
)

type logrusLogger struct { // nolint unused
	logger *logrus.Logger
}

func (log *logrusLogger) init(config Configuration) error {
	log.logger = logrus.New()

	// Logging =================================================================
	// Setup the logger backend using sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/sirupsen/logrus
	log.logger.Formatter = &logrus.JSONFormatter{}

	log.logger.SetReportCaller(true)

	log.setLogLevel(config.Level)

	return nil
}

func (log *logrusLogger) Info(msg string, fields ...Fields) {
	log.converter(fields...).Info(msg)
}

func (log *logrusLogger) Warn(msg string, fields ...Fields) {
	log.converter(fields...).Warn(msg)
}

func (log *logrusLogger) Error(msg string, fields ...Fields) {
	log.converter(fields...).Error(msg)
}

func (log *logrusLogger) Fatal(msg string, fields ...Fields) {
	log.converter(fields...).Fatal(msg)
}

func (log *logrusLogger) SetConfig(config Configuration) error {
	log.setLogLevel(config.Level)

	return nil
}

func (log *logrusLogger) Close() {}

func (log *logrusLogger) converter(fields ...Fields) *logrus.Entry {
	logrusFields := logrus.Fields{}

	for _, field := range fields {
		for k, v := range field {
			logrusFields[k] = v
		}
	}

	entryLog := log.logger.WithFields(logrusFields)
	return entryLog
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
