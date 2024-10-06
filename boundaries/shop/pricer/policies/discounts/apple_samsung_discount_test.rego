package pricing.discount_test

import data.pricing.discount

# Test 1: Apply 5% discount for Apple products
test_apple_discount {
    input := {
        "items": [
            {"productId": "apple1", "brand": "Apple", "price": 100},
            {"productId": "apple2", "brand": "Apple", "price": 200}
        ],
        "params": {
            "apple_samsung_discount": 0.05
        }
    }

    # Evaluate the discount for each Apple product
    discount_apple1 := discount.brand_discount["apple1"] with input as input
    discount_apple2 := discount.brand_discount["apple2"] with input as input

    # Assertions: 5% discount applied
    discount_apple1 == 5.0   # 5% of 100
    discount_apple2 == 10.0  # 5% of 200
}

# Test 2: Apply 5% discount for Samsung products
test_samsung_discount {
    input := {
        "items": [
            {"productId": "samsung1", "brand": "Samsung", "price": 300},
            {"productId": "samsung2", "brand": "Samsung", "price": 400}
        ],
        "params": {
            "apple_samsung_discount": 0.05
        }
    }

    # Evaluate the discount for each Samsung product
    discount_samsung1 := discount.brand_discount["samsung1"] with input as input
    discount_samsung2 := discount.brand_discount["samsung2"] with input as input

    # Assertions: 5% discount applied
    discount_samsung1 == 15.0   # 5% of 300
    discount_samsung2 == 20.0   # 5% of 400
}

# Test 3: No discount for non-Apple/Samsung products
test_no_discount_for_other_brands {
    input := {
        "items": [
            {"productId": "other1", "brand": "OtherBrand", "price": 500},
            {"productId": "other2", "brand": "OtherBrand", "price": 600}
        ],
        "params": {
            "apple_samsung_discount": 0.05
        }
    }

    # Ensure no discount is applied for non-Apple/Samsung products
    not discount.brand_discount["other1"] with input as input
    not discount.brand_discount["other2"] with input as input
}

# Test 4: Total discount for all Apple and Samsung products
test_total_brand_discount {
    input := {
        "items": [
            {"productId": "apple1", "brand": "Apple", "price": 100},
            {"productId": "apple2", "brand": "Apple", "price": 200},
            {"productId": "samsung1", "brand": "Samsung", "price": 300},
            {"productId": "other1", "brand": "OtherBrand", "price": 400}
        ],
        "params": {
            "apple_samsung_discount": 0.05
        }
    }

    # Evaluate the total discount for all Apple and Samsung products
    total_discount := discount.total_brand_discount with input as input

    # Assertions: Total discount should be 5% of (100 + 200 + 300) = 30
    total_discount == 30.0
}
