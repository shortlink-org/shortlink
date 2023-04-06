package saga

import (
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

type StepState int

// Success chain: INIT -> WAIT -> RUN -> DONE
// Fail chain: INIT -> WAIT -> RUN -> REJECT -> FAIL or ROLLBACK
const (
	INIT StepState = iota + 1
	WAIT
	RUN
	DONE
	REJECT
	FAIL
	ROLLBACK
)

type Options struct {
	logger logger.Logger
}

type Option func(*Options)

func Logger(logger logger.Logger) Option {
	return func(args *Options) {
		args.logger = logger
	}
}
