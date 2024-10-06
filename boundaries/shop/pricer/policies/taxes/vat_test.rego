package pricing.vat_test

import data.pricing.vat

# Test 1: Check individual item VAT (20% VAT)
test_individual_item_vat {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200},
            {"productId": "item3", "price": 300}
        ]
    }

    # Evaluate the VAT for each product and ensure it is calculated correctly
    vat_item1 := vat.vat["item1"] with input as input
    vat_item2 := vat.vat["item2"] with input as input
    vat_item3 := vat.vat["item3"] with input as input

    # Assertions
    vat_item1 == 20.0    # 20% of 100
    vat_item2 == 40.0    # 20% of 200
    vat_item3 == 60.0    # 20% of 300
}

# Test 2: Check total VAT for all items
test_total_vat {
    input := {
        "items": [
            {"productId": "item1", "price": 100},
            {"productId": "item2", "price": 200},
            {"productId": "item3", "price": 300}
        ]
    }

    # Evaluate the total VAT and ensure it is calculated correctly
    total := vat.total_vat with input as input

    expected_total_vat := 120.0   # 20% of (100 + 200 + 300)

    # Assertion
    total == expected_total_vat
}
