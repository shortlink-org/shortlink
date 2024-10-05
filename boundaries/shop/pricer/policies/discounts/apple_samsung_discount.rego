package pricing.discount

# 5% discount on Apple and Samsung products
brand_discount {
    brand := input.brand
    discount := if brand == "Apple" || brand == "Samsung" then input.price * 0.05 else 0
}
