//go:build unit || (database && postgres)
// +build unit database,postgres

package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ory/dockertest/v3"
	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	db "github.com/shortlink-org/shortlink/internal/pkg/db/postgres"
)

func BenchmarkPostgresSerial(b *testing.B) {
	ctx := context.Background()

	st := db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(b, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "latest", []string{
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
		var err error

		err = os.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/shortlink?sslmode=disable", resource.GetPort("5432/tcp")))
		require.NoError(b, err, "Cannot set ENV")

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
		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{
			client: st.GetConn().(*pgxpool.Pool),
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
			client: st.GetConn().(*pgxpool.Pool),
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
	resource, err := pool.Run("postgres", "latest", []string{
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
		var err error

		err = os.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/shortlink?sslmode=disable", resource.GetPort("5432/tcp")))
		require.NoError(b, err, "Cannot set ENV")

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
		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{
			client: st.GetConn().(*pgxpool.Pool),
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
			client: st.GetConn().(*pgxpool.Pool),
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
