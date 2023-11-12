package cache

import (
	"context"

	"github.com/go-redis/cache/v9"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidiscompat"
	"github.com/spf13/viper"

	db2 "github.com/shortlink-org/shortlink/internal/pkg/db"
	db "github.com/shortlink-org/shortlink/internal/pkg/db/redis"
)

// New returns a new cache.Client.
func New(ctx context.Context) (*cache.Cache, error) {
	viper.SetDefault("LOCAL_CACHE_TTL", "5m")
	viper.SetDefault("LOCAL_CACHE_COUNT", 1000)

	store := &db.Store{}
	err := store.Init(ctx)
	if err != nil {
		return nil, err
	}

	conn, ok := store.GetConn().(rueidis.Client)
	if !ok {
		return nil, db2.ErrGetConnection
	}

	adapter := &client{
		rueidiscompat.NewAdapter(conn),
	}

	s := cache.New(&cache.Options{
		Redis:      adapter,
		LocalCache: cache.NewTinyLFU(viper.GetInt("LOCAL_CACHE_COUNT"), viper.GetDuration("LOCAL_CACHE_TTL")),
	})

	return s, nil
}
