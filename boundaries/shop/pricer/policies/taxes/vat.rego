package pricing.vat

# Calculate the VAT (20%) for each item
vat[item_id] = tax {
    some i
    item := input.items[i]
    tax := item.price * 0.20
    item_id := item.productId
}

# Calculate the total VAT for all items
total_vat = sum([tax | some i; item := input.items[i]; tax := item.price * 0.20])
