package query

import (
	"context"

	"github.com/go-redis/cache/v9"

	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	v12 "github.com/shortlink-org/shortlink/internal/services/link/domain/link_cqrs/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/query"
)

type Repository interface {
	Get(ctx context.Context, id string) (*v12.LinkView, error)
	List(ctx context.Context, filter *query.Filter) (*v12.LinksView, error)
}

// Store abstract type
type Store struct {
	log       logger.Logger
	store     Repository
	cache     *cache.Cache
	typeStore string
}
