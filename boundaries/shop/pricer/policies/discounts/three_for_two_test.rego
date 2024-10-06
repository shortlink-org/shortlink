package pricing.discount_test

import data.pricing.discount

# Test 1: No discount applied when quantity is less than the minimum required

test_three_for_two_discount_no_discount {
    input := {
        "items": [
            {"productId": "item1", "quantity": 2, "price": 10},  # Unit price: 10
            {"productId": "item2", "quantity": 1, "price": 5}    # Unit price: 5
        ],
        "params": {
            "min_quantity_for_discount": 3
        }
    }
    # Verify that no item qualifies for the discount
    not discount.three_for_two[x] with input as input
    # Verify that the total discount is zero
    total_discount := discount.total_quantity_discount with input as input
    total_discount == 0
}

# Test 2: Discounts applied correctly when quantities meet the minimum required

test_three_for_two_discount_with_discount {
    input := {
        "items": [
            {"productId": "item1", "quantity": 3, "price": 5},   # Unit price: 5
            {"productId": "item2", "quantity": 4, "price": 5}    # Unit price: 5
        ],
        "params": {
            "min_quantity_for_discount": 3
        }
    }
    # Calculate discount for item1
    discount1 := discount.three_for_two[x] with input as {"items": [input.items[0]], "params": input.params}
    discount1 == 5  # 1 set * unit price 5

    # Calculate discount for item2
    discount2 := discount.three_for_two[x] with input as {"items": [input.items[1]], "params": input.params}
    discount2 == 5  # 1 set * unit price 5 (only complete sets count)

    # Verify total discount
    total_discount := discount.total_quantity_discount with input as input
    total_discount == 10  # 5 + 5
}

# Test 3: Discounts applied correctly for multiple items with varying quantities

test_three_for_two_discount_multiple_items {
    input := {
        "items": [
            {"productId": "item1", "quantity": 3, "price": 3},   # Unit price: 3
            {"productId": "item2", "quantity": 6, "price": 3}    # Unit price: 3
        ],
        "params": {
            "min_quantity_for_discount": 3
        }
    }
    # Calculate discount for item1
    discount1 := discount.three_for_two[input.items[0].productId] with input as input
    discount1 == 3  # 1 set * unit price 3

    # Calculate discount for item2
    discount2 := discount.three_for_two[input.items[1].productId] with input as input
    discount2 == 6  # 2 sets * unit price 3

    # Verify total discount
    total_discount := discount.total_quantity_discount with input as input
    total_discount == 9  # 3 + 6
}

# Test 4: Total discount calculation with multiple items

test_total_quantity_discount {
    input := {
        "items": [
            {"productId": "item1", "quantity": 3, "price": 5},   # Unit price: 5
            {"productId": "item2", "quantity": 4, "price": 5},   # Unit price: 5
            {"productId": "item3", "quantity": 2, "price": 5}    # Unit price: 5 (no discount)
        ],
        "params": {
            "min_quantity_for_discount": 3
        }
    }
    # Verify total discount
    total_discount := discount.total_quantity_discount with input as input
    total_discount == 10  # 5 (item1) + 5 (item2) + 0 (item3)
}
