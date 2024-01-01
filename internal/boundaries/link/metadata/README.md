# Metadata

> [!NOTE]
> This ADR introduces a service to the shortlink system, designed to enrich link information.
> Upon the creation of a new link, this service fetches metadata and generates screenshots,
> subsequently storing this data to provide a more detailed description of the links.

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-c4-system.md) - C4 system

### Example request

```
grpcurl -cacert ./ops/cert/intermediate_ca.pem -d '{"Id": "http://google.com"}' localhost:50052 metadata_rpc.Metadata/Set
```

### Use Case

**Use cases**:

- [UC-1](./usecases/parsers/README.md) Parse metadata from URL
- [UC-2](./usecases/screenshot/README.md) Made screenshot from URL
