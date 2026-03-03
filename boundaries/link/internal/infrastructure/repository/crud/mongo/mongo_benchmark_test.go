//go:build unit || (database && mongo)

package mongo

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	db "github.com/shortlink-org/go-sdk/db/drivers/mongo"
	"github.com/shortlink-org/go-sdk/db/options"
)

func BenchmarkMongoSerial(b *testing.B) {
	ctx, cancel := context.WithCancelCause(context.Background())
	b.Cleanup(func() { cancel(nil) })

	st := &db.Store{}

	b.Setenv("STORE_MONGODB_URI", startMongoContainer(b))
	require.NoError(b, st.Init(ctx))

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		store, err := New(ctx, st)
		require.NoError(b, err, "Could not create store")

		for b.Loop() {
			source, err := getLink()
			require.NoError(b, err)

			_, err = store.Add(ctx, source)
			require.NoError(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		newCtx, cancelBatch := context.WithCancelCause(ctx)
		defer cancelBatch(nil)

		storeBatchMode, err := New(newCtx, st)
		require.NoError(b, err, "Could not create store in batch mode")

		for b.Loop() {
			source, err := getLink()
			require.NoError(b, err)

			_, err = storeBatchMode.Add(newCtx, source)
			require.NoError(b, err)
		}
	})
}
