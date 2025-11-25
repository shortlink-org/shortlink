//go:build unit || (database && mysql)

package mysql

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mysqlcontainer "github.com/testcontainers/testcontainers-go/modules/mysql"
	"go.uber.org/goleak"

	db "github.com/shortlink-org/go-sdk/db/drivers/mysql"
	"github.com/shortlink-org/shortlink/boundaries/link/internal/infrastructure/repository/crud/mock"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("github.com/go-sql-driver/mysql.(*mysqlConn).startWatcher.func1"))

	os.Exit(m.Run())
}

func TestMysql(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	st := &db.Store{}

	t.Setenv("STORE_MYSQL_URI", startMySQLContainer(t))
	require.NoError(t, st.Init(ctx))

	// new store
	store, err := New(ctx, st)
	if err != nil {
		t.Fatalf("Could not create store: %s", err)
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		if err != nil {
			t.Fatalf("Could not add link: %s", err)
		}
		assert.Equal(t, mock.AddLink.Hash, link.Hash)
		assert.Equal(t, mock.AddLink.Describe, link.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		if err != nil {
			t.Fatalf("Could not get link: %s", err)
		}
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

func startMySQLContainer(tb testing.TB) string {
	tb.Helper()

	ctx := context.Background()

	container, err := mysqlcontainer.Run(
		ctx,
		"mysql:8",
		mysqlcontainer.WithUsername("shortlink"),
		mysqlcontainer.WithPassword("shortlink"),
		mysqlcontainer.WithDatabase("link"),
	)
	require.NoError(tb, err)

	tb.Cleanup(func() {
		terminateCtx, terminateCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer terminateCancel()

		require.NoError(tb, container.Terminate(terminateCtx))
	})

	return container.MustConnectionString(ctx, "parseTime=true", "loc=UTC")
}
