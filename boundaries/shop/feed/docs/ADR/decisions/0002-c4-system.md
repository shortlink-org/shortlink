# 2. C4 system

Date: 2024-10-05

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
@startuml C4_Context_Diagram_for_Feed_Service
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title System Context Diagram for Feed Service

Person(partner, "Partner", "A partner receiving generated feeds.")
System_Ext(minio, "Minio", "S3-like block storage for storing generated feeds.")

System_Boundary(feed_service_boundary, "Feed Service Boundaries") {
    Container(feed_service, "Feed Service", "Service", "Generates XML feeds for partners based on their preferences.")
    Container(cron_job, "Feed CronJob", "Kubernetes CronJob", "Triggers feed generation every 24 hours.")
    Container(graphql_service, "GraphQL API", "Service", "Fetches data needed to generate feeds.")
}

Rel(cron_job, minio, "Stores generated feeds in", "S3-like block storage")
Rel(partner, feed_service, "Receives generated feeds via", "XML")
Rel(feed_service, graphql_service, "Fetches necessary data from", "GraphQL API")
Rel_Right(feed_service, minio, "Stores generated feeds in", "S3-like block storage")
Rel(cron_job, feed_service, "Triggers feed generation every 24 hours")

@enduml
```
