package postgres

import (
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest"
)

func TestPostgres(t *testing.T) {
	store := PostgresLinkList{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("postgres", "latest", []string{
		"POSTGRES_USER=postgres",
		"POSTGRES_PASSWORD=postgres",
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
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_POSTGRES_URI", fmt.Sprintf("postgres://postgres:postgres@localhost:%s/shortlink?sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			t.Fatalf("Cannot set ENV: %s", err)
		}

		err = store.Init()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		t.Fatalf("Could not connect to docker: %s", err)
	}

	// When you're done, kill and remove the container
	if err := pool.Purge(resource); err != nil {
		t.Fatalf("Could not purge resource: %s", err)
	}
}
