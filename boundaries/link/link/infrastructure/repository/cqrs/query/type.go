package query

import (
	"context"

	"github.com/go-redis/cache/v9"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1"
	v12 "github.com/shortlink-org/shortlink/boundaries/link/link/domain/link_cqrs/v1"
	"github.com/shortlink-org/shortlink/pkg/logger"
)

type Repository interface {
	Get(ctx context.Context, id string) (*v12.LinkView, error)
	List(ctx context.Context, filter *v1.FilterLink) (*v12.LinksView, error)
}

// Store abstract type
type Store struct {
	log       logger.Logger
	store     Repository
	cache     *cache.Cache
	typeStore string
}
