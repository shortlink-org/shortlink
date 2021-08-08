// +build unit database,badger

package badger

import (
	"context"
	"testing"

	"github.com/dgraph-io/badger/v3"
	"github.com/stretchr/testify/assert"

	db "github.com/batazor/shortlink/internal/pkg/db/badger"
	"github.com/batazor/shortlink/internal/services/link/infrastructure/store/crud/mock"
)

//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

func TestBadger(t *testing.T) {
	ctx := context.Background()

	st := db.Store{}
	st.Init(ctx)

	store := Store{
		client: st.GetConn().(*badger.DB),
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		assert.Nil(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.Nil(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}
