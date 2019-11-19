package store

import (
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
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
		// You can't defer this because os.Exit doesn't care for defer
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
		link, err := store.Add(addLink)
		if err != nil {
			t.Error(err)
		}

		if link.Hash != getLink.Hash {
			t.Errorf("Assert hash - %s; Get %s hash", getLink.Hash, link.Hash)
		}
	})

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(ALPHA); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(ZERO); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}
}
