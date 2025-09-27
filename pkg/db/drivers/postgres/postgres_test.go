//go:build unit || (database && postgres)

package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)

	os.Exit(m.Run())
}

func TestPostgres(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "postgres")
	t.Attr("component", "db")
	t.Attr("driver", "postgres")

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	
	store := Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "16.0-alpine", []string{
		"POSTGRES_USER=postgres",
		"POSTGRES_PASSWORD=shortlink",
		"POSTGRES_DB=shortlink",
	})
	if err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		t.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if errRetry := pool.Retry(func() error {
		t.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:shortlink@localhost:%s/shortlink?sslmode=disable", resource.GetPort("5432/tcp")))

		errInit := store.Init(ctx)
		if errInit != nil {
			return errInit
		}

		return nil
	}); errRetry != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		require.NoError(t, errRetry, "Could not connect to docker")
	}

	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}
	})
}
