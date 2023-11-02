package redis

import (
	"context"

	"github.com/redis/rueidis"

	"github.com/shortlink-org/shortlink/internal/pkg/db"
	"github.com/shortlink-org/shortlink/internal/pkg/db/redis"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

type Redis struct {
	client rueidis.Client //nolint:unused // TODO implement me
}

func New() *Redis {
	return &Redis{}
}

func (r *Redis) Init(ctx context.Context) error {
	var ok bool
	store := &redis.Store{}

	err := store.Init(ctx)
	if err != nil {
		return err
	}

	r.client, ok = store.GetConn().(rueidis.Client)
	if !ok {
		return db.ErrGetConnection
	}

	return nil
}

func (r *Redis) Close() error {
	r.client.Close()
	return nil
}

func (r *Redis) Publish(_ context.Context, _ string, _, _ []byte) error {
	// TODO implement me
	panic("implement me")
}

func (r *Redis) Subscribe(_ context.Context, _ string, _ query.Response) error {
	// TODO implement me
	panic("implement me")
}

func (r *Redis) UnSubscribe(_ string) error {
	// TODO implement me
	panic("implement me")
}
