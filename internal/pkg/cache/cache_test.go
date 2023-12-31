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
	"github.com/prometheus/client_golang/prometheus"
	io_prometheus_client "github.com/prometheus/client_model/go"
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
		c, err := cache.New(ctx)
		require.NoError(t, err)

		key := "myKey"
		value := "myValue"

		err = c.Set(&cache2.Item{
			Key:   key,
			Value: value,
		})
		require.NoError(t, err)

		resp := ""
		err = c.Get(ctx, key, &resp)
		require.NoError(t, err)
		require.Equal(t, value, resp)
	})

	t.Run("Test Set and Get with Prometheus Metrics", func(t *testing.T) {
		ctx := context.Background()
		c, err := cache.New(ctx)
		require.NoError(t, err)

		keyExists := "myKey"
		keyNotExists := "unknownKey"
		value := "myValue"

		// Set operation for existing key
		err = c.Set(&cache2.Item{
			Key:   keyExists,
			Value: value,
		})
		require.NoError(t, err)

		// Get operation where key exists
		resp := ""
		err = c.Get(ctx, keyExists, &resp)
		require.NoError(t, err)
		require.Equal(t, value, resp)

		// Get operation where key does not exist
		resp = ""
		err = c.Get(ctx, keyNotExists, &resp)
		require.NoError(t, err)

		// Gather metrics after operations
		metrics, err := prometheus.DefaultGatherer.Gather()
		require.NoError(t, err)

		// Iterate over gathered metrics and check values inline
		for _, m := range metrics {
			if *m.Name == "cache_operation_duration_seconds" {
				for _, metric := range m.Metric {
					if getValueForLabel(metric.Label, "key") == keyExists || getValueForLabel(metric.Label, "key") == keyNotExists {
						require.GreaterOrEqual(t, *metric.Histogram.SampleSum, 0.0, "Duration should be non-negative")
					}
				}
			}

			if *m.Name == "cache_operation_errors_total" {
				for _, metric := range m.Metric {
					if getValueForLabel(metric.Label, "key") == keyNotExists {
						require.GreaterOrEqual(t, *metric.Counter.Value, 1.0, "Error count should be positive for non-existent key")
					}
				}
			}

			if *m.Name == "cache_operations_total" {
				for _, metric := range m.Metric {
					if getValueForLabel(metric.Label, "key") == keyExists {
						require.GreaterOrEqual(t, *metric.Counter.Value, 1.0, "Operation count should be positive for existing key")
					}
				}
			}
		}
	})
}

func getValueForLabel(labels []*io_prometheus_client.LabelPair, labelName string) string {
	for _, l := range labels {
		if l.GetName() == labelName {
			return l.GetValue()
		}
	}

	return ""
}
