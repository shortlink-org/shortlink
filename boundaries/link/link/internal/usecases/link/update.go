package link

import (
	"context"

	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
)

func (uc *UC) Update(ctx context.Context, in *domain.Link) (*domain.Link, error) {
	resp, err := uc.store.Update(ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
