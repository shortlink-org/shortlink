package query

import (
	"context"

	"github.com/go-redis/cache/v8"

	"github.com/batazor/shortlink/internal/pkg/logger"
	v12 "github.com/batazor/shortlink/internal/services/link/domain/link_cqrs/v1"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/query"
)

type Repository interface {
	Get(ctx context.Context, id string) (*v12.LinkView, error)
	List(ctx context.Context, filter *query.Filter) (*v12.LinksView, error)
}

// Store abstract type
type Store struct { // nolint unused
	cache *cache.Cache
	log   logger.Logger

	typeStore string
	store     Repository
}
