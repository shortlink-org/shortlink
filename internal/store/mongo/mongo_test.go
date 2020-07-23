package mongo

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/store/mock"
	"github.com/batazor/shortlink/internal/store/query"
)

// TODO: Problem with testing into GitLab CI
//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

func TestMongo(t *testing.T) {
	store := MongoLinkList{}
	ctx := context.Background()

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mongo", "latest", nil)
	assert.Nil(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s", resource.GetPort("27017/tcp")))
		assert.Nil(t, err, "Cannot set ENV")

		err = store.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		assert.Nil(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		assert.Nil(t, err)
		assert.Equal(t, len(links), 1)
	})

	t.Run("Get list using filter", func(t *testing.T) {
		linkNotValid := "https://google.com"
		filter := &query.Filter{
			Url: &query.StringFilterInput{
				Eq: &mock.GetLink.Url,
				Ne: &linkNotValid,
			},
			Hash: &query.StringFilterInput{Eq: &mock.GetLink.Hash},
		}
		links, err := store.List(ctx, filter)
		assert.Nil(t, err)
		assert.Equal(t, len(links), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.Nil(t, store.Delete(ctx, mock.GetLink.Hash))
	})

	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
