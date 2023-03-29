//go:build unit || (database && ram)
// +build unit database,ram

package ram

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
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
	require.NoError(t, err)

	// Run tests
	t.Run("Close", func(t *testing.T) {
		require.NoError(t, store.Close())
	})
}
