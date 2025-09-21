package cqs

import (
	"context"

	"github.com/go-redis/cache/v9"

	"github.com/shortlink-org/shortlink/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
)

type Repository interface {
	LinkAdd(ctx context.Context, data *v1.Link) (*v1.Link, error)
	LinkUpdate(ctx context.Context, data *v1.Link) (*v1.Link, error)
	LinkDelete(ctx context.Context, id string) error

	// MetadataUpdate(ctx context.Context, data *v12.Meta) (*v12.Meta, error)
}

// Store abstract type
type Store struct {
	log       logger.Logger
	store     Repository
	cache     *cache.Cache
	typeStore string
}
