//go:build unit || (database && leveldb)

package leveldb

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {

	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/syndtr/goleveldb/leveldb.(*DB).mpoolDrain"))

	os.Exit(m.Run())
}

func TestLevelDB(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "leveldb")
	t.Attr("component", "db")
	t.Attr("driver", "leveldb")

	store := Store{}

	err := store.Init(ctx)
	require.NoError(t, err)

	t.Cleanup(func() {
		cancel()
	})
}
