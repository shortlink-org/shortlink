package ram

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/store/mock"
	"github.com/batazor/shortlink/internal/store/options"
)

// TODO: problem with goleak
//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

func TestRAM(t *testing.T) {
	store := RAMLinkList{}

	ctx := context.Background()

	err := store.Init(ctx)
	assert.Nil(t, err)

	t.Run("Create [single]", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Create [batch]", func(t *testing.T) {
		// Set config
		err = os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		assert.Nil(t, err, "Cannot set ENV")

		err = store.Init(ctx)
		assert.Nil(t, err)

		link, err := store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)

		link, err = store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)

		link, err = store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)

		link, err = store.Add(ctx, mock.AddLink)
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
