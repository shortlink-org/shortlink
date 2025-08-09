# Link service

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> Service for work with link-domain.

### Getting started

We use Makefile for build and deploy.

```bash
$> make help # show help message with all commands and targets
```

### ADR

- **Common**:
  - [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
  - [ADR-0003](./docs/ADR/decisions/0003-c4-system.md) - C4 system
- **Infrastructure**:
  - [ADR-0002](./docs/ADR/decisions/0002-store-provider.md) - Store Provider Selection
- **Domain**:
  - [ADR-0004](./docs/ADR/decisions/0004-domain-link.md) - **domain:** _Link_

### Use Cases

- [UC-1](./internal/usecases/link/README.md) Manage link
- [UC-2](./internal/usecases/link_sqrs/README.md) SQRS operations for link
- [UC-3](./internal/usecases/sitemap/README.md) Parse sitemap and save links

### Docs

- [Contributing](./docs/CONTRIBUTING.md)
- [ENV configuration](./docs/env.md)
