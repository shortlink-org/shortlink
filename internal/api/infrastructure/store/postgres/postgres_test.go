package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/atomic"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/api/infrastructure/store/mock"
	"github.com/batazor/shortlink/internal/db/options"
	db "github.com/batazor/shortlink/internal/db/postgres"
)

var linkUniqId atomic.Int64

func TestPostgres(t *testing.T) {
	ctx := context.Background()

	st := db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "latest", []string{
		"POSTGRES_USER=shortlink",
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

		err = os.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://shortlink:shortlink@localhost:%s/shortlink?sslmode=disable", resource.GetPort("5432/tcp")))
		assert.Nil(t, err, "Cannot set ENV")

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

		assert.Nil(t, err, "Could not connect to docker")
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

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Create [batch]", func(t *testing.T) {
		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		assert.Nil(t, err, "Cannot set ENV")

		storeBatchMode := Store{
			client: st.GetConn().(*pgxpool.Pool),
		}

		link, err := storeBatchMode.Add(ctx, getLink())
		assert.Nil(t, err)
		assert.NotNil(t, link.CreatedAt)

		link, err = storeBatchMode.Add(ctx, getLink())
		assert.Nil(t, err)
		assert.NotNil(t, link.CreatedAt)

		link, err = storeBatchMode.Add(ctx, getLink())
		assert.Nil(t, err)
		assert.NotNil(t, link.CreatedAt)

		link, err = storeBatchMode.Add(ctx, getLink())
		assert.Nil(t, err)
		assert.NotNil(t, link.CreatedAt)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		assert.Nil(t, err)
		assert.Equal(t, len(links), 8)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.Nil(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}

func getLink() *link.Link {
	link, _ := link.NewURL(fmt.Sprintf("%s/%d", "http://example.com", linkUniqId.Load()))
	linkUniqId.Inc()
	return link
}
