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

	db "github.com/shortlink-org/shortlink/internal/pkg/db/mongo"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	"github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/repository/crud/mock"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m, goleak.IgnoreTopFunction("github.com/golang/glog.(*fileSink).flushDaemon"))

	os.Exit(m.Run())
}

var linkUniqId atomic.Int64

func TestMongo(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	st := &db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	require.NoError(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp")))
		require.NoError(t, err, "Cannot set ENV")

		err = st.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		require.NoError(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	// new store
	store, err := New(ctx, st)
	if err != nil {
		t.Fatalf("Could not create store: %s", err)
	}

	t.Run("Create [single]", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Create [batch]", func(t *testing.T) {
		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		require.NoError(t, err, "Cannot set ENV")

		newCtx, cancel := context.WithCancel(ctx)

		storeBatchMode, err := New(newCtx, st)
		if err != nil {
			t.Fatalf("Could not create store: %s", err)
		}

		source, err := getLink()
		require.NoError(t, err)
		_, err = storeBatchMode.Add(ctx, source)
		require.NoError(t, err)
		assert.NotNil(t, source.CreatedAt)
		assert.Equal(t, source.Describe, mock.GetLink.Describe)

		source, err = getLink()
		require.NoError(t, err)
		_, err = storeBatchMode.Add(ctx, source)
		require.NoError(t, err)
		assert.NotNil(t, source.CreatedAt)
		assert.Equal(t, source.Describe, mock.GetLink.Describe)

		source, err = getLink()
		require.NoError(t, err)
		_, err = storeBatchMode.Add(ctx, source)
		require.NoError(t, err)
		assert.NotNil(t, source.CreatedAt)
		assert.Equal(t, source.Describe, mock.GetLink.Describe)

		source, err = getLink()
		require.NoError(t, err)
		_, err = storeBatchMode.Add(ctx, source)
		require.NoError(t, err)
		assert.NotNil(t, source.CreatedAt)
		assert.Equal(t, source.Describe, mock.GetLink.Describe)

		t.Cleanup(func() {
			cancel()
		})
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 5)
	})

	t.Run("Get list using filter", func(t *testing.T) {
		linkNotValid := "https://google.com"
		filter := &v1.FilterLink{
			Url: &v1.StringFilterInput{
				Eq: mock.GetLink.Url,
				Ne: linkNotValid,
			},
			Hash: &v1.StringFilterInput{Eq: mock.GetLink.Hash},
		}
		links, err := store.List(ctx, filter)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}

func getLink() (*v1.Link, error) {
	id := linkUniqId.Add(1)

	data := &v1.Link{
		Url:      fmt.Sprintf("%s/%d", "http://example.com", id),
		Describe: mock.AddLink.Describe,
	}

	if err := v1.NewURL(data); err != nil {
		return nil, err
	}

	return data, nil
}
