# Order management system (OMS) service

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> Service for work with carts, orders.

### Getting started

We use Makefile for build and deploy.

```bash
$> make help # show help message with all commands and targets
```

### ADR

- **Common**:
  - [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
  - [ADR-0002](./docs/ADR/decisions/0002-c4-system.md) - C4 system

### Use Cases

- [UC-1](internal/usecases/cart/README.md) Cart workflows
- [UC-2](internal/usecases/order/README.md) Order workflows
