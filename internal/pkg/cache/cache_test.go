//go:build unit || (database && redis)

package cache_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	cache2 "github.com/go-redis/cache/v9"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/internal/pkg/cache"
)

func TestCache(t *testing.T) {
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
		errSetenv := os.Setenv("STORE_REDIS_URI", fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")))
		require.NoError(t, errSetenv, "Cannot set ENV")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, errCache := cache.New(ctx)
		if errCache != nil {
			return errCache
		}
		return nil
	})
	require.NoError(t, err, "Could not connect to Docker")

	t.Cleanup(func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	t.Run("Test Set and Get", func(t *testing.T) {
		ctx := context.Background()
		c, _ := cache.New(ctx)

		key := "myKey"
		value := "myValue"

		err := c.Set(&cache2.Item{
			Key:   key,
			Value: value,
		})
		require.NoError(t, err)

		resp := ""
		err = c.Get(ctx, key, &resp)
		require.NoError(t, err)
		require.Equal(t, value, resp)
	})
}
