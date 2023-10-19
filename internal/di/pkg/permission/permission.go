package permission

import (
	"context"
	"embed"

	"github.com/authzed/authzed-go/v1"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/internal/pkg/auth"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
)

var (
	//go:embed permissions/*
	permissions embed.FS
)

func New(ctx context.Context, log logger.Logger, tracer trace.TracerProvider, monitoring *monitoring.Monitoring) (*authzed.Client, error) {
	permission, err := auth.New(log, tracer, monitoring)
	if err != nil {
		return nil, err
	}

	return permission, nil
}
