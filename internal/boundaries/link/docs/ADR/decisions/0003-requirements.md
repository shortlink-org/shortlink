# 3. Requirements and Consumption Calculations

Date: 2024-01-02

## Status

Accepted

## Requirements

### Functional Requirements

1. Creation of tiny URLs from long URLs.
2. Redirection from tiny URLs to their original long URLs.
3. Analytics for tiny URL usage.

### Performance and Scalability Requirements

* Handle high traffic with minimal latency.
* Efficiently scale to accommodate increasing load.

### Non-Functional Requirements

* Support 50 million daily active users (DAU).
* Maintain a Service Level Agreement (SLA) of 99.95%.
* Anticipate each user creating one tiny URL per week.
* An average of 20 accesses per tiny URL daily.
* Operational lifespan of the service: 5 years.

## Load and Memory Consumption Calculations

### Load Calculation

> [!NOTE]
>
> Daily URL creations = `50,000,000 / 7` ≈ `7,142,857`  
> Daily URL accesses = `50,000,000 * 20` = `1,000,000,000`  
> Total daily requests = `7,142,857 + 1,000,000,000` = `1,007,142,857`  
> Seconds per day = `24 * 60 * 60` = `86,400`
 
> [!TIP]
>
> **Requests Per Second (RPS)** = `1,007,142,857 / 86,400` ≈ `11,655`

### Memory Consumption

> [!NOTE]
> 
> **Storage Per URL**: `50 bytes`  
> **Daily New Data**: `7,142,857 URLs/day * 50 bytes/URL` = `357,142,850 bytes/day` ≈ `0.357 GB/day`  

> [!TIP]
>
> **Yearly Data**: `0.357 GB/day * 365 days` ≈ `130.3 GB/year`  
> **Five-Year Data**: `130.3 GB/year * 5` ≈ `651.5 GB`

### Analytics Calculation
> [!TIP]
>
> **Daily Analytics for Each URL**: `20 accesses/URL/day * 50 bytes/access` = `1,000 bytes/URL/day`
>
> **Total Daily Analytics for All URLs**: `7,142,857 URLs/day * 1,000 bytes/URL/day` ≈ `7,142,857,000 bytes/day` ≈ `7.14 GB/day`
>
> **Total Analytics for 5 Years**: `7.14 GB/day * 365 days/year * 5 years` ≈ `13,061 GB` ≈ `13.06 TB`

### Considerations

* Potential peak usage scenarios and strategies for load balancing.
* Data retention policies and archiving strategies for long-term sustainability.
