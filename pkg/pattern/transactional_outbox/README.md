## Transactional outbox

The **transactional outbox pattern** is a way to ensure that messages are sent to a message broker in a transactional manner. 
This means that the message is only sent if the transaction is committed successfully. 

```plantuml
@startuml C4_Transactional_Outbox

!define C4PlantUML https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master
!include C4PlantUML/C4_Context.puml
!include C4PlantUML/C4_Container.puml
!include C4PlantUML/C4_Component.puml

LAYOUT_WITH_LEGEND()

' Define your system
System_Boundary(s1, "Your System") {
    ContainerDb(db1, "Database", "Relational DBMS", "Stores user data and outbox table")
    Container(app, "Application", "Your favorite technology", "Handles business logic")
    Container(mq, "Message Queue", "RabbitMQ/Kafka", "Buffers messages for asynchronous processing")
}

' Define external systems
System_Ext(externalService, "External Service", "Consumes messages for further processing")

' Show the transactional outbox pattern flow
Rel(app, db1, "Writes to", "1. Performs database transaction including write to outbox table")
Rel(db1, app, "Dispatch", "2. Outbox Poller reads and publishes messages", "Polling mechanism/Transaction log tailing")
Rel_L(app, mq, "3. Publishes messages to", "AMQP/Kafka client")
Rel(mq, externalService, "4. Sends messages to")

@enduml
```

### References

- [Pattern: Transactional outbox](https://microservices.io/patterns/data/transactional-outbox.html)
- [Transactional Outbox Pattern For Guaranteed Delivery In Golang!](https://thegodev.com/transactional-outbox-pattern/)
