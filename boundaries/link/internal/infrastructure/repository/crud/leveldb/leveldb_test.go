//go:build unit || (database && leveldb)

package leveldb

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mock"
	db "github.com/shortlink-org/shortlink/pkg/db/drivers/leveldb"
)

func TestMain(m *testing.M) {

	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"),
		goleak.IgnoreTopFunction("github.com/syndtr/goleveldb/leveldb.(*DB).mpoolDrain"))

	os.Exit(m.Run())
}

func TestLevelDB(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "leveldb")
	t.Attr("component", "link")
	t.Attr("driver", "leveldb")

	st := db.Store{}

	err := st.Init(ctx)
	require.NoError(t, err)

	store := Store{
		client: st.GetConn().(*leveldb.DB),
	}

	t.Run("Create", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "leveldb")
		t.Attr("component", "link")
		t.Attr("driver", "leveldb")

		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "leveldb")
		t.Attr("component", "link")
		t.Attr("driver", "leveldb")

		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "leveldb")
		t.Attr("component", "link")
		t.Attr("driver", "leveldb")

		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "leveldb")
		t.Attr("component", "link")
		t.Attr("driver", "leveldb")

		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})

	t.Cleanup(func() {
		cancel()
	})
}
