package batch

import (
	"context"
)

// NewSync creates a batch that runs in the background like New,
// but it blocks until the passed context is cancelled and all
// pending items have been processed. The first error returned
// by the user-supplied callback (if any) is propagated as the
// returned error.
func NewSync[T any](
	ctx context.Context,
	callback func([]*Item[T]) error,
	opts ...Option[T],
) (*Batch[T], error) {
	// Re-use the asynchronous constructor.
	batch, errChan := New(ctx, callback, opts...)

	var firstErr error
	for err := range errChan { // errChan closes when ctx.Done() is observed.
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}

	return batch, firstErr
}
