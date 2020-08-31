package ram

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestRAM(t *testing.T) {
	// InitStore
	store := Store{}

	ctx := context.Background()

	err := store.Init(ctx)
	assert.Nil(t, err)

	// Run tests
	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
