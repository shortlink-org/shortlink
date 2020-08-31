package scylla

import (
	"fmt"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestScylla(t *testing.T) {
	// db := Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(t, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("scylladb/scylla", "latest", nil)
	assert.Nil(t, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		err = os.Setenv("STORE_SCYLLA_URI", fmt.Sprintf("localhost:%s", resource.GetPort("9042/tcp")))
		assert.Nil(t, err, "Cannot set ENV")

		// if errInit := db.Init(); errInit != nil {
		// 	return errInit
		// }

		return nil
	}); err != nil {
		assert.Nil(t, err, "Could not connect to docker")
	}

	t.Cleanup(func() {
		// When you're done, kill and remove the container
		if err := pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

	// t.Run("Close", func(t *testing.T) {
	// 	assert.Nil(t, db.Close())
	// })
}
