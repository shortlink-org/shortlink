package storeRepository

import (
	"context"

	rpc "github.com/shortlink-org/shortlink/boundaries/metadata/internal/domain/metadata/v1"
)

// Repository defines metadata persistence operations.
type Repository interface {
	Get(ctx context.Context, linkID string) (*rpc.Meta, error)
	Add(ctx context.Context, meta *rpc.Meta) error
}

// MetaStore wraps a concrete Repository implementation selected via configuration.
type MetaStore struct {
	Store Repository
	// notify.Subscriber[link.Link]
	typeStore string
}
