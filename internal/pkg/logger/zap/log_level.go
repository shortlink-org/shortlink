package zap

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/tracer"
)

// Fatal ===============================================================================================================

func (log *Logger) Fatal(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Fatal(msg, zapFields...)
}

func (log *Logger) FatalWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Fatal(msg, zapFields...)
}

// Warn ================================================================================================================

func (log *Logger) Warn(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Warn(msg, zapFields...)
}

func (log *Logger) WarnWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		//nolint:revive // TODO: fix
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Warn(msg, zapFields...)
}

// Error ===============================================================================================================

func (log *Logger) Error(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Error(msg, zapFields...)
}

func (log *Logger) ErrorWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	tags := []attribute.KeyValue{{
		Key:   "error",
		Value: attribute.BoolValue(true),
	}}

	fields, err := tracer.NewTraceFromContext(ctx, msg, tags, fields...)
	if err != nil {
		//nolint:revive // TODO: fix
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Error(msg, zapFields...)
}

// Info ================================================================================================================

func (log *Logger) Info(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Info(msg, zapFields...)
}

func (log *Logger) InfoWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		//nolint:revive // TODO: fix
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Info(msg, zapFields...)
}

// Debug ===============================================================================================================

func (log *Logger) Debug(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Debug(msg, zapFields...)
}

func (log *Logger) DebugWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		//nolint:revive // TODO: fix
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Debug(msg, zapFields...)
}
