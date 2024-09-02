# 2. C4 system

Date: 2024-09-02

## Status

Accepted

## Context

To better understand and communicate the interactions between our system and external entities, 
we need to create a high-level overview that highlights these relationships. 
The C4 System Context Diagram provides a clear and concise way to represent the system's architecture at this level.

## Decision

We will create a System Context Diagram using the C4 model combined with PlantUML. 
This approach allows us to visualize the key components of our system, including the users and external systems 
it interacts with, in a manner that is easy to maintain and understand.

## Consequences

This decision will help provide a clear understanding of how our system fits into the broader context, 
which will be valuable for both development and communication with stakeholders. 
Using PlantUML for this diagram ensures that the diagram is easily updateable and version-controlled.

### System context diagram

```plantuml
@startuml C4_Context_Diagram_for_API_Bridge
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

LAYOUT_WITH_LEGEND()

title System Context Diagram for OMS GraphQL API Bridge

Person(user, "External User", "A user accessing the public API.")
System_Ext(grpc_api, "gRPC API", "Internal gRPC service for OMS operations.")
System(graphql_api_bridge, "GraphQL API Bridge", "Translates gRPC API to GraphQL for public access.")

Rel(user, graphql_api_bridge, "Accesses API via", "GraphQL")
Rel(graphql_api_bridge, grpc_api, "Translates and forwards requests to", "gRPC")

@enduml
```
