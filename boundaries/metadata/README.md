# Metadata

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> This service to the ShortLink system, designed to enrich link information.
> Upon the creation of a new link, this service fetches metadata and generates screenshots,
> subsequently storing this data to provide a more detailed description of the links.

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-c4-system.md) - C4 system

### Use Cases

- [UC-1](./internal/usecases/parsers/README.md) Manage metadata for URL
- [UC-2](./internal/usecases/screenshot/README.md) Made screenshot from URL

### Docs

- [Contributing](./docs/CONTRIBUTING.md)
- [ENV configuration](./docs/env.md)
