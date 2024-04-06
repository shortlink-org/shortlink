## Geolocation service

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> The Geolocation service is a part of the Delivery Boundary in the ShortLink project.
> This service is responsible for handling various location-based functionalities.

### Usecases

- **Real-Time Location Tracking**: Track the geographic location of devices in real-time.
- **Geofencing**: Define geographical boundaries and get alerts when a device enters or exits these areas.
- **Location History**: Maintain a history of location data for analysis and reporting.

#### System context diagram

```plantuml
@startuml ShortLink Project - Geolocation Service System Context Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

Person(customer, "Customer", "Uses the Geolocation Service through a Web or Mobile UI")
Person(admin, "Administrator", "Manages and monitors the Geolocation Service")

System_Boundary(geolocation_service, "Geolocation Service") {
    System(api, "Geolocation API", "Provides an interface for geolocation data")
}

System_Ext(web_ui, "Web UI", "Web Application used by Customers")
System_Ext(mobile_app, "Mobile App", "Mobile Application used by Customers")
System_Ext(database, "External Database", "Stores user and geolocation data")
System_Ext(third_party_services, "Third-Party Services", "e.g., Map APIs, Notification Services")

Rel(customer, web_ui, "Uses")
Rel(customer, mobile_app, "Uses")
Rel(web_ui, api, "Sends requests to")
Rel(mobile_app, api, "Sends requests to")
Rel(api, database, "Reads/Writes data to")
Rel(api, third_party_services, "Interacts with")

SHOW_LEGEND()
@enduml
```

#### Container diagram

```plantuml
@startuml ShortLink Project - Geolocation Service Container Diagram
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

LAYOUT_WITH_LEGEND()

Person(customer, "Customer", "Interacts with the Geolocation API")
Container_Boundary(geolocation_service, "Geolocation Service") {
    ContainerDb(postgres, "PostgreSQL Database", "SQL Database", "Stores geolocation data and histories")
    Container(geolocation_logic, "Geolocation Service Logic", "Go", "Processes and handles geolocation logic")
    Container(api, "Geolocation API", "Go, gRPC", "Provides gRPC interface for geolocation data")
    Container(ui, "Geolocation UI", "Web Application", "User interface for interacting with Geolocation API")
}

Rel(customer, api, "Uses", "gRPC")
Rel(api, geolocation_logic, "Calls", "gRPC")
Rel(geolocation_logic, postgres, "Reads/Writes", "SQL")
Rel(ui, api, "Interacts", "HTTP/WebSocket")

SHOW_LEGEND()
@enduml
```
