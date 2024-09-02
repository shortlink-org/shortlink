## Shop Boundary

> [!NOTE]
> The Shop Boundary in this software system is primarily concerned with the management of goods and services.

| Service     | Description        | Language/Framework   | Docs                            | Status                                                                                                                                                            |
|-------------|--------------------|----------------------|---------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| admin       | Shop admin         | Python (Django)      | [docs](./admin/README.md)       | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-admin&revision=true)](https://argo.shortlink.best/applications/shortlink-admin)               |
| bff         | API Gateway        | NodeJS (Wundergraph) | [docs](./gateway/README.md)     | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-shop-gateway&revision=true)](https://argo.shortlink.best/applications/shortlink-shop-gateway) |
| oms         | Order management   | Temporal             | [docs](./oms/README.md)         | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-oms&revision=true)](https://argo.shortlink.best/applications/shortlink-oms)                   |
| oms-graphql | GraphQL API Bridge | Coming soon          | [docs](./oms-graphql/README.md) | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-oms-graphql&revision=true)](https://argo.shortlink.best/applications/shortlink-oms-graphql)   |
| ui          | Shop service       | JS/NextJS            | [docs](./ui/README.md)          | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-shop-ui&revision=true)](https://argo.shortlink.best/applications/shortlink-shop-ui)           |

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the Shop Boundary
- [README.md](./docs/ADR/README.md) - Architecture Decision Records
