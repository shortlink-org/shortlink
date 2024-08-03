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
System_Ext(temporal, "Temporal", "Orchestration service for managing workflows.")

System_Boundary(oms, "Shop Boundaries") {
    System(oms_service, "OMS Service", "API", "Handles order management API requests.")
    Container(cart_worker, "Cart Worker", "Workflow Worker", "Handles cart workflows.")
    ContainerDb(database, "Database", "SQL Database", "Stores order, customer, and inventory data.")
    Container(cache, "Cache", "In-Memory Data Store", "Caches frequently accessed order data.")
}
Container_Ext(api_gateway, "API Gateway", "API Gateway", "Central entry point for handling order-related requests.")

Rel(customer, api_gateway, "Places orders via", "HTTP/HTTPS")
Rel(api_gateway, mq, "Routes order requests to", "REST API")
Rel(oms_service, database, "Reads/Writes order data", "SQL")
Rel(oms_service, cache, "Utilizes for quick access", "In-Memory")
BiRel_U(oms_service, mq, "Sends/Receives order events", "Messages")
Rel_U(oms_service, temporal, "Uses for workflow orchestration", "Temporal API")
Rel(temporal, cart_worker, "Triggers cart workflows in", "Workflow")

' Use Cases
note right of oms_service
  Use Cases:
  - UC-1: Cart workflows
end note

@enduml
```
