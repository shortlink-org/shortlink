//go:build unit || (database && ram)

package ram

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestRAM(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	// InitStore
	store := Store{}

	err := store.Init(ctx)
	require.NoError(t, err)

	// Run tests
	t.Cleanup(func() {
		cancel()
	})
}
