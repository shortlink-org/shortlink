//go:build unit || (database && mongo)

package mongo

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"

	db "github.com/shortlink-org/shortlink/internal/pkg/db/mongo"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
)

func BenchmarkMongoSerial(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())

	st := &db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(b, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	require.NoError(b, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		b.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp")))

		errInit := st.Init(ctx)
		if errInit != nil {
			return errInit
		}

		return nil
	}); err != nil {
		require.NoError(b, err, "Could not connect to docker")
	}

	b.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store, err := New(ctx, st)
		if err != nil {
			b.Fatalf("Could not create store: %s", err)
		}

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			require.NoError(b, err)

			_, err = store.Add(ctx, source)
			require.NoError(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		newCtx, cancel := context.WithCancel(ctx)

		// create a db
		storeBatchMode, errNewBatchMode := New(newCtx, st)
		if errNewBatchMode != nil {
			b.Fatalf("Could not create store: %s", errNewBatchMode)
		}

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			require.NoError(b, err)

			_, err = storeBatchMode.Add(ctx, source)
			require.NoError(b, err)
		}

		b.Cleanup(func() {
			cancel()
		})
	})
}

func BenchmarkMongoParallel(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())

	st := &db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(b, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mongo", "latest", nil)
	require.NoError(b, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		b.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp")))

		errInit := st.Init(ctx)
		if errInit != nil {
			return errInit
		}

		return nil
	}); err != nil {
		require.NoError(b, err, "Could not connect to docker")
	}

	b.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err = pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store, err := New(ctx, st)
		if err != nil {
			b.Fatalf("Could not create store: %s", err)
		}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				require.NoError(b, err)

				_, err = store.Add(ctx, source)
				require.NoError(b, err)
			}
		})
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		newCtx, cancel := context.WithCancel(ctx)

		// create a db
		storeBatchMode, err := New(newCtx, st)
		if err != nil {
			b.Fatalf("Could not create store: %s", err)
		}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				require.NoError(b, err)

				_, err = storeBatchMode.Add(ctx, source)
				require.NoError(b, err)
			}
		})

		b.Cleanup(func() {
			cancel()
		})
	})
}
