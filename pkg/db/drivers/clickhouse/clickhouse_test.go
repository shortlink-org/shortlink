//go:build unit || (database && clickhouse)

package clickhouse

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

func TestClickHouse(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "clickhouse")
	t.Attr("component", "db")
	t.Attr("driver", "clickhouse")

		t.Attr("type", "unit")
		t.Attr("package", "clickhouse")
		t.Attr("component", "db")
		t.Attr("driver", "clickhouse"), cancel := context.WithCancel(t.Context())
	store := Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("clickhouse/clickhouse-server", "latest", nil)
	require.NoError(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("STORE_CLICKHOUSE_URI", fmt.Sprintf("clickhouse://localhost:%s/default?sslmode=disable", resource.GetPort("9000/tcp")))

		err = store.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		require.NoError(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	t.Run("Close", func(t *testing.T) {
		t.Attr("type", "unit")
		t.Attr("package", "clickhouse")
		t.Attr("component", "db")
		t.Attr("driver", "clickhouse")

		cancel()
	})
}
