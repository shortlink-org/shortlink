package partmap

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkStd(b *testing.B) {
	b.Run("set std concurrently", func(b *testing.B) {
		m := make(map[string]int)
		var wg sync.WaitGroup
		var mu sync.RWMutex

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(i int) {
				key := fmt.Sprint(i)
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
	b.Run("set sync map std concurrently", func(b *testing.B) {
		var m sync.Map
		var wg sync.WaitGroup

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(i int) {
				key := fmt.Sprint(i)
				m.Store(key, i)
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}

func BenchmarkPartitioned(b *testing.B) {
	m, err := New(&HashSumPartitioner{1000}, 1000)
	if err != nil {
		b.Fatalf("Failed to create PartMap: %v", err)
	}

	b.Run("set partitioned concurrently", func(b *testing.B) {
		var wg sync.WaitGroup

		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(i int) {
				key := fmt.Sprint(i)
				if err := m.Set(key, i); err != nil {
					b.Errorf("Failed to set value in PartMap: %v", err)
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}
