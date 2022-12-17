//go:build unit || (database && sqlite)
// +build unit database,sqlite

package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// TODO: fix
	// goleak.VerifyTestMain(m)
}

func TestSQLite(t *testing.T) {
	store := Store{}

	ctx := context.Background()

	err := store.Init(ctx)
	assert.Nil(t, err)

	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
