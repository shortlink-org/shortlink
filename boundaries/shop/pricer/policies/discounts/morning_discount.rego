package pricing.discount

# $5 off between a configurable time range for all items in the cart
time_discount[item_id] = discount {
    item := input.items[_]
    time := input.time
    start_time := input.params.time_discount_start
    end_time := input.params.time_discount_end
    time >= start_time
    time <= end_time
    discount := input.params.time_discount_value
    item_id := item.productId
}

# Total time-based discount for the entire cart
total_time_discount := sum([discount | discount := time_discount[_]])
