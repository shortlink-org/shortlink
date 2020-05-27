package grpcweb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/batazor/shortlink/pkg/link"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestAPI(t *testing.T) {
	ctx := context.Background()

	server := API{ctx: ctx}

	t.Run("GetLink", func(t *testing.T) {
		_, err := server.GetLink(ctx, &link.Link{Hash: "example.com"})
		assert.Equal(t, err.Error(), "Not found subscribe to event METHOD_GET")
	})

	t.Run("GetLinks", func(t *testing.T) {
		_, err := server.GetLinks(ctx, &link.Link{Url: "example.com"})
		assert.Equal(t, err.Error(), "Not found subscribe to event METHOD_LIST")
	})

	t.Run("CreateLink", func(t *testing.T) {
		_, err := server.CreateLink(ctx, &link.Link{Url: "example.com"})
		assert.Equal(t, err.Error(), "Not found subscribe to event METHOD_ADD")
	})

	t.Run("DeleteLink", func(t *testing.T) {
		_, err := server.DeleteLink(ctx, &link.Link{Url: "example.com"})
		assert.Equal(t, err.Error(), "Not found subscribe to event METHOD_DELETE")
	})
}
