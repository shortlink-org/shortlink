package partmap

import (
	"context"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkStd(b *testing.B) {
	b.Attr("type", "unit")
	b.Attr("package", "partmap")
	b.Attr("component", "types")

		b.Attr("type", "unit")
		b.Attr("package", "partmap")
		b.Attr("component", "types")
	

	b.Run("set std concurrently", func(b *testing.B) {
		m := make(map[string]int)
		var wg sync.WaitGroup
		var mu sync.RWMutex

		b.ResetTimer()
		b.ReportAllocs()

		for i := range b.N {
			wg.Add(1)
			go func(i int) {
				key := strconv.Itoa(i)
				mu.Lock()
				m[key] = i
				mu.Unlock()
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}


func BenchmarkSyncStd(b *testing.B) {
	b.Attr("type", "unit")
	b.Attr("package", "partmap")
	b.Attr("component", "types")

		b.Attr("type", "unit")
		b.Attr("package", "partmap")
		b.Attr("component", "types")
	
	b.Run("set sync map std concurrently", func(b *testing.B) {
		var m sync.Map
		var wg sync.WaitGroup

		b.ResetTimer()
		b.ReportAllocs()

		for i := range b.N {
			wg.Add(1)
			go func(i int) {
				key := strconv.Itoa(i)
				m.Store(key, i)
				wg.Done()
			}(i)
		}
		wg.Wait()

	})
}

func BenchmarkPartitioned(b *testing.B) {
	b.Attr("type", "unit")
	b.Attr("package", "partmap")
	b.Attr("component", "types")

		b.Attr("type", "unit")
		b.Attr("package", "partmap")
		b.Attr("component", "types")
	
	m, err := New(&HashSumPartitioner{1000}, 1000)
	if err != nil {
		b.Fatalf("Failed to create PartMap: %v", err)
	}

	b.Run("set partitioned concurrently", func(b *testing.B) {
		var wg sync.WaitGroup

		b.ResetTimer()
		b.ReportAllocs()

		for i := range b.N {
			wg.Add(1)
			go func(i int) {
				key := strconv.Itoa(i)
				if err := m.Set(key, i); err != nil {
					b.Errorf("Failed to set value in PartMap: %v", err)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}
