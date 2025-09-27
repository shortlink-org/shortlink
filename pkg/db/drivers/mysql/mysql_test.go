//go:build unit || (database && mysql)

package mysql

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

func TestMySQL(t *testing.T) {
	t.Attr("type", "unit")
	t.Attr("package", "mysql")
	t.Attr("component", "db")
	t.Attr("driver", "mysql")

	store := Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.Nil(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "latest", []string{"MYSQL_ROOT_PASSWORD=secret"})
	require.Nil(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.Setenv("STORE_MYSQL_URI", fmt.Sprintf("root:secret@(localhost:%s)/mysql?parseTime=true", resource.GetPort("3306/tcp")))

		err = store.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		require.Nil(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})
}
