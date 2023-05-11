# Cache

This repository contains the `cache` package, a Go package that provides a high-level interface to cache operations used 
in the **Shortlink** project. This package leverages `go-redis/cache` to provide a comprehensive caching solution 
that includes both local and Redis-based caching.

## Getting Started

This package is designed to be imported and used in other Go applications.

```go
import "github.com/shortlink-org/shortlink/internal/pkg/cache"
```

## Features

- Easy interface to handle cache operations such as `Set`, `SetXX`, `SetNX`, `Get`, and `Del`.
- Uses a hybrid local and Redis-based caching system for enhanced performance and scalability.
- Uses TinyLFU (Least Frequently Used) algorithm for local cache eviction strategy.

## Usage

### Initialization

To create a new instance of the cache client, use the `New` function:

```go
ctx := context.Background()
cacheClient, err := cache.New(ctx)
if err != nil {
    log.Fatal(err)
}
```

### Setting Values

You can add values to the cache using the `Set`, `SetXX`, or `SetNX` functions:

```go
statusCmd := cacheClient.Set(ctx, "key", "value", time.Minute)
```

### Getting Values

Retrieve values from the cache using the `Get` function:

```go
stringCmd := cacheClient.Get(ctx, "key")
value, err := stringCmd.Result()
if err != nil {
    log.Fatal(err)
}
fmt.Println(value)
```

### Deleting Values

You can delete one or more values from the cache using the `Del` function:

```go
intCmd := cacheClient.Del(ctx, "key1", "key2")
```

## Requirements

This package requires Go 1.18 or later. The Go-Redis/Cache and Rueidis libraries are required and will be installed when you install the package.
