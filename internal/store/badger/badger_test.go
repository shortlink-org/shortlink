package badger

import (
	"testing"

	"github.com/batazor/shortlink/internal/store/mock"
	"github.com/stretchr/testify/assert"
)

func TestBadger(t *testing.T) {
	store := BadgerLinkList{}

	err := store.Init()
	assert.Nil(t, err)

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(mock.GetLink.Hash)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(nil)
		assert.Nil(t, err)
		assert.Equal(t, len(links), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.Nil(t, store.Delete(mock.GetLink.Hash))
	})

	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
