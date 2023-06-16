//go:build unit || (database && postgres)

package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	db "github.com/shortlink-org/shortlink/internal/pkg/db/postgres"
	"github.com/shortlink-org/shortlink/internal/services/link/domain/link/v1"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/mock"
)

var linkUniqId atomic.Int64

func TestPostgres(t *testing.T) {
	ctx := context.Background()

	st := db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "latest", []string{
		"POSTGRES_USER=postgres",
		"POSTGRES_PASSWORD=shortlink",
		"POSTGRES_DB=shortlink",
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/shortlink?sslmode=disable", resource.GetPort("5432/tcp")))
		require.NoError(t, err, "Cannot set ENV")

		err = st.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		require.NoError(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	store := Store{
		client: st.GetConn().(*pgxpool.Pool),
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

		storeBatchMode := Store{
			client: st.GetConn().(*pgxpool.Pool),
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
		assert.Equal(t, len(links.Link), 8)
	})

	t.Run("Delete", func(t *testing.T) {
		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}

func getLink() (*v1.Link, error) {
	data := &v1.Link{
		Url:      fmt.Sprintf("%s/%d", "http://example.com", linkUniqId.Load()),
		Describe: mock.AddLink.Describe,
	}
	if err := v1.NewURL(data); err != nil {
		return nil, err
	}
	linkUniqId.Inc()
	return data, nil
}
