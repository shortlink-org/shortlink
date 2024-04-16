# protoc-gen-go-orm

Protoc plugin for generating Go ORM code from Protocol Buffers (.proto files). 
This tool is designed to simplify the process of creating database interaction layers in Go applications 
by automatically generating ORM-like structures based on your protobuf definitions.

### Features

- **Go Struct Generation**: Automatically generates Go structs with embedded ORM functionality tailored to specific databases (PostgreSQL, MongoDB).
- **Filter and Pagination**: Supports generating code for filtering and pagination to enhance database operations.
- **Multiple Database Support**: Includes options for generating ORM code for PostgreSQL, MongoDB, allowing flexibility based on project requirements.
- **Whitelist Types**: Allows specifying which types to generate ORM code for, enabling selective generation of ORM code.

### Installation

```bash
go install github.com/shortlink-org/shortlink/pkg/protoc/protoc-gen-go-orm
```

### Usage with buf

```yaml
version: v1

managed:
  enabled: true

plugins:
  - plugin: go-orm
    out: boundaries/link/link
    opt:
      - "orm=postgres,pkg=package_name,filter=Link"
      - "paths=source_relative"
```
