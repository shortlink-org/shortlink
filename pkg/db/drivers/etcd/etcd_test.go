//go:build unit || (database && etcd)

package etcd

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {

	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestETCD(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "etcd")
	t.Attr("component", "db")
	t.Attr("driver", "etcd")

	store := Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "docker.io/bitnami/etcd",
		Tag:        "3",
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	require.NoError(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("STORE_ETCD_URI", fmt.Sprintf("localhost:%s", resource.GetPort("2379/tcp")))

		errInit := store.Init(ctx)
		if errInit != nil {
			return errInit
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
}
