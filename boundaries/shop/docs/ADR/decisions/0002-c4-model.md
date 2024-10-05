# 2. C4 Model for Shop boundary context

Date: 2024-01-01

## Status

Accepted

## Context

The Shop Boundary consists of several critical services (Merch, Cart, and Admin) integral to our system's operations 
related to goods and services management. Given the complex interactions and processes handled by these services, 
it is crucial to have a detailed and clear visualization of the architecture. 

The C4 model is renowned for its ability to effectively map and document software architecture, 
making it ideal for our needs to ensure clarity and cohesion across the system.

## Decision

We will apply the C4 model to detail the architecture of the Shop Boundary Context. This includes 
creating System Context, Container, and Component diagrams, and optionally, Class diagrams, 
for each service within the boundary.

## Consequences

By applying the C4 model to the Shop Boundary, we anticipate the following benefits:

+ **Enhanced Understanding:** All stakeholders, from developers to business analysts, will have a clearer understanding of the system architecture.
+ **Improved Communication:** Facilitates better discussions and decision-making regarding changes and enhancements to the system.
+ **Streamlined Development and Maintenance:** With a well-documented architecture, new team members can onboard more quickly, 
and ongoing maintenance can be managed more efficiently.


### C4

#### Level 1: System Context diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml

LAYOUT_WITH_LEGEND()

title System Context diagram for Shop Boundary with External Contexts

Person_Ext(customer, "Customer", "A customer using the online shop.")

System_Boundary(sbs, "Shop Boundary Context") {
    System(ui_service, "UI Service (Next.js)", "UI for the shop boundary")
    System(wundergraph_bff, "BFF (WunderGraph)", "Handles frontend requests via GraphQL and coordinates with backend services.")
    System(admin_service, "Admin Service", "Administers shop settings and user permissions.")
    System(oms_graphql, "OMS-GraphQL", "Service for work with orders via GraphQL API.")
    System(oms_temporal, "OMS (Temporal)", "Service for work with carts and orders.")
    System(email_subscription_service, "Email Subscription (Temporal)", "Handles email subscriptions and notifications.")
    System(feed_service, "Feed Service", "Cron job in Go, generates feeds every 24h and saves them to Minio.")
}

System_Ext(minio_store, "Minio (S3-like block store)", "Stores generated feeds.")

System_Boundary(bbs, "Billing Boundary") {
    System_Ext(billing_service, "Billing Service", "Manages billing and invoices.")
}

System_Boundary(dbs, "Delivery Boundary") {
    System_Ext(delivery_service, "Delivery Service", "Handles logistics and delivery of orders.")
}

Rel(customer, ui_service, "Accesses shop UI through")
Rel(ui_service, wundergraph_bff, "Communicates with via GraphQL")
Rel(wundergraph_bff, oms_graphql, "Coordinates shopping cart and checkout via GraphQL with")
Rel(wundergraph_bff, admin_service, "Admin service managed by administrators via GraphQL")
Rel(oms_graphql, oms_temporal, "Communicates with OMS via gRPC for order management")
Rel(email_subscription_service, customer, "Sends email notifications to")
Rel(oms_temporal, billing_service, "Submits order details to")
Rel(oms_temporal, delivery_service, "Sends order info for delivery to")
Rel(feed_service, oms_graphql, "Fetches data via GraphQL from")
Rel(feed_service, minio_store, "Saves generated feeds to")

@enduml
```

#### Level 2: Container diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title Container diagram for Shop Boundary Context

Person(customer, "Customer", "A customer interacts with the online shopping system.")
Container_Ext(api_gateway, "API Gateway", "API Gateway", "Serves as the central entry point for all incoming requests, directing them to the appropriate services within the system.")
SystemQueue_Ext(mq, "Message Queue", "Handles asynchronous communication and event-driven operations among services.")
Container_Ext(payment_gateway, "Payment Gateway", "External Service", "Securely processes payment transactions and handles financial data exchange.")
System_Boundary(sbs, "Shop Boundary Context") {
    Container(ui_service, "UI Service (Next.js)", "Service", "User interface for customers interacting with the shop.")
    Container(wundergraph_bff, "BFF (WunderGraph)", "Service", "Handles frontend requests via GraphQL and coordinates with backend services.")
    Container(oms_graphql, "OMS-GraphQL", "Service", "Service for work with orders via GraphQL API.")
    Container(oms_temporal, "OMS (Temporal)", "Service", "Service for work with carts and orders using gRPC.")
    Container(admin_service, "Admin Service", "Service", "Administers shop settings, manages user roles and permissions, and performs back-end configuration tasks.")
    Container(email_subscription_service, "Email Subscription Service (Temporal)", "Service", "Manages email subscriptions and notifications.")
    Container(feed_service, "Feed Service", "Service", "Cron job in Go, generates feeds every 24h and saves them to Minio.")
    ContainerDb(shop_db, "Shop Database", "Database", "Central repository for storing all orders, carts, and administrative data.")
    Container(shop_cache, "Shop Cache Server", "Cache", "Improves performance by caching frequently accessed data such as product details and prices.")
    Container_Ext(minio_store, "Minio (S3-like block store)", "External Storage", "Stores generated feeds.")
}

Rel_Left(mq, oms_temporal, "Notifies on cart and order processing events")
Rel_Down(customer, api_gateway, "Submits requests to", "HTTP/HTTPS")
Rel_Down(api_gateway, ui_service, "Routes customer requests to", "HTTP/HTTPS")
Rel(ui_service, wundergraph_bff, "Interacts with for data", "GraphQL")
Rel(wundergraph_bff, oms_graphql, "Coordinates order management via", "GraphQL")
Rel(wundergraph_bff, admin_service, "Routes administrative requests to", "GraphQL")
Rel(oms_graphql, oms_temporal, "Communicates with for order management", "gRPC")
Rel(oms_graphql, shop_db, "Reads and writes data", "SQL")
Rel(admin_service, shop_db, "Reads and writes data", "SQL")
Rel(oms_temporal, shop_cache, "Utilizes for faster data retrieval")
Rel(oms_temporal, payment_gateway, "Connects for payment processing", "API")
Rel_Left(mq, email_subscription_service, "Triggers email notifications")
Rel(feed_service, oms_graphql, "Fetches data via GraphQL from")
Rel(feed_service, minio_store, "Saves generated feeds to")

@enduml
```
