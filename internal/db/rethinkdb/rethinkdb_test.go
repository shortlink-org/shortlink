package rethinkdb

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

func TestRethinkDB(t *testing.T) {
	store := Store{}
	ctx := context.Background()

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("rethinkdb", "2.4-buster-slim", []string{})
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

		err = os.Setenv("STORE_RETHINKDB_URI", fmt.Sprintf("localhost:%s", resource.GetPort("28015/tcp")))
		assert.Nil(t, err, "Cannot set ENV")

		err = store.Init(ctx)
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

	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
