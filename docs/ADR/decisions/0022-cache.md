# 22. Cache Strategy [common]

Date: 2023-12-29

## Status

Accepted

## Context

We are addressing the need for efficient service response times, reduced load on external services, 
and improved data reuse and system resilience. This decision is influenced by a balance between speed, data freshness, 
and system resilience.

> [!NOTE]
> A comprehensive understanding of caching's impact on system performance and stability is key.

## Decision

Our strategy will integrate both internal and external caching methods to optimize performance and scalability.

### Strategy Components

- **Internal Caching**:
  - *Use Case*: For data requiring rapid access.
  - *Benefits*: High speed, no network delays.
- **External Caching**:
  - *Use Case*: For managing larger data volumes.
  - *Benefits*: Ease scalability.

> [!TIP]
> We will employ a mix of 'Cache Aside', 'Cache Through', and 'Cache Ahead' strategies, depending on the specific requirements of each use case.

### Data Eviction Policies

| Policy        | Description                                                                                                                                                       |
|---------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| FIFO          | First-In-First-Out: The oldest data in the cache is evicted first.                                                                                                |
| LIFO          | Last-In-First-Out: The most recently added data is evicted first.                                                                                                 |
| Random        | Random selection for eviction.                                                                                                                                    |
| **LRU**       | Least Recently Used: Data not accessed for the longest time is evicted.                                                                                           |
| MRU           | Most Recently Used: The most recently accessed data is evicted. This is a specific case where older data is preserved over newer.                                 |
| LFU           | Least Frequently Used: Data that is least often accessed is evicted.                                                                                              |
| Belady's OPT  | Optimal Page Replacement: Theoretically the best page replacement algorithm, it evicts pages that will not be used for the longest period in the future.          |
| Second Chance | A modification of FIFO, giving pages a "second chance" if they've been accessed recently.                                                                         |
| Clock         | Similar to Second Chance, organizes pages in a circular queue and gives a second chance before eviction.                                                          |
| 2Q            | Maintains two separate queues, one for recently accessed pages and another for frequently accessed pages, combining recency and frequency aspects of LRU and LFU. |
| SLRU          | Segmented LRU: Divides the cache into two segments, a probationary and a protected segment, to differentiate between frequently and recently used data.           |
| TLRU          | Time Aware LRU: An extension of LRU that considers the age of the data, not just the usage pattern.                                                               |
| LRU-k         | An extension of LRU, maintains the times of the last k references to each page, and uses this data for making eviction decisions.                                 |

> [!NOTE]
> Selecting the right eviction policy is crucial as it affects the efficiency of the caching mechanism. It should align with the specific data access patterns of the application.

## Error Caching

Caching errors can be a strategic approach to further reduce load and avoid repeated fetching of erroneous data.

- **Implementation**: Store error responses in the cache.
- **Benefits**: Subsequent requests will retrieve the error from the cache, preventing an unnecessary load on the data source.
- **Mitigation of Cache Miss Attacks**: Prevents repeated cache misses caused by requests for data that is known to be erroneous.

> [!WARNING]
> Care must be taken to set appropriate expiry times for cached errors to avoid persisting stale error states.

## Effectiveness of Caching

To ensure that the caching system is effective, we will regularly measure its performance.

### Average Response Time Formula

```
AverageTime = DBAccessTime * CacheMissRate + CacheAccessTime
```

- **DBAccessTime**: Time to retrieve data from the database (e.g., 100ms).
- **CacheAccessTime**: Time to retrieve data from the cache (e.g., 20ms).
- **CacheMissRate**: The percentage of cache misses (e.g., 0.1).

A high `CacheMissRate` (e.g., > 0.8) indicates that caching may be counterproductive.

#### Monitoring with Prometheus Metrics

We will use Prometheus metrics to monitor key indicators such as cache hit-and-miss rates, and response times. 
This will help us in identifying performance issues and making the necessary adjustments.

## Consequences

### Benefits

- **Enhanced Service Response**: Significantly faster responses.
- **Reduced External Service Load**: Lower strain on external systems.

### Challenges and Mitigation

> [!WARNING]
> This strategy introduces complexities in cache management, particularly around maintaining data freshness and cache warming post-downtimes.

- **Regular Monitoring**: Essential for identifying and resolving cache pollution and stale data.
- **Parameter Adjustment**: Continuous tuning of cache settings to ensure effectiveness.

> [!NOTE]
> Ongoing assessment and adaptation to changing data patterns and usage scenarios are crucial for the success of our caching strategy.
