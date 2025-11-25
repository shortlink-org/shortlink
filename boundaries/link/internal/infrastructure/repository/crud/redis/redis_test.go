//go:build unit || (database && redis)

package redis

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/rueidis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	rediscontainer "github.com/testcontainers/testcontainers-go/modules/redis"
	"go.uber.org/goleak"

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

	st := db.Store{}

	t.Setenv("STORE_REDIS_URI", startRedisContainer(t))
	require.NoError(t, st.Init(ctx))

	store := Store{
		client: st.GetConn().(rueidis.Client),
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}

func startRedisContainer(tb testing.TB) string {
	tb.Helper()

	ctx := context.Background()

	container, err := rediscontainer.Run(ctx, "redis:7-alpine")
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
