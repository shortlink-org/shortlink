package saga

import (
	"github.com/shortlink-org/go-sdk/logger"
)

type Options struct {
	log     logger.Logger
	limiter int
}

type Option func(*Options)

func SetLogger(log logger.Logger) Option {
	return func(args *Options) {
		args.log = log
	}
}

func SetLimiter(limiter int) Option {
	return func(args *Options) {
		args.limiter = limiter
	}
}
