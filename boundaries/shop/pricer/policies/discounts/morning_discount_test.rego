package pricing.discount_test

import data.pricing.discount

# Test 1: $5 off for all items when time is within the valid range
test_time_discount_within_range {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200}
        ],
        "time": "10:30",
        "params": {
            "time_discount_start": "09:00",
            "time_discount_end": "11:00",
            "time_discount_value": 5
        }
    }

    # Evaluate the time discount for each item
    discount_item1 := discount.time_discount["item1"] with input as input
    discount_item2 := discount.time_discount["item2"] with input as input

    # Assertions: $5 discount applied to each item
    discount_item1 == 5.0
    discount_item2 == 5.0
}

# Test 2: Total time-based discount for all items when time is within range
test_total_time_discount_within_range {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200},
            {"productId": "item3", "price": 300}
        ],
        "time": "10:30",
        "params": {
            "time_discount_start": "09:00",
            "time_discount_end": "11:00",
            "time_discount_value": 5
        }
    }

    # Evaluate the total time-based discount for all items
    total_discount := discount.total_time_discount with input as input

    # Assertions: $5 discount applied to each of the 3 items = 15
    total_discount == 15.0
}

# Test 3: No discount when time is outside the valid range
test_no_time_discount_outside_range {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200}
        ],
        "time": "12:00",  # Outside of the valid time range
        "params": {
            "time_discount_start": "09:00",
            "time_discount_end": "11:00",
            "time_discount_value": 5
        }
    }

    # Ensure no discount is applied outside the valid time range
    not discount.time_discount["item1"] with input as input
    not discount.time_discount["item2"] with input as input
}
