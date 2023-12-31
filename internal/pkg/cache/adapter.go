package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/rueidis/rueidiscompat"
)

type client struct {
	adapter rueidiscompat.Cmdable
}

func (c client) Set(ctx context.Context, key string, value any, ttl time.Duration) *redis.StatusCmd {
	start := time.Now()

	resp := c.adapter.Set(ctx, key, value, ttl)

	// Record the duration and increment the appropriate metric based on the response
	cacheOperationDuration.WithLabelValues("set", key).Observe(time.Since(start).Seconds())
	if resp.Err() != nil {
		cacheOperationErrors.WithLabelValues("set", key).Inc()
	} else {
		cacheOperations.WithLabelValues("set", "success", key).Inc()
	}

	return redis.NewStatusCmd(ctx, resp)
}

func (c client) SetXX(ctx context.Context, key string, value any, ttl time.Duration) *redis.BoolCmd {
	resp := c.adapter.SetXX(ctx, key, value, ttl)
	return redis.NewBoolCmd(ctx, resp)
}

func (c client) SetNX(ctx context.Context, key string, value any, ttl time.Duration) *redis.BoolCmd {
	resp := c.adapter.SetNX(ctx, key, value, ttl)
	return redis.NewBoolCmd(ctx, resp)
}

func (c client) Get(ctx context.Context, key string) *redis.StringCmd {
	start := time.Now()

	resp := c.adapter.Get(ctx, key)

	// Record the duration and increment the appropriate metric based on the response
	cacheOperationDuration.WithLabelValues("get", key).Observe(time.Since(start).Seconds())
	if resp.Err() != nil {
		cacheOperationErrors.WithLabelValues("get", key).Inc()
	} else {
		cacheOperations.WithLabelValues("get", "success", key).Inc()
	}

	return redis.NewStringCmd(ctx, resp)
}

func (c client) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	resp := c.adapter.Del(ctx, keys...)
	return redis.NewIntCmd(ctx, resp)
}
