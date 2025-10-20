package cache

import (
	"context"

	"github.com/go-redis/cache/v9"
	"github.com/redis/rueidis"
	"github.com/redis/rueidis/rueidiscompat"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"

	"github.com/shortlink-org/go-sdk/observability/metrics"
	db2 "github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/db/drivers/redis"
)

// New returns a new cache.Client.
func New(ctx context.Context, tracer trace.TracerProvider, monitor *metrics.Monitoring) (*cache.Cache, error) {
	viper.SetDefault("LOCAL_CACHE_TTL", "5m")
	viper.SetDefault("LOCAL_CACHE_COUNT", 1000)
	viper.SetDefault("LOCAL_CACHE_METRICS_ENABLED", true)

	store := redis.New(tracer, monitor.Metrics)

	err := store.Init(ctx)
	if err != nil {
		return nil, &InitCacheError{err}
	}

	conn, ok := store.GetConn().(rueidis.Client)
	if !ok {
		return nil, db2.ErrGetConnection
	}

	adapter := &client{
		rueidiscompat.NewAdapter(conn),
	}

	s := cache.New(&cache.Options{
		Redis:        adapter,
		LocalCache:   cache.NewTinyLFU(viper.GetInt("LOCAL_CACHE_COUNT"), viper.GetDuration("LOCAL_CACHE_TTL")),
		StatsEnabled: viper.GetBool("LOCAL_CACHE_METRICS_ENABLED"),
	})

	return s, nil
}
