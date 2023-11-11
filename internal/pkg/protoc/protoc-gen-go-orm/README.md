# protoc-gen-go-orm

Protoc plugin for generating Go ORM code from Protocol Buffers (.proto files). 
This tool is designed to simplify the process of creating database interaction layers in Go applications 
by automatically generating ORM-like structures based on your protobuf definitions.

### Features

- Generates Go structs with filter and pagination capabilities based on protobuf messages.
- Easily integrates with existing Protocol Buffer and Go projects.

### Installation

```bash
go install github.com/shortlink-org/shortlink/internal/pkg/protoc/protoc-gen-go-orm
```

### Usage with buf

```yaml
version: v1

managed:
  enabled: true

plugins:
  - plugin: go-orm
    strategy: all
    out: internal/services/link
    opt: paths=source_relative
```
