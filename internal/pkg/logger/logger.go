/*
Logger
*/
package logger

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger/config"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/logrus"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/zap"
)

// New - return new an instance of logger
func New(loggerInstance int, cfg config.Configuration) (Logger, error) {
	var log Logger

	// Check config and set default values if needed
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}

	switch loggerInstance {
	case Zap:
		log, err = zap.New(cfg)
	case Logrus:
		log, err = logrus.New(cfg)
	default:
		return nil, ErrInvalidLoggerInstance
	}

	if err != nil {
		return nil, err
	}

	return log, nil
}
