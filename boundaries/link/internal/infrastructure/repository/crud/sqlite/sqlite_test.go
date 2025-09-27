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
	ctx, cancel := context.WithCancel(context.Background())

	st := &db.Store{}

	t.Setenv("STORE_SQLITE_PATH", "/tmp/links-test.sqlite")

	// Create store
	err := st.Init(ctx)
	require.NoError(t, err)

	// Create repository
	store, err := New(ctx, st)
	require.NoError(t, err)

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})

	t.Run("Close", func(t *testing.T) {
		// Use os.OpenRoot to safely remove the SQLite database file
		root, err := os.OpenRoot("/tmp")
		require.NoError(t, err)
		defer root.Close()

		errDeleteFile := root.Remove("links-test.sqlite")
		require.NoError(t, errDeleteFile)
	})

	t.Cleanup(func() {
		cancel()
	})
}
