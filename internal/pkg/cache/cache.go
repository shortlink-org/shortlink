package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/go-redis/cache/v8"

	db "github.com/batazor/shortlink/internal/pkg/db/redis"
)

func New(ctx context.Context) (*cache.Cache, error) {
	store := &db.Store{}
	err := store.Init(ctx)
	if err != nil {
		return nil, err
	}

	s := cache.New(&cache.Options{
		Redis:      store.GetConn().(*redis.Client),
		LocalCache: cache.NewTinyLFU(1000, 5*time.Minute),
	})

	return s, nil
}
