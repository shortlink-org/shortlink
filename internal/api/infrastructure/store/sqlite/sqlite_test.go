package sqlite

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/api/infrastructure/store/mock"
	db "github.com/batazor/shortlink/internal/db/sqlite"
)

//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

func TestSQLite(t *testing.T) {
	ctx := context.Background()

	st := db.Store{}
	err := st.Init(ctx)
	assert.Nil(t, err)

	store := Store{
		client: st.GetConn().(*sql.DB),
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		assert.Nil(t, err)
		assert.Equal(t, len(links), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.Nil(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}
