```plantuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

' Components
!define actors https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/actors
!define ui https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/ui
!define services https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/services

!include services/billing/service.puml
!include services/billing/db.puml
```
