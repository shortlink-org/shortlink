# 2. C4 system

Date: 2024-05-05

## Status

Accepted

## Context

We need to understand the internal architecture of individual containers in our system.

## Decision

Adopt a Level 3 C4 Component Diagram. This diagram will focus on the components within each container,
showing how they interact internally and with external elements.

## Consequences

Provides an in-depth view of the internal architecture of containers, aiding in development and troubleshooting.

### System context diagram

```plantuml
@startuml C4_Context_Diagram_for_OMS
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title System Context Diagram for Order Management System

Person(customer, "Customer", "A user placing orders through the online platform.")
SystemQueue_Ext(mq, "Message Queue", "Manages event-driven operations.")
System_Boundary(oms, "Shop Boundaries") {
    System(order_service, "OMS", "Processes and manages customer orders.")
    ContainerDb(database, "Database", "Database", "Stores order, customer, and inventory data.")
    Container(cache, "Cache", "Cache", "Caches frequently accessed order data.")
}
Container_Ext(api_gateway, "API Gateway", "API Gateway", "Central entry point for handling order-related requests.")

Rel(customer, api_gateway, "Places orders via")
Rel(api_gateway, order_service, "Routes order requests to")
Rel(order_service, database, "Reads/Writes order data")
Rel(order_service, cache, "Utilizes for quick access")
Rel_U(order_service, mq, "Sends/Receives order events")

' Use Cases
note right of order_service
  Use Cases:
  - UC-1: Create and track orders
  - UC-2: Refund orders

@enduml
```
