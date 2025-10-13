package metadata

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/shortlink-org/go-sdk/logger"
	domain "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/parsers"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/screenshot"
	"github.com/shortlink-org/shortlink/pkg/pattern/saga"
)

type UC struct {
	// usecases
	parserUC     *parsers.UC
	screenshotUC *screenshot.UC

	// common
	log logger.Logger
}

func New(log logger.Logger, parsersUC *parsers.UC, screenshotUC *screenshot.UC) (*UC, error) {
	return &UC{
		// usecases
		parserUC:     parsersUC,
		screenshotUC: screenshotUC,

		// common
		log: log,
	}, nil
}

func errorHelper(ctx context.Context, log logger.Logger, errs []error) error {
	if len(errs) > 0 {
		attrs := make([]slog.Attr, 0, len(errs))
		for index, err := range errs {
			attrs = append(attrs, slog.Any(fmt.Sprintf("stack error: %d", index), err))
		}

		log.ErrorWithContext(ctx, "Error in saga", attrs...)

		return ErrSaga
	}

	return nil
}

// Add adds a metadata
func (uc *UC) Add(ctx context.Context, linkURL string) (*domain.Meta, error) {
	const (
		SAGA_NAME                    = "METADATA_ADD_CONTEXT"
		SAGA_STEP_ADD_META           = "SAGA_STEP_ADD_META"
		SAGA_STEP_ADD_SCREENSHOT     = "SAGA_STEP_ADD_SCREENSHOT"
		SAGA_STEP_GET_SCREENSHOT_URL = "SAGA_STEP_GET_SCREENSHOT_URL"
	)

	meta := &domain.Meta{}

	// create a new saga for set metadata
	sagaSetMetadata, errs := saga.New(SAGA_NAME, saga.SetLogger(uc.log)).
		WithContext(ctx).
		Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaSetMetadata.AddStep(SAGA_STEP_ADD_META).
		Then(func(ctx context.Context) error {
			var err error

			meta, err = uc.parserUC.Set(ctx, linkURL)
			if err != nil {
				return err
			}

			return nil
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaSetMetadata.AddStep(SAGA_STEP_ADD_SCREENSHOT).
		Then(func(ctx context.Context) error {
			err := uc.screenshotUC.Set(ctx, linkURL)
			if err != nil {
				return err
			}

			return nil
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaSetMetadata.AddStep(SAGA_STEP_GET_SCREENSHOT_URL).
		Needs(SAGA_STEP_ADD_SCREENSHOT, SAGA_STEP_ADD_META).
		Then(func(ctx context.Context) error {
			url, err := uc.screenshotUC.Get(ctx, linkURL)
			if err != nil {
				return err
			}

			meta.ImageUrl = url.String()

			return nil
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// run saga
	err := sagaSetMetadata.Play(nil)
	if err != nil {
		return nil, err
	}

	return meta, nil
}
