## Payment Boundary

> [!NOTE]
> The Payment Boundary in a software system deals with the financial transactions aspect of the application. 
> It encompasses the integration and management of payment gateways, processors, and e-wallets to facilitate online 
> transactions. This boundary includes handling customer checkouts, invoicing, and ensuring secure transaction processes. 
> It also involves implementing fraud detection mechanisms to safeguard against unauthorized transactions and managing 
> recurring payments for subscriptions. Effective management within this boundary is crucial for ensuring a smooth, 
> secure, and user-friendly payment experience for customers.

| Service      | Description          | Language/Framework | Docs                                                  | Status                                                                                                                                                            |
|--------------|----------------------|--------------------|-------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| billing      | Billing service      | Go                 | [docs](./internal/boundaries/payment/billing/README.md)         | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-billing&revision=true)](https://argo.shortlink.best/applications/shortlink-billing)           |
| wallet       | Wallet service       | Go (Solidity)      | [docs](./internal/boundaries/payment/wallet/README.md)          |                                                                                                                                                                   |

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the Payment Boundary
