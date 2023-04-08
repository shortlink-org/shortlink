## Use Case: Add item to basket

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml

actor customer

participant "Merch" as merch
participant "Basket" as basket

customer -> merch : Add item to basket
note right of merch #LightBlue
    - Item ID
    - Quantity
end note

merch -> basket : Add item to basket
note right of basket #LightBlue
    - Item ID
    - Quantity
end note

merch <-- basket : Added item to basket
note left of merch #LightBlue
    - Basket ID
    - Item ID
    - Quantity
end note

customer <-- merch : Added item to basket
note left of customer #LightBlue
    - Basket ID
    - Item ID
    - Quantity
end note
```
