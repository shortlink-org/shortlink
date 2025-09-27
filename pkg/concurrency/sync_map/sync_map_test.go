package sync_map_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/pkg/concurrency/sync_map"
)

func Test_SyncMap(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "sync_map")
	t.Attr("component", "concurrency")

		t.Attr("type", "unit")
		t.Attr("package", "sync_map")
		t.Attr("component", "concurrency")
	
	sm := sync_map.New()

	for i := range 1000 {
		sm.Set(i, "value")
	}

	require.Equal(t, "value", sm.Get(5))
}
