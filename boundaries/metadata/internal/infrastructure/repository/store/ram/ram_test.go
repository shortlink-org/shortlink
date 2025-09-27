//go:build unit || (database && ram)

package ram

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRAM(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "ram")
	t.Attr("component", "metadata")
	t.Attr("driver", "ram")

	// InitStore
	store := Store{}

		ctx := t.Context()

	// Run tests
	t.Run("Create", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "ram")
		t.Attr("component", "metadata")
		t.Attr("driver", "ram")

		// errAdd := store.Add(ctx, &v1.Meta{
		// 	FieldMask:   nil,
		// 	Id:          "1",
		// 	ImageUrl:    "",
		// 	Description: "123",
		// 	Keywords:    "",
		// })
		// require.NoError(t, errAdd)
	})

	t.Run("Get", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "ram")
		t.Attr("component", "metadata")
		t.Attr("driver", "ram")

		meta, err := store.Get(ctx, "1")
		require.NoError(t, err)
		assert.Equal(t, meta.Description, "123")
	})
}
