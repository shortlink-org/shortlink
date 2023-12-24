## PartMap

**PartMap** is a concurrent map implementation in Go, designed to reduce lock contention and improve performance in multithreaded environments. 
This package offers a thread-safe map by partitioning the map into several segments, each protected by its own lock.

### Features

- **Concurrency Safe**: Utilizes multiple locks to allow concurrent access with minimal contention.
- **Customizable Partitioning**: Supports custom partition strategies through the `partitioner` interface.
- **Standard Map Interface**: Familiar methods such as `Get`, `Set`, `Delete`, and `Len`.

### Usage

```go
package main

import (
  "fmt"
  "log"

  "github.com/shortlink-org/shortlink/internal/pkg/types/partmap"
)

func main() {
  m, err := partmap.New(&hashSumPartitioner{1000}, 1000)
  if err != nil {
    log.Fatalf("Failed to create PartMap: %v", err)
  }

  // Set a value
  m.Set("key1", "value1")

  // Get a value
  if val, ok := m.Get("key1"); ok {
      fmt.Println("Value:", val)
  }

  // Delete a value
  m.Delete("key1")

  // Get the length
  fmt.Println("Length:", m.Len())
}
```

### Benchmarks

> [!NOTE]
> 
> goos: darwin
> goarch: amd64
> cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz

```shell
BenchmarkStd
BenchmarkStd/set_std_concurrently
BenchmarkStd/set_std_concurrently-16         	  220958	      6605 ns/op	     322 B/op	       4 allocs/op
BenchmarkSyncStd
BenchmarkSyncStd/set_sync_map_std_concurrently
BenchmarkSyncStd/set_sync_map_std_concurrently-16         	  132684	     26067 ns/op	     506 B/op	      10 allocs/op
BenchmarkPartitioned
BenchmarkPartitioned/set_partitioned_concurrently
BenchmarkPartitioned/set_partitioned_concurrently-16      	  251509	      6586 ns/op	     329 B/op	       6 allocs/op
```

### References

- [Writing a Partitioned Cache Using Go Map (x3 Faster than the Standard Map)](https://blog.stackademic.com/writing-a-partitioned-cache-using-go-map-x3-faster-than-the-standard-map-dbfe704fe4bf)
  - [repo](https://github.com/vadiminshakov/partmap/blob/main/crcpartitioner.go)
