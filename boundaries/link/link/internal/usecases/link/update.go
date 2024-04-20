package link

import (
	"context"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
)

func (uc *UC) Update(_ context.Context, _ *domain.Link) (*domain.Link, error) {
	return nil, nil
}
