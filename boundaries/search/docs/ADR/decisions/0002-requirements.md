# 3. Requirements and Consumption Calculations for Search Service

Date: 2024-02-17

## Status

Accepted

## Requirements

### Functional Requirements

1. Full-text search across all documents in the system.
2. Real-time indexing of new or updated documents.
3. Advanced search capabilities, including filtering and sorting.

### Performance and Scalability Requirements

* Quick response times for search queries, even under high load.
* Ability to efficiently scale with growing data volumes and user base.

### Non-Functional Requirements

* Support for querying across millions of documents with minimal latency.
* Maintain a Service Level Agreement (SLA) of 99.9% uptime.
* Capability to handle 100 million search queries per day.
* Ensure data consistency and accuracy in search results.
* Operational lifespan of the search service: 5+ years.

## Load and Memory Consumption Calculations

### Load Calculation

> [!NOTE]
>
> **Daily search queries:** `100,000,000`  
> **Seconds per day:** `24 * 60 * 60` = `86,400`

> [!TIP]
>
> **Queries Per Second (QPS)** = `100,000,000 / 86,400` ≈ `1,157`

### Memory Consumption

> [!NOTE]
>
> **Average Document Size**: `10 KB`  
> **Daily Indexing Volume**: `1,000,000 documents/day * 10 KB/document` = `10,000,000 KB/day` ≈ `9.537 GB/day`

> [!TIP]
>
> **Yearly Data Indexed**: `9.537 GB/day * 365 days` ≈ `3,481 GB/year`  
> **Five-Year Data Indexed**: `3,481 GB/year * 5` ≈ `17,405 GB`

### Considerations

* Elastic scalability to manage sudden spikes in search queries.
* Efficient indexing strategies to minimize latency and optimize query performance.
* Data archiving and purging strategies to manage the index size and ensure system performance over time.
