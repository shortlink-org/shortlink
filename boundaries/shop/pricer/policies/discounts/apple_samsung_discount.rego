package pricing.discount

# 5% discount on Apple and Samsung products using parameterized discount rate
brand_discount {
    input.brand == "Apple"
    discount := input.price * input.params.apple_samsung_discount
}

brand_discount {
    input.brand == "Samsung"
    discount := input.price * input.params.apple_samsung_discount
}
