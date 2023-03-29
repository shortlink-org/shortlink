//go:build unit || (database && leveldb)
// +build unit database,leveldb

package leveldb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLevelDB(t *testing.T) {
	store := Store{}

	ctx := context.Background()

	err := store.Init(ctx)
	require.NoError(t, err)

	t.Run("Close", func(t *testing.T) {
		require.NoError(t, store.Close())
	})
}
