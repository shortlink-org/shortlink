package crud

import (
	"context"

	"github.com/go-redis/cache/v9"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v1 "github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

type Repository interface {
	Get(ctx context.Context, id string) (*v1.Link, error)
	List(ctx context.Context, filter *query.Filter) (*v1.Links, error)

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
