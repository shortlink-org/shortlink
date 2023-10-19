package permission

import (
	"context"
	"embed"

	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/pkg/auth"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
)

var (
	//go:embed permissions/*
	permissions embed.FS
)

func Permission(ctx context.Context, log logger.Logger, tracer trace.TracerProvider, monitoring *monitoring.Monitoring) (*auth.Auth, error) {
	permission, err := auth.New(log, tracer, monitoring)
	if err != nil {
		return nil, err
	}

	err = permission.Migrations(ctx, permissions)
	if err != nil {
		return nil, err
	}

	log.Info("Permission migrations completed")

	return permission, nil
}
