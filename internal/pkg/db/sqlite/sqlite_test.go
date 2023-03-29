//go:build unit || (database && sqlite)
// +build unit database,sqlite

package sqlite

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// TODO: fix
	// goleak.VerifyTestMain(m)
}

func TestSQLite(t *testing.T) {
	store := Store{}

	ctx := context.Background()

	err := store.Init(ctx)
	require.NoError(t, err)

	t.Run("Close", func(t *testing.T) {
		require.NoError(t, store.Close())
	})
}
