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

	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mock"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	m.Run()
}

func TestRAM(t *testing.T) {
	ctx := context.Background()

	t.Run("Create [single]", func(t *testing.T) {
		store, err := New(ctx)
		require.NoError(t, err)

		link, errAdd := store.Add(ctx, mock.AddLink)
		require.NoError(t, errAdd)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)

		t.Cleanup(func() {
			errClose := store.Close()
			require.NoError(t, errClose)
		})
	})

	t.Run("Create [batch]", func(t *testing.T) {
		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		require.NoError(t, err, "Cannot set ENV")

		storeBatchMode, err := New(ctx)
		require.NoError(t, err)

		for i := 0; i < 4; i++ {
			link, errBatchMode := storeBatchMode.Add(ctx, mock.AddLink)
			require.NoError(t, errBatchMode)
			assert.Equal(t, link.Hash, mock.GetLink.Hash)
			assert.Equal(t, link.Describe, mock.GetLink.Describe)
		}

		t.Cleanup(func() {
			errClose := storeBatchMode.Close()
			require.NoError(t, errClose)
		})
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

		t.Cleanup(func() {
			errClose := store.Close()
			require.NoError(t, errClose)
		})
	})

	t.Run("Get list", func(t *testing.T) {
		store, err := New(ctx)
		require.NoError(t, err)

		_, err = store.Add(ctx, mock.GetLink)
		require.NoError(t, err)

		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)

		t.Cleanup(func() {
			errClose := store.Close()
			require.NoError(t, errClose)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		store, err := New(ctx)
		require.NoError(t, err)

		link, err := store.Add(ctx, mock.GetLink)

		require.NoError(t, store.Delete(ctx, link.Hash))

		t.Cleanup(func() {
			errClose := store.Close()
			require.NoError(t, errClose)
		})
	})
}
