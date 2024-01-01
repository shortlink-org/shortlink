# 2. C4 System

Date: 2024-01-01

## Status

Accepted

## Context

The decision to implement the C4 model (Context, Containers, Components, Code) in our system architecture is driven 
by the need for a comprehensive, structured, and scalable approach to documenting and communicating 
the software architecture. The C4 model is renowned for its effectiveness in providing various levels of abstraction, 
which can be understood by different stakeholders ranging from developers to business personnel.

Implementing the C4 model will enable us to:

- Clearly define the high-level context of the system, including key system interactions and integrations.
- Break down the system into containers (applications, data stores, etc.), illustrating the overall system architecture.
- Detail out components within each container, clarifying internal architectures.
- Optionally, delve into the code level to understand how components are implemented.

This approach is expected to streamline our documentation process, 
improve communication about system architecture within the team, and facilitate onboarding of new team members.

## Decision

We will focus specifically on creating **Component Diagrams** as part of the C4 model implementation. 
These diagrams will detail the internal structure and interactions of components within each system container, 
providing a clear view of how various parts of the system work together.

### System context diagram

```plantuml
@startuml Component Diagram

!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

SystemQueue_Ext(mq, "Message Queue", "Receives CRUD events for links.")
Component_Ext(api, "API", "HTTP/REST", "API to retrieve metadata by URL")
Container_Ext(external_sites, "External Websites", "Internet", "External sites for fetching metadata")
Container_Ext(s3_store, "S3 Store (Minio)", "Storage", "Stores screenshots of URLs.")

System_Boundary(metadata_system, "Metadata System") {
    Container(metadata_service, "Metadata Service", "Go", "Service for processing metadata of links.")
    ContainerDb(database, "Database", "Relational DBMS", "Stores metadata of links.")

    Component(create_metadata, "Create Metadata", "Use Case", "Generates metadata for links.")
    Component(create_screenshot, "Create Screenshot", "Use Case", "Generates screenshots for links.")
}

Rel(mq, metadata_service, "Triggers on CRUD link events", "Message/Event")
Rel(metadata_service, database, "Stores metadata")
Rel(metadata_service, s3_store, "Stores screenshots")
Rel(api, metadata_service, "Fetches metadata")
Rel(metadata_service, external_sites, "Fetches metadata")
Rel(metadata_service, create_metadata, "Executes use case")
Rel(metadata_service, create_screenshot, "Executes use case")

@enduml
```
