//go:build unit || (database && cockroachdb)

package cockroachdb

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

func TestCockroachDB(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "cockroachdb")
	t.Attr("component", "db")
	t.Attr("driver", "cockroachdb")

		t.Attr("type", "unit")
		t.Attr("package", "cockroachdb")
		t.Attr("component", "db")
		t.Attr("driver", "cockroachdb")
	
	ctx, cancel := context.WithCancel(context.Background())
	store := Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "cockroachdb/cockroach",
		Tag:        "v23.1.3",
		Env: []string{
			"COCKROACH_PASSWORD=password",
			"COCKROACH_DATABASE=shortlink",
		},
		Cmd: []string{"start-single-node", "--insecure"},
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
		t.Setenv("STORE_COCKROACHDB_URI", fmt.Sprintf("postgresql://root:password@localhost:%s/shortlink?sslmode=disable", resource.GetPort("26257/tcp"))) // Note that the port has changed

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
