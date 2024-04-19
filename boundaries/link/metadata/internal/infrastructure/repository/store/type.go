package storeRepository

import (
	"context"

	rpc "github.com/shortlink-org/shortlink/boundaries/link/metadata/internal/domain/metadata/v1"
)

type Repository interface {
	Get(context.Context, string) (*rpc.Meta, error)
	Add(context.Context, *rpc.Meta) error
}

// Store abstract type
type MetaStore struct {
	Store Repository
	// notify.Subscriber[link.Link]
	typeStore string
}
