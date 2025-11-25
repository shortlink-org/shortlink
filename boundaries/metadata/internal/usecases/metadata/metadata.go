package metadata

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/shortlink-org/go-sdk/cqrs/bus"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/saga"
	v1 "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/parsers"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/screenshot"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UC struct {
	// usecases
	parserUC     *parsers.UC
	screenshotUC *screenshot.UC

	// infrastructure
	eventBus *bus.EventBus

	// common
	log logger.Logger
}

func New(log logger.Logger, parsersUC *parsers.UC, screenshotUC *screenshot.UC, eventBus *bus.EventBus) (*UC, error) {
	return &UC{
		// usecases
		parserUC:     parsersUC,
		screenshotUC: screenshotUC,

		// infrastructure
		eventBus: eventBus,

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
func (uc *UC) Add(ctx context.Context, linkURL string) (*v1.Meta, error) {
	const (
		SAGA_NAME                    = "METADATA_ADD_CONTEXT"
		SAGA_STEP_ADD_META           = "SAGA_STEP_ADD_META"
		SAGA_STEP_ADD_SCREENSHOT     = "SAGA_STEP_ADD_SCREENSHOT"
		SAGA_STEP_GET_SCREENSHOT_URL = "SAGA_STEP_GET_SCREENSHOT_URL"
		SAGA_STEP_UPDATE_META        = "SAGA_STEP_UPDATE_META"
	)

	meta := &v1.Meta{}

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
			// Try to get screenshot URL, but don't fail if screenshot is not available yet
			url, err := uc.screenshotUC.Get(ctx, linkURL)
			if err != nil {
				// Log warning but continue without screenshot URL
				uc.log.Warn("Failed to get screenshot URL, continuing without it",
					slog.String("error", err.Error()),
					slog.String("url", linkURL),
				)
				return nil // Continue saga execution even if screenshot URL is not available
			}

			meta.ImageUrl = url.String()

			return nil
		}).Build()
	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaSetMetadata.AddStep(SAGA_STEP_UPDATE_META).
		Needs(SAGA_STEP_GET_SCREENSHOT_URL).
		Then(func(ctx context.Context) error {
			// Update meta in store with ImageUrl after screenshot URL is retrieved (or without it if screenshot failed)
			// This ensures meta is always persisted with the latest state
			err := uc.parserUC.MetaStore.Store.Add(ctx, meta)
			if err != nil {
				uc.log.Error("Failed to update meta in store",
					slog.String("error", err.Error()),
					slog.String("url", linkURL),
				)
				return err
			}

			uc.log.Info("Meta updated in store",
				slog.String("url", linkURL),
				slog.String("image_url", meta.ImageUrl),
			)

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

	// Publish MetadataExtracted event using EventBus (canonical name: metadata.metadata.extracted.v1)
	// Published after saga completion to ensure all enrichment (including screenshot) is complete
	event := &v1.MetadataExtracted{
		Id:          meta.Id,
		ImageUrl:    meta.ImageUrl,
		Description: meta.Description,
		Keywords:    meta.Keywords,
		OccurredAt:  timestamppb.Now(),
	}

	if err := uc.eventBus.Publish(ctx, event); err != nil {
		uc.log.Error("Failed to publish metadata extracted event",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.MetadataExtractedTopic),
			slog.String("url", linkURL),
		)
		// Don't fail the operation if event publishing fails
	} else {
		uc.log.Info("Metadata extracted event published successfully",
			slog.String("event_type", domain.MetadataExtractedTopic),
			slog.String("url", linkURL),
		)
	}

	return meta, nil
}
