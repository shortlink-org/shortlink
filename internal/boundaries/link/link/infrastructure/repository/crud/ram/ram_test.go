//go:build unit || (database && ram)

package ram

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/repository/crud/mock"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestRAM(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	t.Run("Create [single]", func(t *testing.T) {
		store, err := New(ctx)
		require.NoError(t, err)

		link, errAdd := store.Add(ctx, mock.AddLink)
		require.NoError(t, errAdd)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Create [batch]", func(t *testing.T) {
		// Set config
		t.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		storeBatchMode, err := New(ctx)
		require.NoError(t, err)

		for i := 0; i < 4; i++ {
			link, errBatchMode := storeBatchMode.Add(ctx, mock.AddLink)
			require.NoError(t, errBatchMode)
			assert.Equal(t, link.Hash, mock.GetLink.Hash)
			assert.Equal(t, link.Describe, mock.GetLink.Describe)
		}
	})

	t.Run("Get", func(t *testing.T) {
		store, err := New(ctx)
		require.NoError(t, err)

		link, err := store.Add(ctx, mock.GetLink)
		require.NoError(t, err)

		link, err = store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		store, err := New(ctx)
		require.NoError(t, err)

		_, err = store.Add(ctx, mock.GetLink)
		require.NoError(t, err)

		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		store, err := New(ctx)
		require.NoError(t, err)

		link, err := store.Add(ctx, mock.GetLink)

		require.NoError(t, store.Delete(ctx, link.Hash))
	})

	t.Cleanup(func() {
		cancel()
	})
}
