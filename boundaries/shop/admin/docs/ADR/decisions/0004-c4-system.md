# 4. C4 system

Date: 2024-08-13

## Status

Accepted

## Context

We need a standardized way to describe the architecture of our system. 
The C4 model provides a clear and structured approach to visualize the architecture at different levels of detail. 
This model helps in understanding the system's structure, its components, and their interactions, which is crucial 
for both development and maintenance.

## Decision

We have decided to adopt the C4 model to describe our system architecture. 
We will use PlantUML to create these diagrams and include them in our documentation.

## Consequences

By adopting the C4 model, we ensure a consistent and clear way to document our system architecture.
This approach will make it easier for new developers to understand the system, facilitate better communication among team members, 
and provide a solid foundation for future architectural decisions. 

### Architecture

We use the C4 model to describe architecture.

#### Container diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

title Container Diagram

Person(user, "User")
ContainerDb(database, "Database", "PostgreSQL", "Stores goods and prices data")
Container(adminService, "Admin Service", "Django", "Manages goods and services")
Container(goodsService, "Goods and Prices Service", "Django", "Provides goods and prices")

Rel(user, adminService, "Manage goods and services")
Rel(adminService, goodsService, "CRUD operations")
Rel(goodsService, database, "Persist data")
@enduml
```

#### Use case diagram

The use case diagram shows which functionality of the developed software system is
available to each group of users.

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

title Use Case Diagram

actor User as user
actor Admin as admin

System_Boundary(shopAdminSystem, "Shop Admin System") {
    usecase "Browse Goods" as browseGoods
    usecase "Check Prices" as checkPrices
    usecase "Manage Goods" as manageGoods
    usecase "Set Prices" as setPrices
}

user --> browseGoods : "Uses"
user --> checkPrices : "Uses"
admin --> manageGoods : "Uses"
admin --> setPrices : "Uses"
@enduml
```
