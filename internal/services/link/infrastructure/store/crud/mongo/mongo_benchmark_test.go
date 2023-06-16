//go:build unit || (database && mongo)

package mongo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/shortlink-org/shortlink/internal/pkg/db/mongo"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
)

func BenchmarkPostgresSerial(b *testing.B) {
	ctx := context.Background()

	st := db.Store{}

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
		var err error

		err = os.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp")))
		require.NoError(b, err, "Cannot set ENV")

		err = st.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		require.NoError(b, err, "Could not connect to docker")
	}

	b.Cleanup(func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{
			client: st.GetConn().(*mongo.Client),
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

		// create a db
		storeBatchMode := Store{
			client: st.GetConn().(*mongo.Client),
		}

		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		require.NoError(b, err, "Cannot set ENV")

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			require.NoError(b, err)

			_, err = storeBatchMode.Add(ctx, source)
			require.NoError(b, err)
		}
	})
}

func BenchmarkPostgresParallel(b *testing.B) {
	ctx := context.Background()

	st := db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(b, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mongo", "latest", nil)
	require.NoError(b, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		err = os.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp")))
		require.NoError(b, err, "Cannot set ENV")

		err = st.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		require.NoError(b, err, "Could not connect to docker")
	}

	b.Cleanup(func() {
		// When you're done, kill and remove the container
		if err = pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{
			client: st.GetConn().(*mongo.Client),
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
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		require.NoError(b, err, "Cannot set ENV")

		// create a db
		storeBatchMode := Store{
			client: st.GetConn().(*mongo.Client),
		}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				require.NoError(b, err)

				_, err = storeBatchMode.Add(ctx, source)
				require.NoError(b, err)
			}
		})
	})
}
