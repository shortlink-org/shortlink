package pricing.discount

# 3 for 2 discount
three_for_two {
    count := input.count
    discount := if count >= 3 then input.price / 3 else 0
}
