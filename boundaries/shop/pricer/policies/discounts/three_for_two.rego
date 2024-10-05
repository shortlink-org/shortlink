package pricing.discount

# 3 for 2 discount using parameterized quantity condition
three_for_two {
    input.count >= input.params.min_quantity_for_discount
    discount := input.price / input.params.min_quantity_for_discount
}
