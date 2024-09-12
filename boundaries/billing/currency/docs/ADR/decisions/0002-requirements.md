# 2. Requirements for Currency Service

Date: 2024-09-12

## Status

Accepted

## Requirements

### Functional Requirements

1. **Real-Time Currency Conversion**: The service must provide an API for converting between currencies using real-time exchange rates.
2. **Historical Exchange Rate Data**: The service must offer an API for retrieving historical exchange rate data for any currency pair.
3. **Integration with Third-Party Providers**: The service must fetch exchange rate data from external providers (Bloomberg and Yahoo).
4. **Currency Selection**: Users must be able to select their preferred currency for transactions.
5. **Data Storage**: Historical exchange rate data must be stored for future reference and analytics.

### Performance and Scalability Requirements

- The service must handle a high volume of API requests with minimal latency.
- It should scale efficiently to accommodate increasing traffic and usage.

### Non-Functional Requirements

- **Reliability**: Maintain a Service Level Agreement (SLA) of 99.95% uptime.
- **Availability**: Ensure that real-time exchange rates are consistently available by implementing fallbacks between data providers.
- **Security**: Secure the service to prevent unauthorized access to exchange rate data.
- **Compliance**: Ensure compliance with applicable financial data regulations.

## Load and Consumption Calculations

### Load Calculation

> [!NOTE]
>
> **Daily API requests for exchange rates**: Assuming `500,000` API requests per day.  
> **Seconds per day**: `24 * 60 * 60 = 86,400`

> [!TIP]
>
> **Requests Per Second (RPS)** = `500,000 / 86,400` ≈ `5.78 RPS`

### Memory Consumption

> [!NOTE]
>
> **Storage per exchange rate**: Assume each entry requires `100 bytes`.  
> **Daily new data**: `500,000 API requests * 100 bytes/request` = `50 MB/day`.

> [!TIP]
>
> **Yearly Data**: `50 MB/day * 365 days` ≈ `18.25 GB/year`  
> **Five-Year Data**: `18.25 GB/year * 5` ≈ `91.25 GB`

### Considerations

- **Peak Usage**: Plan for peak usage scenarios where traffic spikes, and ensure load balancing strategies are in place.
- **Data Retention**: Establish data retention and archiving policies for historical exchange rate data.
- **Rate Limiting**: Implement rate limiting to manage requests to third-party providers and avoid exceeding API limits.
