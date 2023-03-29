//go:build unit || (database && badger)
// +build unit database,badger

package badger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

func TestBadger(t *testing.T) {
	store := Store{}
	ctx := context.Background()

	err := store.Init(ctx)
	require.NoError(t, err)

	t.Run("Close", func(t *testing.T) {
		require.NoError(t, store.Close())
	})
}
