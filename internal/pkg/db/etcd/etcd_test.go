//go:build unit || (database && etcd)
// +build unit database,etcd

package etcd

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

func TestETCD(t *testing.T) {
	store := Store{}
	ctx := context.Background()

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("docker.io/bitnami/etcd", "3", nil)
	assert.Nil(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_ETCD_URI", fmt.Sprintf("localhost:%s", resource.GetPort("2379/tcp")))
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

	t.Run("Close", func(t *testing.T) {
		assert.Nil(t, store.Close())
	})
}
