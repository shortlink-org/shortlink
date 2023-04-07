```plantuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

' Components
!define actors https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/actors
!define ui https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/ui
!define services https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/services
!define boundaries https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/boundaries

!include boundaries/chat.puml
!include boundaries/delivery.puml
!include boundaries/gateway.puml
!include boundaries/integration.puml
!include boundaries/link.puml
!include boundaries/marketing.puml
!include boundaries/notification.puml
!include boundaries/payment.puml
!include boundaries/platform.puml
!include boundaries/search.puml
!include boundaries/shortdb.puml
```
