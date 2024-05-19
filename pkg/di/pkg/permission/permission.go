package permission

import (
	"context"
	"embed"

	"github.com/authzed/authzed-go/v1"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/shortlink/pkg/auth"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
)

//go:embed permissions/*
var permissions embed.FS //nolint:unused // ignore

func New(ctx context.Context, log logger.Logger, tracer trace.TracerProvider, monitor *monitoring.Monitoring) (*authzed.Client, error) {
	permission, err := auth.New(log, tracer, monitor)
	if err != nil {
		return nil, err
	}

	return permission, nil
}
