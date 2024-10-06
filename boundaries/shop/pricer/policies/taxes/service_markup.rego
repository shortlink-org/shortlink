package pricing.tax

# Calculate the 5% markup for each item
service_markup[item_id] = tax {
    some i
    item := input.items[i]
    tax := item.price * 0.05
    item_id := item.productId
}

# Calculate the total markup for all items
total_markup = sum([tax | some i; item := input.items[i]; tax := item.price * 0.05])
