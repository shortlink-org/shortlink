# protoc-gen-rich-model

Protoc plugin for generating Go rich model code from Protocol Buffers (.proto files).
This tool is designed to simplify the process of creating rich model structures in Go applications
by automatically generating rich model-like structures based on your protobuf definitions.

### Features

- Generate DDD-like Go structs with rich model capabilities based on protobuf messages.
- **Whitelist Types**: Allows specifying a list of types to generate rich models for.
- **Support additional types**:
  - **url.URL**

### Installation

```bash
go install github.com/shortlink-org/shortlink/pkg/protoc/protoc-gen-rich-model

# for development
go install ./pkg/protoc/protoc-gen-rich-model
```

### Usage with buf

```yaml
version: v1

managed:
  enabled: true

plugins:
  - plugin: rich-model
    strategy: all
    out: boundaries/link/link
    opt:
      - "paths=source_relative"
      - "filter=Link;Links"
```
