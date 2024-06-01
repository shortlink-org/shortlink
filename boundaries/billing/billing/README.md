# Billing

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> Service for work with billing.

### Getting started

We use Makefile for build and deploy.

```bash
$> make help # show help message with all commands and targets
```

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-use-decimal-for-financal-types.md) - Use Decimal for Financial Types

### Use Cases

- [UC-1](./internal/usecases/account/README.md) Works with an account
- [UC-2](./internal/usecases/order/README.md) Works with an order
- [UC-3](./internal/usecases/payment/README.md) Works with a payment
- [UC-4](./internal/usecases/tariff/README.md) Works with a tariff

### Docs

- [ENV configuration](./docs/env.md)
- [Example request](./docs/example-request.md)
