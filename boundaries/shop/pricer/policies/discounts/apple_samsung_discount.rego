package pricing.discount

# 5% discount on Apple products
brand_discount[item_id] = discount {
    item := input.items[_]
    item.brand == "Apple"
    discount := item.price * input.params.apple_samsung_discount
    item_id := item.productId
}

# 5% discount on Samsung products
brand_discount[item_id] = discount {
    item := input.items[_]
    item.brand == "Samsung"
    discount := item.price * input.params.apple_samsung_discount
    item_id := item.productId
}

# Total discount for all Apple and Samsung products
total_brand_discount := sum([discount | discount := brand_discount[_]])
