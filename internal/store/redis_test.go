package store

import (
	"fmt"
	"github.com/ory/dockertest"
	"os"
	"testing"
)

func TestRedis(t *testing.T) {
	store := RedisLinkList{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("redis", "latest", nil)
	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_REDIS_URI", fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")))
		if err != nil {
			t.Fatalf("Cannot set ENV: %s", err)
		}

		err = store.Init()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
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

	t.Run("Get", func(t *testing.T) {
		link, err := store.Get(getLink.Hash)
		if err != nil {
			t.Error(err)
		}

		if link.Hash != getLink.Hash {
			t.Errorf("Assert hash - %s; Get %s hash", getLink.Hash, link.Hash)
		}
	})

	t.Run("Get list", func(t *testing.T) {
		links, err := store.List()
		if err != nil {
			t.Error(err)
		}

		if len(links) != 1 {
			t.Errorf("Assert 1 links; Get %d link(s)", len(links))
		}
	})

	t.Run("Delete", func(t *testing.T) {
		err := store.Delete(getLink.Hash)
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
	if err := pool.Purge(resource); err != nil {
		t.Fatalf("Could not purge resource: %s", err)
	}
}
