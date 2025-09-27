//go:build unit || (database && dgraph)

package dgraph

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/dgraph-io/dgo/v2"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud/mock"
	db "github.com/shortlink-org/shortlink/pkg/db/drivers/dgraph"
)

func TestMain(m *testing.M) {

	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestDgraph(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "dgraph")
	t.Attr("component", "link")
	t.Attr("driver", "dgraph")

	st := db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// create a network with Client.CreateNetwork()
	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
		Name: "shortlink-test",
	})
	if err != nil {
		assert.Errorf(t, err, "Error create docker network")
		os.Exit(1)
	}

	// pulls an image, creates a container based on it and runs it
	ZERO, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "dgraph/dgraph",
		Tag:          "v23.1.0",
		Cmd:          []string{"dgraph", "zero", "--my=test-dgraph-zero:5080"},
		ExposedPorts: []string{"5080"},
		Name:         "test-dgraph-zero",
		NetworkID:    network.ID,
	})
	require.NoError(t, err, "Could not start resource")

	ALPHA, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "dgraph/dgraph",
		Tag:        "v23.1.0",
		Cmd:        []string{"dgraph", "alpha", "--my=localhost:7080", "--lru_mb=2048", fmt.Sprintf("--zero=%s:%s", "test-dgraph-zero", "5080")},
		NetworkID:  network.ID,
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(ZERO); errPurge != nil {
			assert.Errorf(t, errPurge, "Could not purge resource")
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("STORE_DGRAPH_URI", fmt.Sprintf("localhost:%s", ALPHA.GetPort("9080/tcp")))

		errInit := st.Init(ctx)
		if errInit != nil {
			return errInit
		}

		return nil
	}); err != nil {
		assert.Errorf(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(ALPHA); err != nil {
			assert.Errorf(t, err, "Could not purge resource")
		}

		// When you're done, kill and remove the container
		if err := pool.Purge(ZERO); err != nil {
			assert.Errorf(t, err, "Could not purge resource")
		}
	})

	store := Store{
		client: st.GetConn().(*dgo.Dgraph),
	}

	t.Run("Create", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "dgraph")
		t.Attr("component", "link")
		t.Attr("driver", "dgraph")

		link, err := store.Add(ctx, mock.AddLink)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "dgraph")
		t.Attr("component", "link")
		t.Attr("driver", "dgraph")

		link, err := store.Get(ctx, mock.GetLink.Hash)
		require.NoError(t, err)
		assert.Equal(t, link.Hash, mock.GetLink.Hash)
		assert.Equal(t, link.Describe, mock.GetLink.Describe)
	})

	t.Run("Get list", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "dgraph")
		t.Attr("component", "link")
		t.Attr("driver", "dgraph")

		links, err := store.List(ctx, nil)
		require.NoError(t, err)
		assert.Equal(t, len(links.Link), 1)
	})

	t.Run("Delete", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "dgraph")
		t.Attr("component", "link")
		t.Attr("driver", "dgraph")

		require.NoError(t, store.Delete(ctx, mock.GetLink.Hash))
	})
}
