# Link services

> [!NOTE]
> Service for work with link-domain.

**Functionality**:

  * CRUD operations for link
  * SQRS operations for link
  * parse sitemap and save links

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-store-provider.md) - Store Provider Selection

### Architecture

We use the C4 model to describe architecture.

#### System context diagram

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

#### Use case diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/usecase.puml

!include actors/customer.puml

rectangle LinkService {
  usecase (UC-1 Create Link) as UC1
  usecase (UC-2 Read Link) as UC2
  usecase (UC-3 Update Link) as UC3
  usecase (UC-4 Delete Link) as UC4
  usecase (UC-5 SQRS operations for Link) as UC5
  usecase (UC-6 Parse Sitemap and Save Links) as UC6
}

customer --> UC1
customer --> UC2
customer --> UC3
customer --> UC4
customer --> UC5
customer --> UC6
```

**Use cases**:

- UC-1 - Create Link
- UC-2 - Read Link
- UC-3 - Update Link
- UC-4 - Delete Link
- UC-5 - SQRS operations for Link
- UC-6 - Parse Sitemap and Save Links

### Example request

We support reflection for request. You can use [Postman](https://www.postman.com/) or [grpcurl](https://github.com/fullstorydev/grpcurl) for test.

![postman](https://blog.postman.com/wp-content/uploads/2022/01/grpc-author-msg.gif)
