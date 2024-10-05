package pricing.discount

# $5 off between 9am and 11am
time_discount {
    now := input.time
    discount := if now >= "09:00" && now <= "11:00" then 5 else 0
}
