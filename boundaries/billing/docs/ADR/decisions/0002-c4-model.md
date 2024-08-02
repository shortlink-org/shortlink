# 2. C4 Model for Billing boundary context

Date: 2024-01-01

## Status

Accepted

## Context

With the Billing Boundary established to efficiently manage financial transactions within our software system, 
it is crucial to have a clear visualization of its architecture. 

This boundary encompasses services such as Billing and Wallet, each with its intricate operations and interactions 
with external payment systems. 

The C4 model, recognized for its effectiveness in illustrating software architectures, offers a structured approach 
to mapping out these components, fostering a better understanding among stakeholders.

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

title System Context diagram for Billing Boundary Context

Person(user, "Customer", "A customer using the system for transactions.")
System_Ext(payment_gateway, "Payment Gateway", "External system to process payments.")
System_Ext(wallet_provider, "Wallet Provider", "External e-wallet service.")
SystemQueue_Ext(fraud_detection, "Fraud Detection System", "External system to detect and prevent fraudulent transactions.")
System_Boundary(billing_boundary, "Billing Boundary Context") {
    System(billing_service, "Billing Service", "Core service for handling billing and invoicing.")
    System(wallet_service, "Wallet Service", "Manages customer e-wallet transactions.")
}

Rel(user, billing_service, "Initiates payments through")
Rel(billing_service, payment_gateway, "Processes payments with")
Rel(user, wallet_service, "Accesses and manages e-wallet with")
Rel(wallet_service, wallet_provider, "Interacts with")
Rel(billing_service, fraud_detection, "Uses for transaction security")
Rel(wallet_service, fraud_detection, "Uses for transaction security")

@enduml
```

#### Level 2: Container diagram

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

title Container diagram for Billing Boundary Context with Improved Composition

Person(user, "Customer", "A customer using the system for transactions.")
Container_Ext(api_gateway, "API Gateway", "API Gateway", "Central entry point for handling financial transaction requests.")
SystemQueue_Ext(mq, "Message Queue", "Handles event-driven financial operations.")
Container_Ext(payment_processor, "Payment Processor", "External Payment Processor", "Processes online payments through various gateways.")
Container_Ext(wallet_api, "Wallet API", "External E-Wallet Service", "Manages e-wallet transactions.")
Container_Ext(fraud_system, "Fraud Detection System", "External System", "Detects and prevents fraudulent transactions.")
System_Boundary(billing_boundary, "Billing Boundary Context") {
    Container(billing_service, "Billing Service", "Go", "Handles billing, invoicing, and payment processing, including interactions with Stripe and PayPal.", "left")
    Container(wallet_service, "Wallet Service", "Go (Solidity)", "Manages customer e-wallet transactions and balances.", "right")
    ContainerDb(billing_db, "Billing Database", "Database", "Stores transaction and invoice data.", "left")
    ContainerDb(wallet_db, "Wallet Database", "Database", "Stores e-wallet balances and transaction history.", "right")
}

Rel(user, api_gateway, "Initiates transaction via")
Rel(api_gateway, billing_service, "Routes payment requests to")
Rel(api_gateway, wallet_service, "Routes wallet queries to")
Rel(mq, billing_service, "Publishes transaction events to")
Rel_R(billing_service, payment_processor, "Processes payments with")
Rel_L(wallet_service, wallet_api, "Interacts with for e-wallet transactions")
Rel(billing_service, fraud_system, "Analyzes transactions with")
Rel(wallet_service, fraud_system, "Verifies transactions with")
Rel_D(billing_service, billing_db, "Reads/Writes transaction data")
Rel_D(wallet_service, wallet_db, "Reads/Writes wallet data")

@enduml
```
