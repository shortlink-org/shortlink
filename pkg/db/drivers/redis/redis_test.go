//go:build unit || (database && redis)

package redis

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestRedis(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "redis")
	t.Attr("component", "db")
	t.Attr("driver", "redis")

		t.Attr("type", "unit")
		t.Attr("package", "redis")
		t.Attr("component", "db")
		t.Attr("driver", "redis")
	
	ctx, cancel := context.WithCancel(context.Background())
	store := Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "redis",
		Tag:        "7-alpine",
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	require.NoError(t, err, "Could not start resource")

	// setting the max wait time for the container to start
	pool.MaxWait = time.Minute * 5

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		t.Setenv("STORE_REDIS_URI", fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")))

		errInit := store.Init(ctx)
		if errInit != nil {
			return errInit
		}

		return nil
	})
	require.NoError(t, err, "Could not connect to Docker")

	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})
}
