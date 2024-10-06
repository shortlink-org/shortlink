package pricing.discount

# General 10% discount on all items in the cart
general_discount[item_id] = discount {
    item := input.items[_]
    discount := item.price * 0.10
    item_id := item.productId
}

# Total general discount for all items in the cart
total_general_discount := sum([discount | discount := general_discount[_]])
