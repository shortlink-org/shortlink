//go:build unit || (media && s3)

package s3

import (
	"context"
	"fmt"
	"net/http"
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

func TestMinio(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	client := &Client{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "Could not connect to docker")

	options := &dockertest.RunOptions{
		Repository: "minio/minio",
		Tag:        "RELEASE.2023-12-23T07-19-11Z",
		Cmd:        []string{"server", "--address", ":9000", "/data"},
		Env: []string{
			"MINIO_ROOT_USER=minio_access_key",
			"MINIO_ROOT_PASSWORD=minio_secret_key",
		},
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(options)
	require.NoError(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	endpoint := fmt.Sprintf("localhost:%s", resource.GetPort("9000/tcp"))
	err = os.Setenv("S3_ENDPOINT", fmt.Sprintf("localhost:%s", resource.GetPort("9000/tcp")))
	require.NoError(t, err, "Cannot set ENV")

	if err := pool.Retry(func() error {
		url := fmt.Sprintf("http://%s/minio/health/live", endpoint)

		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status code not OK")
		}

		client, err = New(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		require.NoError(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		cancel()

		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}

		// drop downloaded file
		err := os.Remove("./fixtures/download.json")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("UploadFile", func(t *testing.T) {
		err := client.CreateBucket(ctx, "test")
		if err != nil {
			t.Fatal(err)
		}

		err = client.UploadFile(ctx, "test", "test", "./fixtures/test.json")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DownloadFile", func(t *testing.T) {
		err := client.DownloadFile(ctx, "test", "test", "./fixtures/download.json")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListFiles", func(t *testing.T) {
		files, err := client.ListFiles(ctx, "test")
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, []string{"test"}, files)
	})

	t.Run("FileExists", func(t *testing.T) {
		exists, err := client.FileExists(ctx, "test", "test")
		if err != nil {
			t.Fatal(err)
		}

		require.Equal(t, true, exists)
	})

	t.Run("DeleteFile", func(t *testing.T) {
		err := client.RemoveFile(ctx, "test", "test")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("FileNoExists", func(t *testing.T) {
		exists, err := client.FileExists(ctx, "test", "test")
		// The specified key does not exist
		if err != nil {
			require.Equal(t, "The specified key does not exist.", err.Error())
		}

		require.Equal(t, false, exists)
	})

	t.Run("RemoveBucket", func(t *testing.T) {
		err := client.RemoveBucket(ctx, "test")
		if err != nil {
			t.Fatal(err)
		}
	})
}
