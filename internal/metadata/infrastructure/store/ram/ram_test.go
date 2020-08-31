package ram

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/metadata/infrastructure/store/mock"
)

func TestRAM(t *testing.T) {
	// InitStore
	store := Store{}

	ctx := context.Background()

	// Run tests
	t.Run("Create", func(t *testing.T) {
		errAdd := store.Add(ctx, mock.AddMetaLink)
		assert.Nil(t, errAdd)
	})

	t.Run("Get", func(t *testing.T) {
		meta, err := store.Get(ctx, mock.GetMetaLink.Id)
		assert.Nil(t, err)
		assert.Equal(t, meta.Description, mock.GetMetaLink.Description)
	})
}
