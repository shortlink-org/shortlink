package sqlite

import (
	"context"
	"testing"

	"github.com/batazor/shortlink/internal/store/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestSQLite(t *testing.T) {
	store := SQLiteLinkList{}

	ctx := context.Background()

	err := store.Init(ctx)
	assert.Nil(t, err)

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

	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
