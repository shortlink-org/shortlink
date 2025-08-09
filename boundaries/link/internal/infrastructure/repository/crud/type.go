package crud

import (
	"context"

	"github.com/go-redis/cache/v9"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	types "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

// Repository abstract
type Repository interface {
	Get(ctx context.Context, id string) (*v1.Link, error)
	List(ctx context.Context, filter *types.FilterLink) (*v1.Links, error)

	Add(ctx context.Context, in *v1.Link) (*v1.Link, error)
	Update(ctx context.Context, in *v1.Link) (*v1.Link, error)
	Delete(ctx context.Context, id string) error
}

// Store abstract type
type Store struct {
	log       logger.Logger
	store     Repository
	cache     *cache.Cache
	typeStore string
}
