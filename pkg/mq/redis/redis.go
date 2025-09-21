package redis

import (
	"context"

	"github.com/redis/rueidis"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/db"
	"github.com/shortlink-org/shortlink/pkg/db/drivers/redis"
	"github.com/shortlink-org/shortlink/pkg/mq/query"
)

type Redis struct {
	client rueidis.Client //nolint:unused // TODO implement me
}

func New() *Redis {
	return &Redis{}
}

func (r *Redis) Init(ctx context.Context, log logger.Logger) error {
	var ok bool
	mq := &redis.Store{}

	err := mq.Init(ctx)
	if err != nil {
		return err
	}

	r.client, ok = mq.GetConn().(rueidis.Client)
	if !ok {
		return db.ErrGetConnection
	}

	// Graceful shutdown
	go func() {
		<-ctx.Done()

		if errClose := r.close(); errClose != nil {
			log.Error("Redis close error", field.Fields{
				"error": errClose.Error(),
			})
		}
	}()

	return nil
}

// close - close connection
//
//nolint:unparam // ignore unused parameter
func (r *Redis) close() error {
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
