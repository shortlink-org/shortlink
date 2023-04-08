## Use Case: Update item quantity in basket

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml

actor customer

participant "Merch" as merch
participant "Basket" as basket

customer -> merch : Update item quantity in basket
note right of merch #LightBlue
    - Basket ID
    - Item ID
    - Quantity
end note

merch -> basket : Update item quantity in basket
note right of basket #LightBlue
    - Basket ID
    - Item ID
    - Quantity
end note

merch <-- basket : Updated item quantity in basket
note left of merch #LightBlue
    - Basket ID
    - Item ID
    - Quantity
end note

customer <-- merch : Updated item quantity in basket
note left of customer #LightBlue
    - Basket ID
    - Item ID
    - Quantity
end note
```
