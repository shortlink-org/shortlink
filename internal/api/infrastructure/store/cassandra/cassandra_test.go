package cassandra

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gocql/gocql"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/api/infrastructure/store/mock"
	db "github.com/batazor/shortlink/internal/db/cassandra"
)

//func TestMain(m *testing.M) {
//	goleak.VerifyTestMain(m)
//}

func TestCassandra(t *testing.T) {
	ctx := context.Background()
	st := db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("cassandra", "latest", nil)
	assert.Nil(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		err = os.Setenv("STORE_CASSANDRA_URI", fmt.Sprintf("localhost:%s", resource.GetPort("9042/tcp")))
		assert.Nil(t, err, "Cannot set ENV")

		err = st.Init(ctx)
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

	store := Store{
		client: st.GetConn().(*gocql.Session),
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(ctx, mock.AddLink)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(ctx, mock.GetLink.Hash)
		assert.Nil(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(ctx, nil)
		assert.Nil(t, err)
		assert.Equal(t, len(links), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		assert.Nil(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}
