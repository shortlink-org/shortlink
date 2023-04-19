## Referral Service

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project

**Functionality**:

  * CRUD operations for referral

### Architecture

We use C4 model for describe architecture.

#### Container diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml

!include boundaries/marketing.puml

customer --> marketingBoundary : uses
```

#### Use case diagram

The use case diagram shows which functionality of the developed software system is
available to each group of users.

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/usecase.puml

!include actors/customer.puml
!include actors/manager.puml

rectangle Referral {
  usecase (UC-1 CRUD referral) as UC1
  usecase (UC-2 Use referral) as UC2
}

customer --> UC2
manager --> UC1
```

**Use cases**:

- [UC-1](./application/crud_referral/README.md) - CRUD referral
- [UC-2](./application/use_referral/README.md) - Use referral

### Getting Started

```
$> python3 main.py
```
