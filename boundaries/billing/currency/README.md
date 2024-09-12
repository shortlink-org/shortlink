# Currency

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> This service enables users to select their preferred currency and automatically convert it in real-time. 
> We use subscriptions to Bloomberg and Yahoo to fetch current exchange rates for selected currency pairs. 
> The service also provides historical exchange rate data for analytics and reporting.

### Getting started

We use Makefile for build and deploy.

```bash
$> make help # show help message with all commands and targets
```

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init-project.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-requirements.md) - Requirements for Currency Service
- [ADR-0003](./docs/ADR/decisions/0003-c4-model.md) - C4 Model
- [ADR-0004](./docs/ADR/decisions/0004-api-design.md) - API Design
- [ADR-0005](./docs/ADR/decisions/0005-handling-divergence-in-exchange-rates-between-providers.md) - Handling Divergence in Exchange Rates Between Providers

### Use Cases

- [UC-1](./usecases/load_exchange_rates/README.md) Load Exchange Rates
- [UC-2](./usecases/discrepancies/README.md) Handle Exchange Rate Discrepancies Using Weighted Average Approach
