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
* An Average of 20 accesses per tiny URL daily.
* Operational lifespan of the service: 5 years.

## Load and Memory Consumption Calculations

### Load Calculation

1. **Daily Active Users (DAU)**: 50,000,000
2. **Tiny URL Creations Per Week**: Each user creates one tiny URL per week, so in a day, \( \frac{50,000,000}{7} \approx 7,142,857 \) tiny URLs are created.
3. **Tiny URL Accesses Per Day**: Each tiny URL is accessed 20 times, which translates to \( 7,142,857 \times 20 = 142,857,140 \) accesses.
4. **Total Daily Requests**: Sum of creations and accesses is \( 7,142,857 + 142,857,140 = 150,000,000 \).
5. **Requests Per Second (RPS)**: For average load, divide total daily requests by seconds in a day, \( \frac{150,000,000}{86,400} \approx 1,736 \) RPS.

### Memory Consumption

1. **Storage Per URL**: Assuming an average of 500 bytes per URL (including metadata and tiny URL).
2. **Daily New Data**: \( 7,142,857 \times 500 \) bytes.
3. **Yearly Data**: \( 7,142,857 \times 500 \times 365 \) bytes.
4. **Five-Year Data**: Multiply yearly data by 5.

### Analytics Calculation

1. **Click Data Per URL**: Assuming 50 bytes per click.
2. **Daily Click Data**: \( 142,857,140 \times 50 \) bytes.
3. **Storage for 5 Years**: Daily click data multiplied by 365 days and then by 5 years.

These calculations provide a basic understanding of the load and memory requirements.

### Considerations

* Potential peak usage scenarios and strategies for load balancing.
* Data retention policies and archiving strategies for long-term sustainability.
