## Use Case: Remove item from basket

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml

actor customer

participant "Merch" as merch
participant "Basket" as basket

customer -> merch : Remove item from basket
note right of merch #LightBlue
    - Basket ID
    - Item ID
end note

merch -> basket : Remove item from basket
note right of basket #LightBlue
    - Basket ID
    - Item ID
end note

merch <-- basket : Removed item from basket
note left of merch #LightBlue
    - Basket ID
    - Item ID
end note

customer <-- merch : Removed item from basket
note left of customer #LightBlue
    - Basket ID
    - Item ID
end note
```
