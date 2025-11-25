# 3. C4 System

Date: 2025-11-25

## Status

Accepted

## Context

The decision to implement the C4 model (Context, Containers, Components, Code) in our BFF service architecture is driven
by the need for a comprehensive, structured, and scalable approach to documenting and communicating
the software architecture. The C4 model is renowned for its effectiveness in providing various levels of abstraction,
which can be understood by different stakeholders ranging from developers to business personnel.

The BFF service is built using **Go (Golang)** and leverages several key technologies:

- **chi router** for HTTP routing and middleware
- **oapi-codegen** for OpenAPI 3.0 code generation
- **Redis** for caching to improve performance
- **Ory Kratos** for authentication
- **SpiceDB** for authorization and permission checks
- **OpenTelemetry** for distributed tracing and observability
- **gRPC** for communication with backend services

Implementing the C4 model will enable us to:

- Clearly define the high-level context of the BFF service, including key system interactions and integrations.
- Break down the service into containers (HTTP API, controllers, etc.), illustrating the overall system architecture.
- Detail out components within each container, clarifying internal architectures.
- Optionally, delve into the code level to understand how components are implemented.

This approach is expected to streamline our documentation process,
improve communication about system architecture within the team, and facilitate onboarding of new team members.

## Decision

We will focus specifically on creating **System Context**, **Container**, and **Component Diagrams** as part of the C4 model implementation.
These diagrams will detail the internal structure and interactions of components within the BFF service,
providing a clear view of how various parts of the system work together.

### System context diagram

```plantuml
@startuml C4_Context_Diagram_for_BFF_Service
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

LAYOUT_WITH_LEGEND()

title System Context Diagram for BFF Service

Person(user, "User", "An end user who interacts with the web application to manage short links.")
System_Ext(web_app, "Web Application", "Next.js/React", "Frontend web application that provides UI for managing short links.")
System_Ext(kratos, "Ory Kratos", "Identity Provider", "Handles user authentication, registration, and identity management.")
System_Ext(spicedb, "SpiceDB", "Authorization Service", "Handles permission checks using Zanzibar-style fine-grained authorization.")
System_Boundary(link_boundary, "Link Boundary Context") {
    System(link_service, "Link Service", "Domain Service", "Core service for managing short links and link operations.")
    System(metadata_service, "Metadata Service", "Domain Service", "Service for fetching and storing link metadata and screenshots.")
}
System_Boundary(bff_boundary, "BFF Boundary") {
    System(bff_service, "BFF Service", "API Gateway", "Backend-for-Frontend service that aggregates and transforms data for the web application.")
}

Rel(user, web_app, "Uses", "HTTPS")
Rel(web_app, bff_service, "Makes API requests", "HTTPS/REST")
Rel(bff_service, kratos, "Validates user sessions", "HTTPS")
Rel(bff_service, spicedb, "Checks user permissions", "gRPC")
Rel(bff_service, link_service, "Proxies link operations", "gRPC")
Rel(bff_service, metadata_service, "Queries metadata", "gRPC")

' Interaction notes
note right of kratos
  **Ory Kratos**
  Only user session validation
  - Validates authentication tokens
  - Checks if user is logged in
  - Provides user identity information
end note

note right of spicedb
  **SpiceDB**
  Object-level permissions
  - "can read link"
  - "can modify link"
  - Fine-grained authorization
  - Zanzibar-style permissions
end note

' Use Cases
note right of bff_service
  Use Cases:
  1. Manage links (CRUD operations)
  2. CQRS operations for links
  3. Sitemap operations
end note

@enduml
```

### Container diagram

```plantuml
@startuml Container_Diagram_for_BFF_Service
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title Container Diagram for BFF Service

Person(web_client, "Web Client", "A web application user interacting with the system.")
System_Ext(kratos, "Ory Kratos", "Handles user authentication and identity management.")
System_Ext(spicedb, "SpiceDB", "Handles permission checks using Zanzibar-style authorization.")
System_Ext(link_service, "Link Service", "Service for managing short links.")
System_Ext(link_command, "Link Command Service", "CQRS command operations.")
System_Ext(link_query, "Link Query Service", "CQRS query operations.")
System_Ext(sitemap_service, "Sitemap Service", "Handles sitemap operations.")

System_Boundary(bff_system, "BFF System") {
    Container(bff_api, "BFF API", "Go/HTTP", "HTTP REST API server using chi router.")
    ContainerDb(cache, "Redis Cache", "Redis", "Caches frequently accessed data to improve performance.")
}

Rel(web_client, bff_api, "Makes requests", "HTTPS/REST")
Rel(bff_api, kratos, "Authenticates via", "HTTPS")
Rel(bff_api, spicedb, "Checks permissions via", "gRPC")
Rel(bff_api, cache, "Reads/Writes cached data", "Redis Protocol")
Rel(bff_api, link_service, "Proxies requests to", "gRPC")
Rel(bff_api, link_command, "Sends commands to", "gRPC")
Rel(bff_api, link_query, "Sends queries to", "gRPC")
Rel(bff_api, sitemap_service, "Proxies requests to", "gRPC")

@enduml
```

### Component diagram

```plantuml
@startuml Component_Diagram_for_BFF_Service
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title Component Diagram for BFF Service

Person(web_client, "Web Client", "A web application user.")
System_Ext(kratos, "Ory Kratos", "Handles user authentication and identity management.")
System_Ext(spicedb, "SpiceDB", "Handles permission checks using Zanzibar-style authorization.")
System_Ext(link_service, "Link Service", "Manages short links.")
System_Ext(link_command, "Link Command Service", "CQRS command operations.")
System_Ext(link_query, "Link Query Service", "CQRS query operations.")
System_Ext(sitemap_service, "Sitemap Service", "Handles sitemap operations.")
ContainerDb_Ext(cache, "Redis Cache", "Redis", "Caches frequently accessed data.")

System_Boundary(bff_system, "BFF System") {
    Container(bff_api, "BFF API", "Go/HTTP", "HTTP REST API server using chi router.")

    Container_Boundary(interface_layer, "Interface Layer") {
        Container_Boundary(rest_handlers, "REST Handlers") {
            Component(link_controller, "Link Controller", "Go", "Handles link CRUD operations and data transformation.")
            Component(cqrs_controller, "CQRS Controller", "Go", "Handles CQRS operations for links with data aggregation.")
            Component(sitemap_controller, "Sitemap Controller", "Go", "Handles sitemap operations.")
        }
        Container_Boundary(middleware, "Middleware Stack") {
            Component(tracing_middleware, "Tracing Middleware", "Go", "Handles distributed tracing (1st in pipeline).")
            Component(logger_middleware, "Logger Middleware", "Go", "Logs HTTP requests and responses (2nd in pipeline).")
            Component(metrics_middleware, "Metrics Middleware", "Go", "Collects metrics and observability data (3rd in pipeline).")
            Component(csrf_middleware, "CSRF Middleware", "Go", "Protects against CSRF attacks (4th in pipeline).")
            Component(auth_middleware, "Auth Middleware", "Go", "Validates authentication tokens via Ory Kratos (5th in pipeline).")
            Component(permission_middleware, "Permission Middleware", "Go", "Checks permissions via SpiceDB (6th in pipeline).")
        }
    }

    Container_Boundary(application_layer, "Application Layer") {
        Component(aggregation_service, "Aggregation Service", "Go", "Aggregates data from multiple backend services.")
        Component(adapter_service, "Adapter Service", "Go", "Transforms and adapts data between frontend and backend formats.")
    }

    Container_Boundary(infrastructure_layer, "Infrastructure Layer") {
        Component(cache_component, "Redis Client", "Go", "Manages Redis cache operations for performance optimization.")
        Component(grpc_link_client, "gRPC Link Client", "Go", "Client for Link Service gRPC communication.")
        Component(grpc_command_client, "gRPC Command Client", "Go", "Client for Link Command Service gRPC communication.")
        Component(grpc_query_client, "gRPC Query Client", "Go", "Client for Link Query Service gRPC communication.")
        Component(grpc_sitemap_client, "gRPC Sitemap Client", "Go", "Client for Sitemap Service gRPC communication.")
        Component(error_transformer, "Error Transformer", "Go", "Transforms domain and gRPC errors into structured JSON responses.")
    }

    Container_Boundary(cross_cutting, "Cross-Cutting Concerns") {
        Container_Boundary(observability, "Observability") {
            Component(logger_middleware, "Logger Middleware", "Go", "Logs HTTP requests and responses.")
            Component(metrics_middleware, "Metrics Middleware", "Go", "Collects metrics and observability data.")
            Component(tracing_middleware, "Tracing Middleware", "Go", "Handles distributed tracing.")
            Component(profiling_component, "Profiling Component", "Go", "Provides pprof endpoints for performance profiling.")
            Component(flight_trace_component, "Flight Trace", "Go", "Records request traces for debugging and analysis.")
        }
        Component(i18n_component, "i18n Component", "Go", "Provides internationalization support (en-GB, de-DE, fr-CH).")
    }
}

Rel(web_client, bff_api, "Makes requests", "HTTPS/REST")
Rel(bff_api, tracing_middleware, "Processes through", "1. Request pipeline")
Rel(tracing_middleware, logger_middleware, "Processes through", "2. Request pipeline")
Rel(logger_middleware, metrics_middleware, "Processes through", "3. Request pipeline")
Rel(metrics_middleware, csrf_middleware, "Processes through", "4. Request pipeline")
Rel(csrf_middleware, auth_middleware, "Processes through", "5. Request pipeline")
Rel(auth_middleware, permission_middleware, "Processes through", "6. Request pipeline")
Rel(permission_middleware, link_controller, "Routes to", "After middleware")
Rel(permission_middleware, cqrs_controller, "Routes to", "After middleware")
Rel(permission_middleware, sitemap_controller, "Routes to", "After middleware")
Rel(auth_middleware, kratos, "Validates tokens", "HTTPS")
Rel(permission_middleware, spicedb, "Checks permissions", "gRPC")

' Middleware order note
note right of middleware
  **Middleware Pipeline Order:**
  1. Tracing (captures full request context)
  2. Logger (includes trace-id in logs)
  3. Metrics (includes user-id in metrics)
  4. CSRF (security validation)
  5. Auth (user session validation)
  6. Permission (object-level checks)

  This order ensures trace-id and user-id
  are available in all observability data.
end note

' Security interaction notes
note right of auth_middleware
  **Ory Kratos Interaction:**
  Only user session validation
  - Validates authentication tokens
  - Checks if user is logged in
  - Provides user identity information
end note

note right of permission_middleware
  **SpiceDB Interaction:**
  Object-level permissions
  - "can read link"
  - "can modify link"
  - Fine-grained authorization
  - Zanzibar-style permissions
end note
Rel(link_controller, aggregation_service, "Uses for data aggregation")
Rel(cqrs_controller, aggregation_service, "Uses for data aggregation")
Rel(sitemap_controller, adapter_service, "Uses for data transformation")
Rel(aggregation_service, cache_component, "Uses for caching", "Redis Protocol")
Rel(adapter_service, cache_component, "Uses for caching", "Redis Protocol")
Rel(cache_component, cache, "Reads/Writes", "Redis Protocol")
Rel(aggregation_service, grpc_link_client, "Uses", "gRPC")
Rel(aggregation_service, grpc_command_client, "Uses", "gRPC")
Rel(aggregation_service, grpc_query_client, "Uses", "gRPC")
Rel(adapter_service, grpc_sitemap_client, "Uses", "gRPC")
Rel(grpc_link_client, link_service, "Calls", "gRPC")
Rel(grpc_command_client, link_command, "Calls", "gRPC")
Rel(grpc_query_client, link_query, "Calls", "gRPC")
Rel(grpc_sitemap_client, sitemap_service, "Calls", "gRPC")
Rel(link_controller, error_transformer, "Uses for error transformation")
Rel(cqrs_controller, error_transformer, "Uses for error transformation")
Rel(sitemap_controller, error_transformer, "Uses for error transformation")
Rel(bff_api, i18n_component, "Uses for localization")
Rel(bff_api, profiling_component, "Exposes via", "/debug/pprof")
Rel(bff_api, flight_trace_component, "Uses for debugging")

@enduml
```

## Consequences

### Benefits

- **Improved understanding and communication**: Clear visualization of BFF service architecture helps all stakeholders understand the system design.
- **Increased efficiency**: Well-documented architecture reduces time spent on understanding system interactions during development and maintenance.
- **Better integration clarity**: Clear visualization of how the BFF service integrates with Ory Kratos, SpiceDB, and backend services.
- **Enhanced onboarding**: New team members can quickly understand the system structure through visual diagrams.
- **Facilitated discussions**: Architecture diagrams serve as a common reference point for technical discussions and decision-making.

### Technical Details

The diagrams illustrate:

- **Request flow**: From web client through middleware stack to controllers and backend services.
- **Security layers**: Authentication (Ory Kratos) and authorization (SpiceDB) integration points.
- **Performance optimization**: Redis caching strategy for frequently accessed data.
- **Observability**: Metrics, tracing, profiling, and flight trace components for monitoring and debugging.
- **Error handling**: Structured error transformation from domain/gRPC errors to client-friendly JSON responses.
- **Internationalization**: Multi-language support (en-GB, de-DE, fr-CH) for localized responses.
