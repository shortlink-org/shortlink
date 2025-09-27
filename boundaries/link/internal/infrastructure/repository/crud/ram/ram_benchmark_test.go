//go:build unit || (database && ram)

package ram

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"

	v1 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/domain/link/v1"
	"github.com/shortlink-org/shortlink/pkg/db/options"
)

var linkUniqId atomic.Int64

func BenchmarkRAMSerial(b *testing.B) {
	b.Attr("type", "unit")
	b.Attr("package", "ram")
	b.Attr("component", "link")
	b.Attr("driver", "ram")

		b.Attr("type", "unit")
		b.Attr("package", "ram")
		b.Attr("component", "link")
		b.Attr("driver", "ram"), cancel := context.WithCancel(t.Context())
	defer cancel()

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			require.NoError(b, err)

			_, err = store.Add(ctx, source)
			require.NoError(b, err)
		}
	})

	b.Run("Create [batch]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

		// Set config
		b.Setenv("STORE_MODE_WRITE", strconv.Itoa(options.MODE_BATCH_WRITE))

		for i := 0; i < b.N; i++ {
			source, err := getLink()
			require.NoError(b, err)

			_, err = store.Add(ctx, source)
			require.NoError(b, err)
		}
	})
}

func BenchmarkRAMParallel(b *testing.B) {
	b.Attr("type", "unit")
	b.Attr("package", "ram")
	b.Attr("component", "link")
	b.Attr("driver", "ram")

		b.Attr("type", "unit")
		b.Attr("package", "ram")
		b.Attr("component", "link")
		b.Attr("driver", "ram")

		ctx := t.Context()

	b.Run("Create [single]", func(b *testing.B) {
		b.ReportAllocs()

		// create a db
		store := Store{}

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

		// create a db
		store := Store{}

		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				source, err := getLink()
				require.NoError(b, err)

				_, err = store.Add(ctx, source)
				require.NoError(b, err)
			}
		})
	})
}

// getLink constructs a new Link using the LinkBuilder.
func getLink() (*v1.Link, error) {
	id := linkUniqId.Add(1)
	url := fmt.Sprintf("http://example.com/%d", id)
	describe := "Generated link description"

	linkBuilder := v1.NewLinkBuilder().
		SetURL(url).
		SetDescribe(describe)
	link, err := linkBuilder.Build()
	if err != nil {
		return nil, err
	}

	return link, nil
}
