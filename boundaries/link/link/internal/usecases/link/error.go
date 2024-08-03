package link

import (
	"context"
	"errors"
	"fmt"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/logger/field"
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
		errList := field.Fields{}
		for index := range errs {
			errList[fmt.Sprintf("stack error: %d", index)] = errs[index]
		}

		log.ErrorWithContext(ctx, "Error create a new link", errList)

		return ErrCreateLink
	}

	return nil
}
