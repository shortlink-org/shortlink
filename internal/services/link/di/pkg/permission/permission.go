package permission

import (
	"context"

	"github.com/shortlink-org/shortlink/internal/pkg/auth"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
)

func Permission(ctx context.Context, log logger.Logger) (*authzed.Client, error) {
	permission, err := auth.New()
	if err != nil {
		return nil, err
	}

	return permission, nil
}
