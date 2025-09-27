//go:build unit || (database && mongo)

package mongo

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"

	db "github.com/shortlink-org/shortlink/pkg/db/drivers/mongo"
	"github.com/shortlink-org/shortlink/pkg/db/options"
)

func BenchmarkMongoSerial(b *testing.B) {
	b.Attr("type", "unit")
	b.Attr("package", "mongo")
	b.Attr("component", "link")
	b.Attr("driver", "mongo")

		b.Attr("type", "unit")
		b.Attr("package", "mongo")
		b.Attr("component", "link")
		b.Attr("driver", "mongo"), cancel := context.WithCancel(t.Context())
	defer cancel()

	st := &db.Store{}

	pool, err := dockertest.NewPool("")
	require.NoError(b, err, "Could not connect to docker")

	resource, err := pool.Run("mongo", "7.0", nil)
	require.NoError(b, err, "Could not start resource")

	if err := pool.Retry(func() error {
		b.Setenv("STORE_MONGODB_URI", fmt.Sprintf("mongodb://localhost:%s/shortlink", resource.GetPort("27017/tcp")))
		return st.Init(ctx)
	}); err != nil {
		require.NoError(b, err, "Could not connect to docker")
	}

	b.Cleanup(func() {
		if err := pool.Purge(resource); err != nil {
			b.Fatalf("Could not purge resource: %s", err)
		}
	})

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		store, err := New(ctx, st)
		require.NoError(b, err, "Could not create store")

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			require.NoError(b, err)

			_, err = store.Add(ctx, source)
			require.NoError(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		newCtx, cancelBatch := context.WithCancel(ctx)
		defer cancelBatch()

		storeBatchMode, err := New(newCtx, st)
		require.NoError(b, err, "Could not create store in batch mode")

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			require.NoError(b, err)

			_, err = storeBatchMode.Add(newCtx, source)
			require.NoError(b, err)
		}
	})
}
