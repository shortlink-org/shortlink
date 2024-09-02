## Merch [DEPRECATED]

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> This service is responsible for the merch store.

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-use-dapr.md) - Use Dapr

### Architecture

### Container diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml
!include boundaries/billing.puml

customer --> paymentBoundary : Builds a shopping basket
```

### Component diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include services/api-gateway/ext.puml
Container_Ext(mq, "Message queue", "RabbitMQ, Kafka, etc.")

System_Boundary(merchBoundary, "Merch boundary") {
  
    System_Boundary(api_layer, "API Layer") {
        Container(grpc, "gRPC", "Go")
    }
    
    System_Boundary(application_layer, "Application Layer") {
        System_Boundary(command, "Command Layer") {
            Container(add_item_to_basket_command, "Add item to basket", "", "UC-1 Add item to basket")
            Container(remove_item_from_basket_command, "Remove item from basket", "", "UC-2 Remove item from basket")
            Container(update_item_quantity_in_basket_command, "Update item quantity in basket", "", "UC-3 Update item quantity in basket")
        }
        
        System_Boundary(query, "Query Layer") {
            Container(view_items_in_basket_query, "View items in basket", "", "UC-4 View items in basket")
        }
    }
    
    System_Boundary(domain_layer, "Domain Layer") {
        Component(merch, "Merch")
        Component(basket, "Basket")
    }
    
    System_Boundary(infrastructure_layer, "Infrastructure Layer") {
        Component(repository, "Database", "PostgreSQL")
    }
  
}

Rel(api_gateway, merchBoundary, "Adds, deletes items from the basket, decorates it", "gRPC")
Rel(mq, merchBoundary, "New product added/remains of existing product changed", "Async, AMQP, Kafka")
Rel(api_layer, application_layer, "Uses")
Rel(command, domain_layer, "Uses")
Rel(command, infrastructure_layer, "Uses")
Rel(infrastructure_layer, domain_layer, "Uses")

Lay_U(domain_layer, infrastructure_layer)
```

### Use case diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/usecase.puml

!include actors/customer.puml

rectangle "Basket CRUD" {
  usecase (UC-1 Add item to basket) as UC1
  usecase (UC-2 Remove item from basket) as UC2
  usecase (UC-3 Update item quantity in basket) as UC3
  usecase (UC-4 View items in basket) as UC4
  
  url of UC1 is [[./usecases/crud/UC-1.md]]
  url of UC2 is [[./usecases/crud/UC-2.md]]
  url of UC3 is [[./usecases/crud/UC-3.md]]
  url of UC4 is [[./usecases/crud/UC-4.md]]
}

customer --> UC1
customer --> UC2
customer --> UC3
customer --> UC4
```

**Use cases**:

- [UC-1](./usecases/crud_item/README.md) - UC-1 CRUD Item
- [UC-2](./usecases/order-processor/README.md) - UC-2 order processor
- [UC-3](./usecases/use_basket/README.md) - UC-3 Use basket
- [UC-4](./usecases/delivery_order_to_customer/README.md) - UC-4 Delivery order to customer

