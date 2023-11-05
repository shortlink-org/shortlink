package ctx

import (
	"context"
)

func New() (context.Context, func(), error) {
	ctx, cancel := context.WithCancel(context.Background())

	cb := func() {
		cancel()
	}

	return ctx, cb, nil
}
