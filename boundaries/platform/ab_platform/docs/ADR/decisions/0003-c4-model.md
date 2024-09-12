# 3. C4 Model

Date: 2024-09-13

## Status

Accepted

## Context

The A/B Platform will be built using a monolithic architecture to support experimentation, including test case management, 
user assignment, and metrics collection. To clearly visualize the architecture, we will apply the C4 model,
which focuses on describing the system's components at different levels (System Context, Container, and Component diagrams).

The first focus is on visualizing the CRUD module for test cases within the platform, as it forms the foundation of the system.

## Decision

We will apply the C4 model to describe the architecture of the A/B Platform. 

This includes creating the following diagrams:

- **System Context Diagram**: Shows the system's interactions with external actors and systems.
- **Container Diagram**: Illustrates the high-level components of the system and their interactions.

### C4 Model Diagrams

#### System Context Diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

title System Context - A/B Platform

Person(user, "Experiment Manager", "Manages test cases and experiments.")
System(ab_platform, "A/B Platform", "Platform to run and manage A/B tests.")
System_Ext(prometheus, "Prometheus", "Collects and monitors platform metrics.")

Rel(user, ab_platform, "Creates and manages test cases")
Rel(ab_platform, prometheus, "Sends experiment metrics for monitoring")
@enduml
```

#### Container Diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

title Container Diagram - A/B Platform

Person(user, "Experiment Manager")
System_Boundary(ab_platform, "A/B Platform Monolith") {
    Container(web_ui, "Web UI", "React", "User interface for managing test cases and experiments")
    Container(api, "API", "Go", "API layer for handling experiment-related requests")
    Container(db, "Database", "PostgreSQL", "Stores test cases, variations, and results")
    Container(metrics_collector, "Metrics Collector", "Prometheus", "Collects and monitors performance metrics")
}

Rel(user, web_ui, "Interacts with")
Rel(web_ui, api, "Submits test cases")
Rel(api, db, "Reads and writes data")
Rel(api, metrics_collector, "Sends experiment metrics")
@enduml
```

## Consequences

- **Visualization**: The C4 model provides a clear and concise way to visualize the architecture of the A/B Platform.
- **Communication**: Diagrams help communicate the system's structure and interactions to stakeholders and team members.
- **Documentation**: The diagrams serve as living documentation that can be updated as the system evolves.
