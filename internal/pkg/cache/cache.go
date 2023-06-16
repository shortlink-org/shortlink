package cache

import (
	"context"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidiscompat"

	db "github.com/shortlink-org/shortlink/internal/pkg/db/redis"
)

// New returns a new cache.Client.
func New(ctx context.Context) (*cache.Cache, error) {
	store := &db.Store{}
	err := store.Init(ctx)
	if err != nil {
		return nil, err
	}

	adapter := &client{
		rueidiscompat.NewAdapter(store.GetConn().(rueidis.Client)),
	}

	s := cache.New(&cache.Options{
		Redis:      adapter,
		LocalCache: cache.NewTinyLFU(1000, 5*time.Minute), // nolint:gomnd
	})

	return s, nil
}
