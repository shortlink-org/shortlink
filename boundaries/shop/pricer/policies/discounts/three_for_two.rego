package pricing.discount

# "3 for 2" discount for items with quantity >= 3 in the cart
three_for_two[item_id] = discount {
    item := input.items[_]
    item.quantity >= input.params.min_quantity_for_discount
    sets := floor(item.quantity / input.params.min_quantity_for_discount)
    discount := sets * item.price  # item.price is the unit price
    item_id := item.productId
}

# Total "3 for 2" discount for qualifying items
total_quantity_discount := sum([discount | discount := three_for_two[_]])
