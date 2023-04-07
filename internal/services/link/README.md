# Link services

Service for work with link-domain.

**Functionality**:

  * CRUD operations for link
  * SQRS operations for link
  * parse sitemap and save links

### Architecture

We use C4 model for describe architecture.

#### Context diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml

!include boundaries/gateway.puml
!include boundaries/auth.puml
!include boundaries/link.puml

customer --> gatewayBoundary : uses
gatewayBoundary --> authBoundary : check auth
gatewayBoundary --> linkBoundary : create link
```

### Store provider

<details><summary>Click to expand</summary>

> support - enabled batch mode; filter, etc...  
> scale - scalability/single mode

| Name                            | Support   | Scale    |
|---------------------------------|-----------|----------|
| RAM                             | ✅         | ❌       |
| MongoDB                         | ✅         | ✅       |
| Postgres                        | ✅         | ✅       |
| Redis                           | ❌         | ✅       |
| LevelDB                         | ❌         | ❌       |
| Badger                          | ❌         | ❌       |
| SQLite                          | ❌         | ❌       |
| DGraph                          | ❌         | ✅       |

</details>

### Example request

<details><summary>Click to expand</summary>

We support reflection for request. You can use [Postman](https://www.postman.com/) or [grpcurl](https://github.com/fullstorydev/grpcurl) for test.

![postman](https://blog.postman.com/wp-content/uploads/2022/01/grpc-author-msg.gif)

</details>

### Changelog

<details><summary>Click to expand</summary>

- [19.09.2022] Drop support database: MySQL
- [04.08.2021] Drop support database: scylla, cassandra

</details>
