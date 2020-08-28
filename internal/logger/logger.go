/*
Logger
*/

package logger

import (
	"errors"
	"os"
	"time"
)

// NewLogger - return new an instance of logger
func NewLogger(loggerInstance int, config Configuration) (Logger, error) {
	var log Logger

	switch loggerInstance {
	case Zap:
		log = &zapLogger{}
	case Logrus:
		log = &logrusLogger{}
	default:
		return nil, errors.New("Invalid logger instance")
	}

	// Init logger
	validateConfig(&config)
	if err := log.init(config); err != nil {
		return nil, err
	}

	return log, nil
}

func validateConfig(config *Configuration) { // nolint unused
	if config.Writer == nil {
		config.Writer = os.Stdout
	}

	if config.TimeFormat == "" {
		config.TimeFormat = time.RFC3339Nano
	}
}
