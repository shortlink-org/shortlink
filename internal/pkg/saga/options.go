package saga

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

type Options struct {
	logger  logger.Logger
	limiter int
}

type Option func(*Options)

func Logger(logger logger.Logger) Option {
	return func(args *Options) {
		args.logger = logger
	}
}

func SetLimiter(limiter int) Option {
	return func(args *Options) {
		args.limiter = limiter
	}
}
