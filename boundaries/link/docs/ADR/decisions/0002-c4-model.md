# 2. C4 Model for Link boundary context

Date: 2024-01-01

## Status

Accepted

## Context

With the Link Boundary Context encompassing multiple services (Link, Metadata, and Proxy), there is a pressing need 
for a comprehensive visualization of this intricate architecture. The C4 model, renowned for its effectiveness in mapping 
software architecture, is apt for this purpose, ensuring clarity and structure.

## Decision

We will apply the C4 model to detail the architecture of the Link Boundary Context. This includes 
creating System Context, Container, and Component diagrams, and optionally, Class diagrams, 
for each service within the boundary.

## Consequences

+ Improved understanding and communication of system architecture.
+ Increased efficiency in development and maintenance.

### C4

#### Level 1: System Context diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

LAYOUT_WITH_LEGEND()

title System Context diagram for Link Boundary Context

Person(user, "User", "A user of the ShortLink system.")
SystemQueue_Ext(mq, "Message Queue", "Handles event-driven operations for the system.")
System_Boundary(slb, "Link Boundary Context") {
    System(link_service, "Link Service", "Service for managing short links.")
    System(metadata_service, "Metadata Service", "Service for link metadata and screenshots.")
    System(proxy_service, "Proxy Service", "Service for redirecting short links to original URLs.")
}

Rel(user, link_service, "Uses")
Rel(user, metadata_service, "Queries")
Rel(user, proxy_service, "Redirects through")
Rel_U(mq, link_service, "Sends events to")
Rel_U(mq, metadata_service, "Sends events to")

@enduml
```

#### Level 2: Container diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title Container diagram for Link Boundary Context

Person(user, "User", "A user of the ShortLink system.")
Container_Ext(api_gateway, "API Gateway", "API Gateway", "Central entry point for handling requests.")
SystemQueue_Ext(mq, "Message Queue", "Handles event-driven operations.")
Container_Ext(s3_minio, "S3 (MinIO)", "Storage", "Stores screenshots and metadata.")
System_Boundary(slb, "Link Boundary Context") {
    Container(link_service, "Link Service", "Service", "Manages CRUD operations for short links.")
    Container(proxy_service, "Proxy Service", "Service", "Redirects short links to original URLs.")
    Container(metadata_service, "Metadata Service", "Service", "Generates and stores metadata and screenshots for links.")
    ContainerDb(metadata_db, "Metadata Database", "Database", "Stores metadata and screenshot data.")
    ContainerDb(link_db, "Link Database", "Database", "Stores short link data.")
    Container(link_cache, "Link Cache Server", "Cache", "Caches frequently accessed link data.")
    Container(metadata_cache, "Metadata Cache Server", "Cache", "Caches frequently accessed metadata.")
}

Rel_D(mq, link_service, "Exchanges events")
Rel_D(mq, metadata_service, "Exchanges events")
Rel(user, api_gateway, "Makes requests to")
Rel(api_gateway, link_service, "Routes to")
Rel(api_gateway, metadata_service, "Routes to")
Rel(api_gateway, proxy_service, "Routes to")
Rel(link_service, link_db, "Reads/Writes data")
Rel(metadata_service, metadata_db, "Reads/Writes data")
Rel(link_service, link_cache, "Utilizes")
Rel(metadata_service, metadata_cache, "Utilizes")
Rel(metadata_service, s3_minio, "Stores data in")

@enduml
```
