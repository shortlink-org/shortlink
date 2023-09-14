# Auth Package

This Go package provides functionality to migrate permission 
schemas to a [SpiceDB](https://github.com/authzed/spicedb) instance. 
It retrieves schema files (with `.zed` extension) embedded in the application and 
writes them to a SpiceDB server using the Authzed API.

## Usage

This package includes the `Migrations` function, which retrieves permission schemas and writes them to the SpiceDB server.

```go
err := auth.Migrations(context.Background())
if err != nil {
    log.Fatalf("Failed to migrate permissions: %v", err)
}
```

This function reads configurations from environment variables:

| Name                  | Description                         | Default Value                      |
|-----------------------|-------------------------------------|------------------------------------|
| `SPICE_DB_API`        | The address of the SpiceDB API.     | `"shortlink.spicedb-operator:50051"`           |
| `SPICE_DB_COMMON_KEY` | The shared key for the SpiceDB API. | `"secret-shortlink-preshared-key"` |

