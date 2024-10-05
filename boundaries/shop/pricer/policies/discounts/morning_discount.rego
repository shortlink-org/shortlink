package pricing.discount

# $5 off between a configurable time range using parameters
time_discount {
    input.time >= input.params.time_discount_start
    input.time <= input.params.time_discount_end
    discount := input.params.time_discount_value
}
