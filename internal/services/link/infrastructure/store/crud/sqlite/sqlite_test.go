//go:build unit || (database && sqlite)
// +build unit database,sqlite

package sqlite

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	db "github.com/shortlink-org/shortlink/internal/pkg/db/sqlite"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/mock"
)

//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

func TestSQLite(t *testing.T) {
	ctx := context.Background()

	err := os.Setenv("STORE_SQLITE_PATH", "/tmp/links-test.sqlite")
	require.NoError(t, err, "Cannot set ENV")

	st := db.Store{}
	err = st.Init(ctx)
	require.NoError(t, err)

	store := Store{
		client: st.GetConn().(*sql.DB),
	}

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
		errDeleteFile := os.Remove(viper.GetString("STORE_SQLITE_PATH"))
		require.NoError(t, errDeleteFile)
	})
}
