package pricing.tax_test

import data.pricing.tax

# Test 1: Check individual item markup (5% markup)
test_individual_item_markup {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200},
            {"productId": "item3", "price": 300}
        ]
    }

    # Evaluate the service_markup for each product and ensure it is calculated correctly
    service_markup_item1 := tax.service_markup["item1"] with input as input
    service_markup_item2 := tax.service_markup["item2"] with input as input
    service_markup_item3 := tax.service_markup["item3"] with input as input

    # Assertions
    service_markup_item1 == 5.0    # 5% of 100
    service_markup_item2 == 10.0   # 5% of 200
    service_markup_item3 == 15.0   # 5% of 300
}

# Test 2: Check total markup for all items
test_total_markup {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200},
            {"productId": "item3", "price": 300}
        ]
    }

    # Evaluate the total_markup and ensure it is calculated correctly
    total := tax.total_markup with input as input

    expected_total_markup := 30.0   # 5% of (100 + 200 + 300)

    # Assertion
    total == expected_total_markup
}
