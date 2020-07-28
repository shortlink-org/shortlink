package ram

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/batazor/shortlink/internal/store/mock"
	"github.com/batazor/shortlink/internal/store/options"
)

func BenchmarkRAM(b *testing.B) {
	store := RAMLinkList{}

	ctx := context.Background()

	b.Run("Create [single]", func(b *testing.B) {
		// create a store
		err := store.Init(ctx)
		assert.Nil(b, err)

		data := mock.AddLink

		for i := 0; i < b.N; i++ {
			data.Url = data.Url + "/" + string(i)
			_, err := store.Add(ctx, data)
			assert.Nil(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		// Set config
		err := os.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))
		assert.Nil(b, err, "Cannot set ENV")

		// create a store
		err = store.Init(ctx)
		assert.Nil(b, err)

		data := mock.AddLink

		for i := 0; i < b.N; i++ {
			data.Url = data.Url + "/" + string(i)
			_, err := store.Add(ctx, data)
			assert.Nil(b, err)
		}
	})
}
