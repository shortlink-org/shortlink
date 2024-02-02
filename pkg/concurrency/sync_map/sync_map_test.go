package sync_map

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SyncMap(t *testing.T) {
	sm := New()

	for i := 0; i < 1000; i++ {
		sm.Set(i, "value")
	}

	require.Equal(t, "value", sm.Get(5))
}
