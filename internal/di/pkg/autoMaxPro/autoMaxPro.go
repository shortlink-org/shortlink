package autoMaxPro

import (
	"fmt"

	"go.uber.org/automaxprocs/maxprocs"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

type AutoMaxPro *string

// New - Automatically set GOMAXPROCS to match Linux container CPU quota
func New(log logger.Logger) (AutoMaxPro, func(), error) {
	undo, err := maxprocs.Set(maxprocs.Logger(func(s string, args ...any) {
		log.Info(fmt.Sprintf(s, args...))
	}))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		undo()
	}

	return nil, cleanup, nil
}
