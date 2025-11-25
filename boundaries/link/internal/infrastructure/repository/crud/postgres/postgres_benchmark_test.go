//go:build unit || (database && postgres)

package postgres

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	db "github.com/shortlink-org/go-sdk/db/drivers/postgres"
	"github.com/shortlink-org/go-sdk/db/options"
)

func BenchmarkPostgresSerial(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(cancel)

	st := &db.Store{}

	b.Setenv("STORE_POSTGRES_URI", startPostgresContainer(b))
	require.NoError(b, st.Init(ctx))

	// new store
	store, err := New(ctx, st)
	if err != nil {
		b.Fatalf("Could not create store: %s", err)
	}

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		for b.Loop() {
			source, err := getLink()
			require.NoError(b, err)

			_, err = store.Add(ctx, source)
			require.NoError(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		// new store
		storeBatchMode, err := New(ctx, st)
		if err != nil {
			b.Fatalf("Could not create store: %s", err)
		}

		for b.Loop() {
			source, err := getLink()
			require.NoError(b, err)

			_, err = storeBatchMode.Add(ctx, source)
			require.NoError(b, err)
		}
	})
}

func BenchmarkPostgresParallel(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	b.Cleanup(cancel)

	st := &db.Store{}

	b.Setenv("STORE_POSTGRES_URI", startPostgresContainer(b))
	require.NoError(b, st.Init(ctx))

	// new store
	store, err := New(ctx, st)
	if err != nil {
		b.Fatalf("Could not create store: %s", err)
	}

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_SINGLE_WRITE))

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				require.NoError(b, err)

				_, err = store.Add(ctx, source)
				require.NoError(b, err)
			}
		})
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		// new store
		storeBatchMode, err := New(ctx, st)
		if err != nil {
			b.Fatalf("Could not create store: %s", err)
		}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				require.NoError(b, err)

				_, err = storeBatchMode.Add(ctx, source)
				require.NoError(b, err)
			}
		})
	})
}
