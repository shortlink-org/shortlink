package pricing.discount_test

import data.pricing.discount

# Test 1: General 10% discount on individual items
test_general_discount {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200},
            {"productId": "item3", "price": 300}
        ]
    }

    # Evaluate the general discount for each item
    discount_item1 := discount.general_discount["item1"] with input as input
    discount_item2 := discount.general_discount["item2"] with input as input
    discount_item3 := discount.general_discount["item3"] with input as input

    # Assertions: 10% discount applied
    discount_item1 == 10.0   # 10% of 100
    discount_item2 == 20.0   # 10% of 200
    discount_item3 == 30.0   # 10% of 300
}

# Test 2: Total general discount for all items in the cart
test_total_general_discount {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200},
            {"productId": "item3", "price": 300}
        ]
    }

    # Evaluate the total general discount for all items
    total_discount := discount.total_general_discount with input as input

    # Assertions: Total discount should be 10% of (100 + 200 + 300) = 60
    total_discount == 60.0
}
