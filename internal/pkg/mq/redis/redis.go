package redis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

type Config struct{}

type Redis struct {
	*Config
	client redis.UniversalClient //nolint:unused
}

func New() *Redis {
	return &Redis{}
}

func (r Redis) Init(ctx context.Context) error {
	// TODO implement me
	panic("implement me")
}

func (r Redis) Close() error {
	// TODO implement me
	panic("implement me")
}

func (r Redis) Publish(ctx context.Context, target string, message query.Message) error {
	// TODO implement me
	panic("implement me")
}

func (r Redis) Subscribe(target string, message query.Response) error {
	// TODO implement me
	panic("implement me")
}

func (r Redis) UnSubscribe(target string) error {
	// TODO implement me
	panic("implement me")
}
