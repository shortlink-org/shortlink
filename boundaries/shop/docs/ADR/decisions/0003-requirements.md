# 3. Requirements and Consumption Calculations

Date: 2024-01-02

## Status

Accepted

## Requirements

### Functional Requirements

1. Secure processing of approximately 100 transactions daily.
2. Efficient handling of shop inventory, user management, and order processing.
3. Reliable order tracking and analytics for transactions.

### Performance and Scalability Requirements

* Ensure system stability and responsiveness under the daily transaction load.
* Scale efficiently during peak retail periods such as holidays and promotions.

### Non-Functional Requirements

* Maintain a high level of data integrity and transaction security.
* Ensure a Service Level Agreement (SLA) of 99.95% uptime.
* Designed to support future expansion in products and transaction volume.

## Load and Memory Consumption Calculations

### Load Calculation

> [!NOTE]
>
> **Daily Transactions:** `100 transactions/day`
> **Seconds per day:** `24 * 60 * 60` = `86,400`

> [!TIP]
>
> **Transactions Per Second (TPS)** = `100 transactions / 86,400` ≈ `0.00116 TPS`

### Memory Consumption

> [!NOTE]
>
> **Storage Per Transaction Record**: `2 KB`  
> **Daily New Data**: `100 transactions/day * 2 KB/transaction` = `200 KB/day`

> [!TIP]
>
> **Yearly Data**: `200 KB/day * 365 days` ≈ `73 MB/year`
> **Five-Year Data**: `73 MB/year * 5` ≈ `365 MB`

### Considerations

* The system should be capable of handling sudden surges in transactions, especially during promotional events.
* Data backup and recovery plans must be robust to ensure continuity and security of transaction data.
