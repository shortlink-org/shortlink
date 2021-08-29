package cqs

import (
	"context"

	"github.com/go-redis/cache/v8"

	"github.com/batazor/shortlink/internal/pkg/logger"
	v1 "github.com/batazor/shortlink/internal/services/link/domain/link/v1"
	v12 "github.com/batazor/shortlink/internal/services/metadata/domain/metadata/v1"
)

type Repository interface {
	LinkAdd(ctx context.Context, data *v1.Link) (*v1.Link, error)
	LinkUpdate(ctx context.Context, data *v1.Link) (*v1.Link, error)
	LinkDelete(ctx context.Context, id string) error

	MetadataUpdate(ctx context.Context, data *v12.Meta) (*v12.Meta, error)
}

// Store abstract type
type Store struct { // nolint unused
	cache *cache.Cache
	log   logger.Logger

	typeStore string
	store     Repository
}
