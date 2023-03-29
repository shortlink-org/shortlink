//go:build unit || (database && ram)
// +build unit database,ram

package ram

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRAM(t *testing.T) {
	// InitStore
	store := Store{}

	ctx := context.Background()

	// Run tests
	t.Run("Create", func(t *testing.T) {
		errAdd := store.Add(ctx, mock.AddMetaLink)
		require.NoError(t, errAdd)
	})

	t.Run("Get", func(t *testing.T) {
		meta, err := store.Get(ctx, mock.GetMetaLink.Id)
		require.NoError(t, err)
		assert.Equal(t, meta.Description, mock.GetMetaLink.Description)
	})
}
