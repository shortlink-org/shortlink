# 3. C4 system

Date: 2024-01-03

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
@startuml C4_Context_Diagram_for_Link_Service
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title System Context Diagram for Link Service

Person(customer, "Customer", "A user interacting with the system.")
System(auth_service, "Authentication Service", "Handles user authentication.")
SystemQueue_Ext(mq, "Message Queue", "Manages event-driven operations.")
System_Boundary(lb, "Link Boundaries") {
    System(link_service, "Link Service", "Manages and processes short links.")
    ContainerDb(database, "Database", "Database", "Stores link data.")
    Container(cache, "Cache", "Cache", "Caches frequently accessed data.")
}
Container(api_gateway, "API Gateway", "API Gateway", "Central entry point for handling requests.")

Rel(customer, api_gateway, "Uses")
Rel(api_gateway, auth_service, "Authenticates via")
Rel_R(auth_service, link_service, "Validates users for")
Rel(api_gateway, link_service, "Routes requests to")
Rel(link_service, database, "Reads/Writes data")
Rel(link_service, cache, "Utilizes for performance")
Rel(link_service, mq, "Sends/Receives events")

' Use Cases
note right of link_service
  Use Cases:
  1. Manage link
  2. SQRS operations for link
  3. Parse sitemap and save links
end note

@enduml
```
