//go:build unit

package v1

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestNewURL(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "v1")
	t.Attr("component", "link")
	

	source := "http://test.com"

	t.Run("create new", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "v1")
		t.Attr("component", "link")

		link, err := NewLinkBuilder().SetURL(source).Build()

		require.NoError(t, err, "Assert nil. Got: %s", err)
		assert.Equal(t, "test.com", link.GetUrl().Host)
	})

	t.Run("create hash", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "v1")
		t.Attr("component", "link")

		success := "e59277e699090366a7751e5c0c5aceaa89379c6ab65f0cf08ef9f62f133c8d07bd2eb6740f5fd65f816b103e2242b7a54a1c1fbb23749609613def7cc6a335d3"
		response := createHash([]byte("hello world"), []byte("salt"))
		assert.Equal(t, success, response)
	})
}
