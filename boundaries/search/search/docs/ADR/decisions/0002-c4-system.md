# 2. C4 System

Date: 2024-02-17

## Status

Accepted

## Context

To enhance our software system's search capabilities, it's crucial to have a clear understanding 
of the search service's internal architecture. This involves dissecting the components within the search service container, 
their roles, and how they interact both internally and externally.

## Decision

We have decided to adopt a Level 3 C4 Component Diagram for the search service. This diagram will detail the components 
within the search service container, illustrating their internal interactions and how they connect with external 
systems and services. The focus will be on delineating the components responsible for processing search queries, 
handling indexing, and integrating with database and other system domains for real-time data synchronization.

## Consequences

Employing a Level 3 C4 Component Diagram offers a comprehensive view into the search service's internal architecture. 
This depth of insight is invaluable for development, maintenance, and troubleshooting, ensuring that the team 
can effectively understand and navigate the complexities of the search functionality.

### System Context Diagram

```plantuml
@startuml C4_Context_Diagram_for_Search_Service
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title System Context Diagram for Search Service

Person(user, "User", "A user searching for information.")
System_Boundary(ss, "Search Service") {
    Container(search_api, "Search API", "gRPC API", "Handles incoming search requests.")
    ContainerDb(search_db, "Search Database", "Elasticsearch", "Indexes and stores documents for search.")
    Container(search_indexer, "Search Indexer", "Component", "Processes and indexes new or updated documents.")
    Container(query_processor, "Query Processor", "Component", "Parses and executes search queries against Elasticsearch.")
}
System_Ext(other_domains, "Other System Domains", "Integrates for real-time data updates.")

Rel(user, search_api, "Submits search queries via")
Rel(search_api, query_processor, "Forwards queries to")
Rel(query_processor, search_db, "Searches in")
Rel(search_indexer, search_db, "Indexes documents in")
Rel(search_api, other_domains, "Receives data updates from")
Rel(other_domains, search_api, "Sends updates to")

' Use Cases
note right of search_api
  Use Cases:
  1. Full-text search queries
  2. Real-time indexing of documents
  3. Advanced search features like filtering and sorting
end note

@enduml
```
