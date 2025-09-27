//go:build unit || (database && ram)

package ram

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mock"
	"github.com/shortlink-org/shortlink/pkg/db/options"
)

func TestMain(m *testing.M) {

	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestRAM(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "ram")
	t.Attr("component", "link")
	t.Attr("driver", "ram")

	defer cancel()

	store, err := New(ctx)
	require.NoError(t, err)

	t.Run("Create [single]", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "ram")
		t.Attr("component", "link")
		t.Attr("driver", "ram")

		createdLink, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err, "Failed to add Link to store")
		assert.Equal(t, mock.AddLink.GetHash(), createdLink.GetHash(), "Hashes should match")
		assert.Equal(t, mock.AddLink.GetDescribe(), createdLink.GetDescribe(), "Descriptions should match")
		assert.False(t, createdLink.GetCreatedAt().GetTime().IsZero(), "CreatedAt should be set")
	})

	t.Run("Create [batch]", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "ram")
		t.Attr("component", "link")
		t.Attr("driver", "ram")

		// Set config
		t.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		storeBatchMode, err := New(ctx)
		require.NoError(t, err)

		for i := 0; i < 4; i++ {
			// Use LinkBuilder to create a new Link instance
			linkURL := fmt.Sprintf("http://example.com/batch/%d", linkUniqId.Add(1))
			linkBuilder := v1.NewLinkBuilder().
				SetURL(linkURL).
				SetDescribe("Batch link description")
			link, err := linkBuilder.Build()
			require.NoError(t, err, "Failed to build Link using LinkBuilder")

			createdLink, err := storeBatchMode.Add(ctx, link)
			require.NoError(t, err, "Failed to add Link to store in batch mode")
			assert.Equal(t, link.GetHash(), createdLink.GetHash(), "Hashes should match")
			assert.Equal(t, link.GetDescribe(), createdLink.GetDescribe(), "Descriptions should match")
			assert.False(t, createdLink.GetCreatedAt().GetTime().IsZero(), "CreatedAt should be set")
		}
	})

	t.Run("Get", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "ram")
		t.Attr("component", "link")
		t.Attr("driver", "ram")

		retrievedLink, err := store.Get(ctx, mock.GetLink.GetHash())
		require.NoError(t, err, "Failed to get Link from store")
		assert.Equal(t, mock.GetLink.GetHash(), retrievedLink.GetHash(), "Hashes should match")
		assert.Equal(t, mock.GetLink.GetDescribe(), retrievedLink.GetDescribe(), "Descriptions should match")
	})

	t.Run("Get list", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "ram")
		t.Attr("component", "link")
		t.Attr("driver", "ram")

		// Set up data needed for the test
		// Add multiple links
		for i := 0; i < 5; i++ {
			link, err := getLink()
			require.NoError(t, err)
			_, err = store.Add(ctx, link)
			require.NoError(t, err)
		}

		links, err := store.List(ctx, nil)
		require.NoError(t, err, "Failed to list Links from store")
		assert.GreaterOrEqual(t, len(links.GetLinks()), 5, "Should have at least 5 Links")
	})

	t.Run("Delete", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "ram")
		t.Attr("component", "link")
		t.Attr("driver", "ram")

		// Set up data needed for the test
		linkBuilder := v1.NewLinkBuilder().
			SetURL("http://example.com/delete").
			SetDescribe("Delete link description")
		link, err := linkBuilder.Build()
		require.NoError(t, err, "Failed to build Link using LinkBuilder")

		_, err = store.Add(ctx, link)
		require.NoError(t, err, "Failed to add Link to store")

		// Proceed with the test
		err = store.Delete(ctx, link.GetHash())
		require.NoError(t, err, "Failed to delete Link from store")
	})

	t.Cleanup(func() {
		cancel()
	})
}
