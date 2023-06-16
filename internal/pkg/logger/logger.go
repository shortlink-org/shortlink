/*
Logger
*/
package logger

import (
	"errors"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/logrus"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/zap"
)

// NewLogger - return new an instance of logger
func NewLogger(loggerInstance int, config config.Configuration) (Logger, error) {
	var log Logger

	// Check config and set default values if needed
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	switch loggerInstance {
	case Zap:
		log, err = zap.New(config)
	case Logrus:
		log, err = logrus.New(config)
	default:
		return nil, errors.New("invalid logger instance")
	}

	if err != nil {
		return nil, err
	}

	return log, nil
}
