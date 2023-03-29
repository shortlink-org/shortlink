//go:build unit || (database && leveldb)
// +build unit database,leveldb

package leveldb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/syndtr/goleveldb/leveldb"

	db "github.com/shortlink-org/shortlink/internal/pkg/db/leveldb"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/mock"
)

func TestLevelDB(t *testing.T) {
	ctx := context.Background()

	st := db.Store{}

	err := st.Init(ctx)
	require.NoError(t, err)

	store := Store{
		client: st.GetConn().(*leveldb.DB),
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
