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

Person(customer, "Customer", "A customer using the online shop.")

System_Boundary(sbs, "Shop Boundary Context") {
    System(merch_service, "Merch Service", "Handles merchandise listings and inventory.")
    System(admin_service, "Admin Service", "Administers shop settings and user permissions.")
    System(cart_service, "Cart Service", "Manages customer shopping cart and checkout process.")
}

System_Boundary(bbs, "Billing Boundary") {
    System_Ext(billing_service, "Billing Service", "Manages billing and invoices.")
}

System_Boundary(dbs, "Delivery Boundary") {
    System_Ext(delivery_service, "Delivery Service", "Handles logistics and delivery of orders.")
}

Rel(customer, merch_service, "Browses and orders merch from")
Rel(customer, cart_service, "Adds items to and checks out via")
Rel(customer, admin_service, "Managed by administrators for")
Rel_D(cart_service, bbs, "Submits order details to")
Rel_D(cart_service, dbs, "Sends order info for delivery to")

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
    Container(merch_service, "Merch Service", "Service", "Provides merchandise listings, manages inventory, and handles product updates.")
    Container(cart_service, "Cart Service", "Service", "Responsible for shopping cart management, session tracking, and checkout operations.")
    Container(admin_service, "Admin Service", "Service", "Administers shop settings, manages user roles and permissions, and performs back-end configuration tasks.")
    ContainerDb(shop_db, "Shop Database", "Database", "Central repository for storing all merchandise, cart, and administrative data.")
    Container(shop_cache, "Shop Cache Server", "Cache", "Improves performance by caching frequently accessed data such as product details and prices.")
}

Rel_D(mq, merch_service, "Sends inventory updates and order confirmations")
Rel_D(mq, cart_service, "Notifies on cart updates and checkout events")
Rel(customer, api_gateway, "Submits requests to", "HTTP/HTTPS")
Rel(api_gateway, merch_service, "Routes requests to", "HTTP/HTTPS")
Rel(api_gateway, cart_service, "Routes requests to", "HTTP/HTTPS")
Rel(api_gateway, admin_service, "Routes requests to", "HTTP/HTTPS")
Rel(merch_service, shop_db, "Reads and writes data", "SQL")
Rel(cart_service, shop_db, "Reads and writes data", "SQL")
Rel(admin_service, shop_db, "Reads and writes data", "SQL")
Rel(merch_service, shop_cache, "Utilizes for faster data retrieval")
Rel(cart_service, payment_gateway, "Connects for payment processing", "API")

@enduml
```
