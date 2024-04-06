## Shop admin

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> The Shop admin in this software system is primarily concerned with the management of goods and services.

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-auth-mode.md) - Auth mode

### Architecture

We use the C4 model to describe architecture.

#### Container diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

title Container Diagram

actor User
database "Database" as DB
[Admin Service] as Admin
[Goods and Prices Service] as GPS

User --> Admin : Manage goods and services
Admin --> GPS : CRUD operations
GPS --> DB : Persist data

@enduml
```

#### Use case diagram

The use case diagram shows which functionality of the developed software system is
available to each group of users.

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/usecase.puml

actor User
actor Admin

User --> (Browse Goods)
User --> (Check Prices)
Admin --> (Manage Goods)
Admin --> (Set Prices)
```

**Use cases**:

- [UC-1](src/usecases/manage_goods/README.md) - Manage goods


### Contributing

- [CONTRIBUTING.md](./CONTRIBUTING.md)
