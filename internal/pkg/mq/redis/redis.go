package redis

import (
	"context"

	"github.com/redis/rueidis"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/query"
)

type Config struct{}

type Redis struct {
	*Config
	client rueidis.Client //nolint:unused
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

func (r Redis) Publish(ctx context.Context, target string, routingKey []byte, payload []byte) error {
	// TODO implement me
	panic("implement me")
}

func (r Redis) Subscribe(ctx context.Context, target string, message query.Response) error {
	// TODO implement me
	panic("implement me")
}

func (r Redis) UnSubscribe(target string) error {
	// TODO implement me
	panic("implement me")
}
