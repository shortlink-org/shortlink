// +build database,ram

package ram

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/atomic"

	"github.com/batazor/shortlink/internal/api/infrastructure/store/mock"
	"github.com/batazor/shortlink/internal/db/options"
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
			assert.Nil(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		assert.Nil(b, err, "Cannot set ENV")

		data := mock.AddLink

		for i := 0; i < b.N; i++ {
			data.Url = fmt.Sprintf("%s/%d", data.Url, i)
			_, err := store.Add(ctx, data)
			assert.Nil(b, err)
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
				assert.Nil(b, err)

				atom.Inc()
			}
		})
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		assert.Nil(b, err, "Cannot set ENV")

		// create a db
		store := Store{}

		data := mock.AddLink
		var atom atomic.Int64

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				data.Url = fmt.Sprintf("%s/%d", data.Url, atom.Load())
				_, err := store.Add(ctx, data)
				assert.Nil(b, err)

				atom.Inc()
			}
		})
	})
}
