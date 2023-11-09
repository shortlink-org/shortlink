package saga

import (
	"context"
)

type Error string

const ContextErrorKey = Error("saga-error")

func WithError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, ContextErrorKey, err)
}

func GetError(ctx context.Context) error {
	err := ctx.Value(ContextErrorKey)
	if err == nil {
		return nil
	}

	return err.(error)
}
