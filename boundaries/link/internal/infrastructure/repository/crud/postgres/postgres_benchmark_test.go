//go:build unit || (database && postgres)

package postgres

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"

	db "github.com/shortlink-org/shortlink/pkg/db/drivers/postgres"
	"github.com/shortlink-org/shortlink/pkg/db/options"
)

func BenchmarkPostgresSerial(b *testing.B) {
	b.Attr("type", "benchmark")
	b.Attr("package", "postgres")
	b.Attr("component", "link")
	b.Attr("driver", "postgres")

	st := &db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(b, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("ghcr.io/dbsystel/postgresql-partman", "16", []string{
		"POSTGRES_USER=postgres",
		"POSTGRES_PASSWORD=shortlink",
		"POSTGRES_DB=shortlink",
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			b.Fatalf("Could not purge resource: %s", errPurge)
		}

		b.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		b.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/shortlink?sslmode=disable", resource.GetPort("5432/tcp")))

		errInit := st.Init(ctx)
		if errInit != nil {
			return errInit
		}

		return nil
	}); err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			b.Fatalf("Could not purge resource: %s", errPurge)
		}

		require.NoError(b, err, "Could not connect to docker")
	}

	b.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	// new store
	store, err := New(ctx, st)
	if err != nil {
		b.Fatalf("Could not create store: %s", err)
	}

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

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

		// new store
		storeBatchMode, err := New(ctx, st)
		if err != nil {
			b.Fatalf("Could not create store: %s", err)
		}

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			require.NoError(b, err)

			_, err = storeBatchMode.Add(ctx, source)
			require.NoError(b, err)
		}
	})
}

func BenchmarkPostgresParallel(b *testing.B) {
	b.Attr("type", "benchmark")
	b.Attr("package", "postgres")
	b.Attr("component", "link")
	b.Attr("driver", "postgres")

	st := &db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(b, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("ghcr.io/dbsystel/postgresql-partman", "16", []string{
		"POSTGRES_USER=postgres",
		"POSTGRES_PASSWORD=shortlink",
		"POSTGRES_DB=shortlink",
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			b.Fatalf("Could not purge resource: %s", errPurge)
		}

		b.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		b.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/shortlink?sslmode=disable", resource.GetPort("5432/tcp")))

		err = st.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			b.Fatalf("Could not purge resource: %s", errPurge)
		}

		require.NoError(b, err, "Could not connect to docker")
	}

	b.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	// new store
	store, err := New(ctx, st)
	if err != nil {
		b.Fatalf("Could not create store: %s", err)
	}

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_SINGLE_WRITE))

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

		// new store
		storeBatchMode, err := New(ctx, st)
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
	})
}
