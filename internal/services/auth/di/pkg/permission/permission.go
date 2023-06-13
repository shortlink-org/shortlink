package permission

import (
	"context"
	"embed"

	"github.com/shortlink-org/shortlink/internal/pkg/auth"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

var (
	//go:embed permissions/*
	permissions embed.FS
)

func Permission(ctx context.Context, log logger.Logger) (*auth.Auth, error) {
	permission, err := auth.New()
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
