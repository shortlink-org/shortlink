```plantuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Container.puml

' Components
!define actors https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/actors
!define ui https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/ui

!include ui/landing/landing.puml

!include ui/next/next.puml
!include ui/next/gateway.puml

!include ui/ui_kit/ui_kit.puml
```
