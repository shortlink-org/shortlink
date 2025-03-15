### Boundaries

> [!TIP]
>
> Our project follows Domain-Driven Design (DDD) principles, organizing code into distinct domains for clarity and easier updates.

| Bounded Context                  | Description              | Type subdomain | Docs                                                  |
|----------------------------------|--------------------------|----------------|-------------------------------------------------------|
| API Gateway                      | Gateway for all services | Supporting     | [docs](./boundaries/api/README.md)                    |
| Auth Boundary                    | Auth services            | Generic        | [docs](https://github.com/shortlink-org/auth)         |
| Billing Boundary [External]      | Payment services         | Generic        | [docs](https://github.com/shortlink-org/billing)      |
| Chat Boundary [External]         | Chat services            | Supporting     | [docs](https://github.com/shortlink-org/chat)         |
| Delivery Boundary [External]     | Delivery services        | Supporting     | [docs](https://github.com/shortlink-org/delivery)     |
| DS Boundary                      | Data Science services    | Supporting     | [docs](./boundaries/ds/README.md)                     |
| Link Boundary [External]         | Link services            | Core           | [docs](https://github.com/shortlink-org/link)         |
| Marketing Boundary [External]    | Marketing services       | Supporting     | [docs](https://github.com/shortlink-org/marketing)    |
| Notification Boundary [External] | Notification services    | Generic        | [docs](https://github.com/shortlink-org/notification) |
| Platform Boundary                | Platform services        | Supporting     | [docs](./boundaries/platform/README.md)               |
| Search Boundary [External]       | Search services          | Supporting     | [docs](https://github.com/shortlink-org/search)       |
| Shop Boundary [External]         | Shop services            | Supporting     | [docs](https://github.com/shortlink-org/shop)         |
| ShortDB Boundary [External]      | ShortDB services         | Supporting     | [docs](https://github.com/shortlink-org/shortdb)      |

#### Draft context
 
- [accounting](./draft/accounting)
