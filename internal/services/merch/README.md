## Merch

This service is responsible for the merch store.

### Architecture

### Container diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml
!include boundaries/payment.puml

customer --> paymentBoundary : Builds a shopping basket
```

### Use case diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/usecase.puml

!include actors/customer.puml

rectangle Basket {
  usecase (UC-1 Add item to basket) as UC1
  usecase (UC-2 Remove item from basket) as UC2
  usecase (UC-3 Update item quantity in basket) as UC3
  usecase (UC-4 View items in basket) as UC4
  
  url of UC1 is [[./usecases/UC-1.md]]
  url of UC2 is [[./usecases/UC-2.md]]
  url of UC3 is [[./usecases/UC-3.md]]
  url of UC4 is [[./usecases/UC-4.md]]
}

customer --> UC1
customer --> UC2
customer --> UC3
customer --> UC4
```

**Use cases**:

- [UC-1](./usecases/UC-1.md) - Add item to basket
- [UC-2](./usecases/UC-2.md) - Remove item from basket
- [UC-3](./usecases/UC-3.md) - Update item quantity in basket
- [UC-4](./usecases/UC-4.md) - View items in basket
