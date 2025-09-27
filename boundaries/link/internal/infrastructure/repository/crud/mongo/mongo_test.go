//go:build unit || (database && mongo)

package mongo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
	"go.uber.org/goleak"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mock"
	filter2 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mongo/filter"
	domain "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/types/v1"
	db "github.com/shortlink-org/shortlink/pkg/db/drivers/mongo"
	"github.com/shortlink-org/shortlink/pkg/db/options"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
	os.Exit(m.Run())
}

var linkUniqId atomic.Int64

func TestMongo(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "mongo")
	t.Attr("component", "link")
	t.Attr("driver", "mongo")

		t.Attr("type", "unit")
		t.Attr("package", "mongo")
		t.Attr("component", "link")
		t.Attr("driver", "mongo")
	
	ctx, cancel := context.WithCancel(context.Background())

	st := &db.Store{}

	// Set up Docker MongoDB
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to Docker")

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "7.0",
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	require.NoError(t, err, "Could not start MongoDB Docker container")

	// Exponential backoff-retry to ensure MongoDB is ready
	if err := pool.Retry(func() error {
		mongoURI := fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp"))
		t.Setenv("STORE_MONGODB_URI", mongoURI)
		return st.Init(ctx)
	}); err != nil {
		require.NoError(t, err, "Could not connect to MongoDB Docker container")
	}

	// Cleanup Docker resources after tests
	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge MongoDB Docker container: %s", err)
		}
	})

	// Initialize the store
	store, err := New(ctx, st)
	require.NoError(t, err, "Could not create MongoDB store")

	t.Run("Create [single]", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mongo")
		t.Attr("component", "link")
		t.Attr("driver", "mongo")

		createdLink, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err, "Failed to add Link to store")
		assert.Equal(t, mock.AddLink.GetHash(), createdLink.GetHash(), "Hashes should match")
		assert.Equal(t, mock.AddLink.GetDescribe(), createdLink.GetDescribe(), "Descriptions should match")
		assert.False(t, createdLink.GetCreatedAt().GetTime().IsZero(), "CreatedAt should be set")
	})

	t.Run("Create [batch]", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mongo")
		t.Attr("component", "link")
		t.Attr("driver", "mongo")

		// Set config to batch mode
		t.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		// Initialize store in batch mode
		storeBatchMode, err := New(ctx, st)
		require.NoError(t, err, "Could not create store in batch mode")

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
		t.Attr("package", "mongo")
		t.Attr("component", "link")
		t.Attr("driver", "mongo")

		retrievedLink, err := store.Get(ctx, mock.GetLink.GetHash())
		require.NoError(t, err, "Failed to get Link from store")
		assert.Equal(t, mock.GetLink.GetHash(), retrievedLink.GetHash(), "Hashes should match")
		assert.Equal(t, mock.GetLink.GetDescribe(), retrievedLink.GetDescribe(), "Descriptions should match")
	})

	t.Run("Get list", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mongo")
		t.Attr("component", "link")
		t.Attr("driver", "mongo")

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

	t.Run("Get list using filter", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mongo")
		t.Attr("component", "link")
		t.Attr("driver", "mongo")

		// Set up data needed for the test
		linkBuilder := v1.NewLinkBuilder().
			SetURL("http://example.com/filter").
			SetDescribe("Filter link description")
		link, err := linkBuilder.Build()
		require.NoError(t, err, "Failed to build Link using LinkBuilder")

		_, err = store.Add(ctx, link)
		require.NoError(t, err, "Failed to add Link to store")

		linkNotValid := "https://google.com"
		filter := &filter2.FilterLink{
			Url: &domain.StringFilterInput{
				Eq: link.GetUrl().String(),
				Ne: linkNotValid,
			},
			Hash: &domain.StringFilterInput{Eq: link.GetHash()},
		}
		links, err := store.List(ctx, (*domain.FilterLink)(filter))
		require.NoError(t, err, "Failed to list Links with filter")
		assert.Len(t, links.GetLinks(), 1, "Should have exactly 1 Link")
	})

	t.Run("Delete", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "mongo")
		t.Attr("component", "link")
		t.Attr("driver", "mongo")

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
}

// getLink constructs a new Link using the LinkBuilder.
func getLink() (*v1.Link, error) {
	id := linkUniqId.Add(1)
	url := fmt.Sprintf("http://example.com/%d", id)
	describe := "Generated link description"

	linkBuilder := v1.NewLinkBuilder().
		SetURL(url).
		SetDescribe(describe)
	link, err := linkBuilder.Build()
	if err != nil {
		return nil, err
	}

	return link, nil
}
