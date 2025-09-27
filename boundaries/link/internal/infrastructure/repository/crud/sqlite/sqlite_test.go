//go:build unit || (database && sqlite)

package sqlite

import (
	"context"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mock"
	db "github.com/shortlink-org/shortlink/pkg/db/drivers/sqlite"
)

func TestMain(m *testing.M) {

	goleak.VerifyTestMain(m,
		goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"))

	os.Exit(m.Run())
}

func TestSQLite(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "sqlite")
	t.Attr("component", "link")
	t.Attr("driver", "sqlite")

	st := &db.Store{}

	t.Setenv("STORE_SQLITE_PATH", "/tmp/links-test.sqlite")

	// Create store
	err := st.Init(ctx)
	require.NoError(t, err)

	// Create repository
	store, err := New(ctx, st)
	require.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "sqlite")
		t.Attr("component", "link")
		t.Attr("driver", "sqlite")

		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "sqlite")
		t.Attr("component", "link")
		t.Attr("driver", "sqlite")

		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get list", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "sqlite")
		t.Attr("component", "link")
		t.Attr("driver", "sqlite")

		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "sqlite")
		t.Attr("component", "link")
		t.Attr("driver", "sqlite")

		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})

	t.Run("Close", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "sqlite")
		t.Attr("component", "link")
		t.Attr("driver", "sqlite")

		errDeleteFile := os.Remove(viper.GetString("STORE_SQLITE_PATH"))
		require.NoError(t, errDeleteFile)
	})

	t.Cleanup(func() {
		cancel()
	})
}
