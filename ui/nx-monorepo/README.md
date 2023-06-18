## nx-monorepo

This repository is a monorepo for the nx workspace.

### ADR

- [ADR-0001](./docs/adr/0001-init.md) - Init. Use Nx
- [ADR-0002](./docs/adr/0002-security.md) - Implementing Security Measures in UI
- [ADR-0003](./docs/adr/0003-transition-to-cloudflare-apps.md) - Transition to Cloudflare Apps

### Services

| Service | Description               | Language/Framework | Docs                                 |
|---------|---------------------------|--------------------|--------------------------------------|
| eslint  | Custom rules for eslint   | JavaScript         | [docs](./packages/eslint/README.md)  |
| landing | Welcome page              | JS/NextJS          | [docs](./packages/landing/README.md) |
| next    | UI service                | JS/NextJS          | [docs](./packages/next/README.md)    |
| ui-kit  | UI kit for ShortLink      | JS/React           | [docs](./packages/ui-kit/README.md)  |

#### Status

This table shows the status of the services deploy to the cluster by ArgoCD.

| Service   | Status                                                                                                                                                  |
|-----------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| next      | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-next&revision=true)](https://shortlink.best/argo/cd/applications/shortlink-next) |
| ui-kit    | [![App Status](https://shortlink.best/argo/cd/api/badge?name=shortlink-next&revision=true)](https://shortlink.best/argo/cd/applications/ui-kit)         |
