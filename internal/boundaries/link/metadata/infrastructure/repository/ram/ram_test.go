//go:build unit || (database && ram)

package ram

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/domain/metadata/v1"
)

func TestRAM(t *testing.T) {
	// InitStore
	store := Store{}

	ctx := context.Background()

	// Run tests
	t.Run("Create", func(t *testing.T) {
		errAdd := store.Add(ctx, &v1.Meta{
			FieldMask:   nil,
			Id:          "1",
			ImageUrl:    "",
			Description: "123",
			Keywords:    "",
		})
		require.NoError(t, errAdd)
	})

	t.Run("Get", func(t *testing.T) {
		meta, err := store.Get(ctx, "1")
		require.NoError(t, err)
		assert.Equal(t, meta.Description, "123")
	})
}
