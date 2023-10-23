package redis

import (
	"context"

	"github.com/redis/rueidis"

	"github.com/shortlink-org/shortlink/internal/pkg/db/redis"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

type Redis struct {
	client rueidis.Client //nolint:unused
}

func New() *Redis {
	return &Redis{}
}

func (r *Redis) Init(ctx context.Context) error {
	store := &redis.Store{}

	err := store.Init(ctx)
	if err != nil {
		return err
	}

	r.client = store.GetConn().(rueidis.Client)

	return nil
}

func (r *Redis) Close() error {
	r.client.Close()
	return nil
}

func (r *Redis) Publish(ctx context.Context, target string, routingKey, payload []byte) error {
	// TODO implement me
	panic("implement me")
}

func (r *Redis) Subscribe(ctx context.Context, target string, message query.Response) error {
	// TODO implement me
	panic("implement me")
}

func (r *Redis) UnSubscribe(target string) error {
	// TODO implement me
	panic("implement me")
}
