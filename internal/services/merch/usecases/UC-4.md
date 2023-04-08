## Use Case: View items in basket

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml

actor customer

participant "Merch" as merch
participant "Basket" as basket

customer -> merch : View items in basket
note right of merch #LightBlue
    - Basket ID
end note

merch -> basket : View items in basket
note right of basket #LightBlue
    - Basket ID
end note

merch <-- basket : Items in basket
note left of merch #LightBlue
    - Basket ID
    - Items
end note

customer <-- merch : Items in basket
note left of customer #LightBlue
    - Basket ID
    - Items
end note
```
