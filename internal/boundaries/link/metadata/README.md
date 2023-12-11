# Metadata

Service got get metadata by URL.

### DDD

- application (Write business logic)
- domain (Entity struct that represent mapping to data model)
- infrastructure (Solves backend technical topics)

### Example request

```
grpcurl -cacert ./ops/cert/intermediate_ca.pem -d '{"Id": "http://google.com"}' localhost:50052 metadata_rpc.Metadata/Set
```

### Use Case

**Use cases**:

- [UC-1](./application/parsers/README.md) Parse metadata from URL
