package mongo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"

	db "github.com/batazor/shortlink/internal/db/mongo"
	"github.com/batazor/shortlink/internal/db/options"
)

func BenchmarkPostgresSerial(b *testing.B) {
	ctx := context.Background()

	st := db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(b, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mongo", "latest", nil)
	assert.Nil(b, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error

		err = os.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp")))
		assert.Nil(b, err, "Cannot set ENV")

		err = st.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		assert.Nil(b, err, "Could not connect to docker")
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
			client: st.GetConn().(*mongo.Client),
		}

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			assert.Nil(b, err)

			_, err = store.Add(ctx, source)
			assert.Nil(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		storeBatchMode := Store{
			client: st.GetConn().(*mongo.Client),
		}

		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		assert.Nil(b, err, "Cannot set ENV")

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			assert.Nil(b, err)

			_, err = storeBatchMode.Add(ctx, source)
			assert.Nil(b, err)
		}
	})
}

func BenchmarkPostgresParallel(b *testing.B) {
	ctx := context.Background()

	st := db.Store{}

	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	assert.Nil(b, err, "Could not connect to docker")

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mongo", "latest", nil)
	assert.Nil(b, err, "Could not start resource")

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		err = os.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp")))
		assert.Nil(b, err, "Cannot set ENV")

		err = st.Init(ctx)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		assert.Nil(b, err, "Could not connect to docker")
	}

	b.Cleanup(func() {
		// When you're done, kill and remove the container
		if err = pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{
			client: st.GetConn().(*mongo.Client),
		}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				assert.Nil(b, err)

				_, err = store.Add(ctx, source)
				assert.Nil(b, err)
			}
		})
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		assert.Nil(b, err, "Cannot set ENV")

		// create a db
		storeBatchMode := Store{
			client: st.GetConn().(*mongo.Client),
		}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				assert.Nil(b, err)

				_, err = storeBatchMode.Add(ctx, source)
				assert.Nil(b, err)
			}
		})
	})
}
