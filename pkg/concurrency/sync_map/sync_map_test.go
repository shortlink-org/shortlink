package sync_map

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SyncMap(t *testing.T) {
	sm := New()

	for i := range 1000 {
		sm.Set(i, "value")
	}

	require.Equal(t, "value", sm.Get(5))
}
