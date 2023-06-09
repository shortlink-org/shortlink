//go:build unit || auth

package auth

import (
	"context"
	"embed"
	"fmt"
	"os"
	"testing"
	"testing/fstest"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

var (
	//go:embed permissions/*
	permissions embed.FS
)

// TestGetPermissions tests the GetPermissions function.
func TestGetPermissions(t *testing.T) {
	// Create a mock file system
	mockFS := fstest.MapFS{
		"test1.zed.yaml": &fstest.MapFile{Data: []byte(`
schema: |-
  text: 123
`)},
		"test2.zed.yaml": &fstest.MapFile{Data: []byte(`
schema:
`)},
		"test3.txt": &fstest.MapFile{Data: []byte("content3")},
	}

	permissionsData, err := GetPermissions(mockFS)
	require.NoError(t, err)

	// Expecting 2 files with .zed extension
	require.Len(t, permissionsData, 2)

	// Check the content of the first file
	require.Equal(t, "text: 123", permissionsData[0].Schema)

	// Check the content of the second file
	require.Equal(t, "", permissionsData[1].Schema)
}

func TestSpiceDB(t *testing.T) {
	ctx := context.Background()

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository:   "authzed/spicedb",
		Tag:          "latest",
		Cmd:          []string{"serve-testing"},
		ExposedPorts: []string{"50051/tcp", "50052/tcp"},
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
		var errSetenv error

		errSetenv = os.Setenv("SPICE_DB_API", fmt.Sprintf("localhost:%s", resource.GetPort("50051/tcp")))
		require.NoError(t, errSetenv, "Cannot set ENV")

		errMigrations := Migrations(ctx, permissions)
		require.NoError(t, errMigrations, "Cannot migrate")

		return nil
	}); errRetry != nil {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}

		require.NoError(t, errRetry, "Could not connect to docker")
	}

	t.Cleanup(func() {
		// When you're done, kill and remove the container
		if errPurge := pool.Purge(resource); errPurge != nil {
			t.Fatalf("Could not purge resource: %s", errPurge)
		}
	})
}
