//go:build unit || (database && redis)

package redis

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	rediscontainer "github.com/testcontainers/testcontainers-go/modules/redis"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/trace/noop"
	"go.uber.org/goleak"

	"github.com/shortlink-org/go-sdk/config"
	db "github.com/shortlink-org/go-sdk/db/drivers/redis"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/mock"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"))

	os.Exit(m.Run())
}

func TestRedis(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	t.Setenv("STORE_REDIS_URI", startRedisContainer(t))

	cfg, err := config.New()
	require.NoError(t, err)

	st := db.New(noop.NewTracerProvider(), metric.NewMeterProvider(), cfg)
	require.NoError(t, st.Init(ctx))

	store, err := New(ctx, st, cfg)
	require.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, mock.GetLink.GetHash(), link.GetHash())
		assert.Equal(t, mock.GetLink.GetDescribe(), link.GetDescribe())
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.GetHash())
		require.NoError(t, err)
		assert.Equal(t, mock.GetLink.GetHash(), link.GetHash())
		assert.Equal(t, mock.GetLink.GetDescribe(), link.GetDescribe())
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, 1, links.Count())
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mock.GetLink.GetHash()))
	})
}

func startRedisContainer(tb testing.TB) string {
	tb.Helper()

	ctx := context.Background()

	container, err := rediscontainer.Run(ctx, "redis:8-alpine")
	require.NoError(tb, err)

	tb.Cleanup(func() {
		terminateCtx, terminateCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer terminateCancel()

		require.NoError(tb, container.Terminate(terminateCtx))
	})

	endpoint, err := container.PortEndpoint(ctx, "6379/tcp", "")
	require.NoError(tb, err)

	return endpoint
}
