package leveldb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelDB(t *testing.T) {
	store := Store{}

	ctx := context.Background()

	err := store.Init(ctx)
	assert.Nil(t, err)

	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
