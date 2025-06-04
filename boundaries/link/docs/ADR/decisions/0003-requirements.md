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
> **Daily URL creations:** `5,000,000 / 7` ≈ `714,285`
> **Daily URL accesses:** `5,000,000 * 20` = `100,000,000`
> **Total daily requests:** `714,285 + 100,000,000` = `100,714,285`
> **Seconds per day:** `24 * 60 * 60` = `86,400`
 
> [!TIP]
>
> **Requests Per Second (RPS)** = `100,714,285 / 86,400` ≈ `1,165 RPS`

### Memory Consumption

> [!NOTE]
> 
> **Storage Per URL**: `50 bytes`  
> **Daily New Data**: `714,285 URLs/day * 50 bytes/URL` = `35.7 MB/day` 

> [!TIP]
>
> **Yearly Data**: `35.7 MB/day * 365 days` ≈ `13 GB/year` 
> **Five-Year Data**: `13 GB/year * 5` ≈ `65 GB`

### Considerations

* Potential peak usage scenarios and strategies for load balancing.
* Data retention policies and archiving strategies for long-term sustainability.
