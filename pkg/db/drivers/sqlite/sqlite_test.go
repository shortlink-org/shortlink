//go:build unit || (database && sqlite)

package sqlite

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

func TestSQLite(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "sqlite")
	t.Attr("component", "db")
	t.Attr("driver", "sqlite")

		t.Attr("type", "unit")
		t.Attr("package", "sqlite")
		t.Attr("component", "db")
		t.Attr("driver", "sqlite"), cancel := context.WithCancel(t.Context())
	store := Store{}

	err := store.Init(ctx)
	require.NoError(t, err)

	t.Cleanup(func() {
		cancel()
	})
}
