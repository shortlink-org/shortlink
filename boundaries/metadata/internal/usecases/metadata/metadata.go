package metadata

import (
	"context"
	"errors"
	"log/slog"

	"github.com/shortlink-org/go-sdk/cqrs/bus"
	"github.com/shortlink-org/go-sdk/logger"
	"github.com/shortlink-org/go-sdk/saga"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain"
	domainerrors "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/errors"
	v1 "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/parsers"
	"github.com/shortlink-org/shortlink/boundaries/metadata/internal/usecases/screenshot"
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

const (
	OpParserSet     = "metadata.parser.set"
	OpScreenshotSet = "metadata.screenshot.set"
	OpSagaPlay      = "metadata.play"
	OpStoreUpdate   = "metadata.store.update"
)

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
	if len(errs) == 0 {
		return nil
	}

	for index, err := range errs {
		log.ErrorWithContext(ctx, "Saga error",
			slog.Int("stack_index", index),
			slog.Any("error", err),
		)
	}

	return domainerrors.Normalize(OpSagaPlay, errors.Join(errs...))
}

// Add adds a metadata
func (uc *UC) Add(ctx context.Context, linkURL string) (*v1.Meta, error) { //nolint:maintidx,funlen // saga orchestration is inherently complex
	const (
		SAGA_NAME                    = "METADATA_ADD"
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
			ctx, span := otel.Tracer("metadata.uc.parser").Start(ctx, "saga: SAGA_STEP_ADD_META",
				trace.WithSpanKind(trace.SpanKindInternal),
			)
			defer span.End()

			span.SetAttributes(
				attribute.String("step", SAGA_STEP_ADD_META),
				attribute.String("status", "run"),
				attribute.String("link.url", linkURL),
			)

			m, stepErr := uc.parserUC.Set(ctx, linkURL)
			if stepErr != nil {
				span.RecordError(stepErr)
				span.SetStatus(otelcodes.Error, stepErr.Error())
				return domainerrors.Normalize(OpParserSet, stepErr)
			}

			meta = m
			span.SetAttributes(
				attribute.String("meta.id", m.GetId()),
			)
			span.SetStatus(otelcodes.Ok, "Metadata parsed successfully")

			return nil
		}).Build()

	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaSetMetadata.AddStep(SAGA_STEP_ADD_SCREENSHOT).
		Then(func(ctx context.Context) error {
			ctx, span := otel.Tracer("metadata.uc.screenshot").Start(ctx, "saga: SAGA_STEP_ADD_SCREENSHOT",
				trace.WithSpanKind(trace.SpanKindInternal),
			)
			defer span.End()

			span.SetAttributes(
				attribute.String("step", SAGA_STEP_ADD_SCREENSHOT),
				attribute.String("status", "run"),
				attribute.String("link.url", linkURL),
			)

			stepErr := uc.screenshotUC.Set(ctx, linkURL)
			if stepErr != nil {
				span.RecordError(stepErr)
				span.SetStatus(otelcodes.Error, stepErr.Error())
				return domainerrors.Normalize(OpScreenshotSet, stepErr)
			}

			span.SetStatus(otelcodes.Ok, "Screenshot processing started successfully")
			return nil
		}).Build()

	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaSetMetadata.AddStep(SAGA_STEP_GET_SCREENSHOT_URL).
		Needs(SAGA_STEP_ADD_SCREENSHOT, SAGA_STEP_ADD_META).
		Then(func(ctx context.Context) error {
			ctx, span := otel.Tracer("metadata.uc.screenshot").Start(ctx, "saga: SAGA_STEP_GET_SCREENSHOT_URL",
				trace.WithSpanKind(trace.SpanKindInternal),
			)
			defer span.End()

			span.SetAttributes(
				attribute.String("step", SAGA_STEP_GET_SCREENSHOT_URL),
				attribute.String("status", "run"),
				attribute.String("link.url", linkURL),
			)

			// Try to get screenshot URL, but don't fail if screenshot is not available yet
			url, stepErr := uc.screenshotUC.Get(ctx, linkURL)
			if stepErr != nil {
				// Log warning but continue without screenshot URL
				span.AddEvent("Screenshot URL not available yet, continuing without it")
				span.SetAttributes(attribute.Bool("screenshot.available", false))
				uc.log.WarnWithContext(ctx, "Failed to get screenshot URL, continuing without it",
					slog.String("error", stepErr.Error()),
					slog.String("url", linkURL),
				)

				span.SetStatus(otelcodes.Ok, "Continuing without screenshot URL")
				return nil // Continue saga execution even if screenshot URL is not available
			}

			meta.ImageUrl = url.String()
			span.SetAttributes(
				attribute.String("screenshot.url", url.String()),
				attribute.Bool("screenshot.available", true),
			)
			span.SetStatus(otelcodes.Ok, "Screenshot URL retrieved successfully")

			return nil
		}).Build()

	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	_, errs = sagaSetMetadata.AddStep(SAGA_STEP_UPDATE_META).
		Needs(SAGA_STEP_GET_SCREENSHOT_URL).
		Then(func(ctx context.Context) error {
			ctx, span := otel.Tracer("metadata.uc.store").Start(ctx, "saga: SAGA_STEP_UPDATE_META",
				trace.WithSpanKind(trace.SpanKindInternal),
			)
			defer span.End()

			span.SetAttributes(
				attribute.String("step", SAGA_STEP_UPDATE_META),
				attribute.String("status", "run"),
				attribute.String("link.url", linkURL),
				attribute.String("meta.id", meta.GetId()),
			)

			// Update meta in store with ImageUrl after screenshot URL is retrieved (or without it if screenshot failed)
			// This ensures meta is always persisted with the latest state
			storeErr := uc.parserUC.MetaStore.Store.Add(ctx, meta)
			if storeErr != nil {
				span.RecordError(storeErr)
				span.SetStatus(otelcodes.Error, storeErr.Error())
				uc.log.ErrorWithContext(ctx, "Failed to update meta in store",
					slog.String("error", storeErr.Error()),
					slog.String("url", linkURL),
				)

				return domainerrors.Normalize(OpStoreUpdate, storeErr)
			}

			span.SetAttributes(
				attribute.String("meta.image_url", meta.GetImageUrl()),
			)
			span.SetStatus(otelcodes.Ok, "Meta updated in store successfully")
			uc.log.InfoWithContext(ctx, "Meta updated in store",
				slog.String("url", linkURL),
				slog.String("image_url", meta.GetImageUrl()),
			)

			return nil
		}).Build()

	if err := errorHelper(ctx, uc.log, errs); err != nil {
		return nil, err
	}

	// run saga
	if err := sagaSetMetadata.Play(nil); err != nil {
		return nil, domainerrors.Normalize(OpSagaPlay, err)
	}

	// Publish MetadataExtracted event using EventBus (canonical name: metadata.metadata.extracted.v1)
	// Published after saga completion to ensure all enrichment (including screenshot) is complete
	ctx, span := otel.Tracer("metadata.uc.event").Start(ctx, "metadata.uc.publish_metadata_extracted",
		trace.WithSpanKind(trace.SpanKindProducer),
	)
	defer span.End()

	span.SetAttributes(
		attribute.String("messaging.system", "kafka"),
		attribute.String("messaging.destination", domain.MetadataExtractedTopic),
		attribute.String("messaging.destination_kind", "topic"),
		attribute.String("messaging.operation", "publish"),
		attribute.String("event.type", domain.MetadataExtractedTopic),
		attribute.String("link.url", linkURL),
		attribute.String("meta.id", meta.GetId()),
	)

	event := &v1.MetadataExtracted{
		Id:          meta.GetId(),
		ImageUrl:    meta.GetImageUrl(),
		Description: meta.GetDescription(),
		Keywords:    meta.GetKeywords(),
		OccurredAt:  timestamppb.Now(),
	}

	if err := uc.eventBus.Publish(ctx, event); err != nil {
		span.RecordError(err)
		span.SetStatus(otelcodes.Error, err.Error())
		uc.log.ErrorWithContext(ctx, "Failed to publish metadata extracted event",
			slog.String("error", err.Error()),
			slog.String("event_type", domain.MetadataExtractedTopic),
			slog.String("url", linkURL),
		)
		// Don't fail the operation if event publishing fails
	} else {
		span.SetAttributes(
			attribute.String("event_type", domain.MetadataExtractedTopic),
		)
		span.SetStatus(otelcodes.Ok, "Metadata extracted event published successfully")
		uc.log.InfoWithContext(ctx, "Metadata extracted event published successfully",
			slog.String("event_type", domain.MetadataExtractedTopic),
			slog.String("url", linkURL),
		)
	}

	return meta, nil
}

