package link

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/shortlink-org/go-sdk/logger"
)

var ErrCreateLink = errors.New("error create a new link")

// NotFoundByHash is an error when the link is not found by hash
type NotFoundByHash struct {
	Hash string
}

func (e NotFoundByHash) Error() string {
	return "link not found by hash: " + e.Hash
}

// errorHelper is a helper function to log errors
func errorHelper(ctx context.Context, log logger.Logger, errs []error) error {
	if len(errs) > 0 {
		attrs := make([]slog.Attr, 0, len(errs))
		for index, err := range errs {
			attrs = append(attrs, slog.Any(fmt.Sprintf("stack error: %d", index), err))
		}

		log.ErrorWithContext(ctx, "Error create a new link", attrs...)

		return ErrCreateLink
	}

	return nil
}
