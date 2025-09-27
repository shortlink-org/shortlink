//go:build unit || (database && badger)

package badger

import (
	"context"
	"os"
	"testing"

	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mock"
	db "github.com/shortlink-org/shortlink/pkg/db/drivers/badger"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"))

	os.Exit(m.Run())
}

func TestBadger(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "badger")
	t.Attr("component", "link")
	t.Attr("driver", "badger")

		t.Attr("type", "unit")
		t.Attr("package", "badger")
		t.Attr("component", "link")
		t.Attr("driver", "badger"), cancel := context.WithCancel(t.Context())

	st := db.Store{}
	err := st.Init(ctx)
	require.NoError(t, err)

	store := Store{
		client: st.GetConn().(*badger.DB),
	}

	t.Run("Create", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "badger")
		t.Attr("component", "link")
		t.Attr("driver", "badger")

		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "badger")
		t.Attr("component", "link")
		t.Attr("driver", "badger")

		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "badger")
		t.Attr("component", "link")
		t.Attr("driver", "badger")

		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "badger")
		t.Attr("component", "link")
		t.Attr("driver", "badger")

		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})

	t.Cleanup(func() {
		cancel()
	})
}
