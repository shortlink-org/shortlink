package dgraph

import (
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"

	"github.com/batazor/shortlink/internal/store/mock"
)

func TestDgraph(t *testing.T) {
	store := DGraphLinkList{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	// create a network with Client.CreateNetwork()
	network, err := pool.Client.CreateNetwork(docker.CreateNetworkOptions{
		Name: "shortlink-test",
	})
	if err != nil {
		t.Errorf("Error create docker network: %s", err)
		os.Exit(1)
	}

	// pulls an image, creates a container based on it and runs it
	ZERO, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "dgraph/dgraph",
		Tag:          "latest",
		Cmd:          []string{"dgraph", "zero", "--my=test-dgraph-zero:5080"},
		ExposedPorts: []string{"5080"},
		Name:         "test-dgraph-zero",
		NetworkID:    network.ID,
	})
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}

	ALPHA, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "dgraph/dgraph",
		Tag:        "latest",
		Cmd:        []string{"dgraph", "alpha", "--my=localhost:7080", "--lru_mb=2048", fmt.Sprintf("--zero=%s:%s", "test-dgraph-zero", "5080")},
		NetworkID:  network.ID,
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(ZERO); errPurge != nil {
			t.Errorf("Could not purge resource: %s", errPurge)
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_DGRAPH_URI", fmt.Sprintf("localhost:%s", ALPHA.GetPort("9080/tcp")))
		if err != nil {
			t.Errorf("Cannot set ENV: %s", err)
			return nil
		}

		err = store.Init()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		t.Errorf("Could not connect to docker: %s", err)
	}

	t.Run("Create", func(t *testing.T) {
		link, err := store.Add(mock.AddLink)
		if err != nil {
			t.Error(err)
		}

		if link == nil {
			t.Fatalf("Assert link; Get nil")
		}

		if link.Hash != mock.GetLink.Hash {
			t.Errorf("Assert hash - %s; Get %s hash", mock.GetLink.Hash, link.Hash)
		}
	})

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(mock.GetLink.Hash)
		if err != nil {
			t.Error(err)
		}

		if link == nil {
			t.Fatalf("Assert link; Get nil")
		}

		if link.Hash != mock.GetLink.Hash {
			t.Errorf("Assert hash - %s; Get %s hash", mock.GetLink.Hash, link.Hash)
		}
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List(nil)
		if err != nil {
			t.Error(err)
		}

		if len(links) != 1 {
			t.Errorf("Assert 1 links; Get %d link(s)", len(links))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		link, err := store.Add(mock.AddLink)
		if err != nil {
			t.Error(err)
		}

		err = store.Delete(link.Hash)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Close", func(t *testing.T) {
		err := store.Close()
		if err != nil {
			t.Error(err)
		}
	})

	// When you're done, kill and remove the container
	if err := pool.Purge(ALPHA); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}

	// When you're done, kill and remove the container
	if err := pool.Purge(ZERO); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}
}
