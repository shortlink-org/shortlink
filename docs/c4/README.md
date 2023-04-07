### Context diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml

LAYOUT_TOP_DOWN()
LAYOUT_WITH_LEGEND()

!include actors/customer.puml
!include boundaries/gateway.puml

customer --> gatewayBoundary
```
