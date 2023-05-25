package zap

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
	"github.com/shortlink-org/shortlink/internal/pkg/logger/tracer"
)

// Fatal ===============================================================================================================

func (log *ZapLogger) Fatal(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Fatal(msg, zapFields...)
}

func (log *ZapLogger) FatalWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Fatal(msg, zapFields...)
}

// Warn ================================================================================================================

func (log *ZapLogger) Warn(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Warn(msg, zapFields...)
}

func (log *ZapLogger) WarnWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Warn(msg, zapFields...)
}

// Error ===============================================================================================================

func (log *ZapLogger) Error(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Error(msg, zapFields...)
}

func (log *ZapLogger) ErrorWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	tags := []attribute.KeyValue{{
		Key:   "error",
		Value: attribute.BoolValue(true),
	}}

	fields, err := tracer.NewTraceFromContext(ctx, msg, tags, fields...)
	if err != nil {
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Error(msg, zapFields...)
}

// Info ================================================================================================================

func (log *ZapLogger) Info(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Info(msg, zapFields...)
}

func (log *ZapLogger) InfoWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Info(msg, zapFields...)
}

// Debug ===============================================================================================================

func (log *ZapLogger) Debug(msg string, fields ...field.Fields) {
	zapFields := log.converter(fields...)
	log.Logger.Debug(msg, zapFields...)
}

func (log *ZapLogger) DebugWithContext(ctx context.Context, msg string, fields ...field.Fields) {
	fields, err := tracer.NewTraceFromContext(ctx, msg, nil, fields...)
	if err != nil {
		log.Logger.Ctx(ctx).Error(fmt.Sprintf("Error send span to openTelemetry: %s", err.Error()))
	}

	zapFields := log.converter(fields...)
	log.Logger.Ctx(ctx).Debug(msg, zapFields...)
}
