# Billing

> [!NOTE]
> Service for work with billing.

**Functionality**:

  * CRUD operations for payment
  * SQRS operations for payment
  * change tariff

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project

### Architecture

We use the C4 model for describe architecture.

#### Context diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml

!include boundaries/gateway.puml
!include boundaries/auth.puml
!include boundaries/payment.puml

customer -right-> gatewayBoundary : uses
gatewayBoundary -down-> authBoundary : check auth
gatewayBoundary -right-> paymentBoundary : create payment
```

#### Use case diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/usecase.puml

!include actors/customer.puml

rectangle BillingService {
    usecase (UC-1 Create Payment) as UC1
    usecase (UC-2 Read Payment) as UC2
    usecase (UC-3 Update Payment) as UC3
    usecase (UC-4 Delete Payment) as UC4
    usecase (UC-5 Change Tariff) as UC5
}

customer --> UC1
customer --> UC2
customer --> UC3
customer --> UC4
customer --> UC5
```

**Use cases**:

- UC-1 - Create Payment
- UC-2 - Read Payment
- UC-3 - Update Payment
- UC-4 - Delete Payment
- UC-5 - Change tariff

### Example request

We support reflection for request. You can use [Postman](https://www.postman.com/) or [grpcurl](https://github.com/fullstorydev/grpcurl) for test.

![postman](https://blog.postman.com/wp-content/uploads/2022/01/grpc-author-msg.gif)
