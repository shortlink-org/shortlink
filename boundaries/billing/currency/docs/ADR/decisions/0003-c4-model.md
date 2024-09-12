# 3. C4 Model

Date: 2024-09-12

## Status

Accepted

## Context

To provide clarity in the architecture of the new currency conversion service, we will utilize the C4 model. 
This model will help visualize the system’s structure and components, allowing developers and stakeholders to easily 
understand how the system is built and how different parts interact. The C4 model will include System Context 
and Container diagrams, with the addition of a cache store to optimize performance.

## Decision

We will apply the C4 model to the Currency Service to create System Context and Container diagrams to better visualize 
the architecture. A cache layer will be added to store frequently accessed exchange rates and reduce the load on external 
APIs and databases.

## Consequences

- Improved understanding of the overall architecture and the integration with external services like Bloomberg and Yahoo.
- Better communication between team members and stakeholders.
- Increased efficiency in future development and maintenance of the service.

### C4

#### Level 1: System Context Diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

LAYOUT_WITH_LEGEND()

title System Context diagram for Currency Service

Person(user, "User", "A user of the billing system selecting currency for transactions.")
System_Ext(bloomberg_api, "Bloomberg API", "Provides real-time exchange rates.")
System_Ext(yahoo_api, "Yahoo API", "Provides real-time exchange rates.")
System_Boundary(currency_service, "Currency Service") {
    System(currency_conversion, "Currency Conversion Service", "Converts currency in real-time and provides exchange rates.")
}

Rel(user, currency_conversion, "Requests currency conversion or exchange rates")
Rel_U(bloomberg_api, currency_conversion, "Fetches exchange rates from")
Rel_U(yahoo_api, currency_conversion, "Fetches exchange rates from")

@enduml
```

#### Level 2: Container Diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title Container diagram for Currency Service

Person(user, "User", "A user of the billing system selecting currency for transactions.")
Container_Ext(bloomberg_api, "Bloomberg API", "API", "Provides real-time exchange rates.")
Container_Ext(yahoo_api, "Yahoo API", "API", "Provides real-time exchange rates.")
System_Boundary(currency_service, "Currency Service") {
    Container(currency_api, "Currency API", "API", "Exposes endpoints for real-time and historical exchange rates.")
    Container(exchange_handler, "Exchange Handler", "Service", "Handles currency conversion logic and manages cache.")
    Container(cache_store, "Cache Store", "Cache", "Stores frequently accessed exchange rates for quick retrieval.")
    ContainerDb(rate_db, "Rate Database", "Database", "Stores historical exchange rate data.")
}

Rel(user, currency_api, "Requests rates through")
Rel(currency_api, exchange_handler, "Delegates conversion to")
Rel(exchange_handler, cache_store, "Checks cache for frequently accessed rates")
Rel(exchange_handler, rate_db, "Stores and retrieves historical rates from")
Rel(exchange_handler, bloomberg_api, "Fetches real-time rates from")
Rel(exchange_handler, yahoo_api, "Fetches real-time rates from")

@enduml
```
