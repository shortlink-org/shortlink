//go:build unit || (database && ram)
// +build unit database,ram

package ram

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	"github.com/shortlink-org/shortlink/internal/pkg/db/options"
	"github.com/shortlink-org/shortlink/internal/services/link/infrastructure/store/crud/mock"
)

func BenchmarkRAMSerial(b *testing.B) {
	ctx := context.Background()

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		data := mock.AddLink

		for i := 0; i < b.N; i++ {
			data.Url = fmt.Sprintf("%s/%d", data.Url, i)
			_, err := store.Add(ctx, data)
			require.NoError(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		require.NoError(b, err, "Cannot set ENV")

		data := mock.AddLink

		for i := 0; i < b.N; i++ {
			data.Url = fmt.Sprintf("%s/%d", data.Url, i)
			_, err := store.Add(ctx, data)
			require.NoError(b, err)
		}
	})
}

func BenchmarkRAMParallel(b *testing.B) {
	ctx := context.Background()

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		data := mock.AddLink
		var atom atomic.Int64

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.Url = fmt.Sprintf("%s/%d", data.Url, atom.Load())
				_, err := store.Add(ctx, data)
				require.NoError(b, err)

				atom.Inc()
			}
		})
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		require.NoError(b, err, "Cannot set ENV")

		// create a db
		store := Store{}

		data := mock.AddLink
		var atom atomic.Int64

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.Url = fmt.Sprintf("%s/%d", data.Url, atom.Load())
				_, err := store.Add(ctx, data)
				require.NoError(b, err)

				atom.Inc()
			}
		})
	})
}
