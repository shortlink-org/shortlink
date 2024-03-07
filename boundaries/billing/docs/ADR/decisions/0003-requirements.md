# 3. Requirements and Consumption Calculations

Date: 2024-03-07

## Status

Accepted

## Requirements

### Functional Requirements

1. Secure processing of financial transactions through various payment gateways, including Stripe and PayPal.
2. Management of customer e-wallets for transactions within the system.
3. Fraud detection and prevention in financial transactions.
4. Handling recurring payments for subscriptions.
5. Generating invoices and receipts for transactions.

### Performance and Scalability Requirements

* Ensure transaction processing with minimal latency.
* Scale dynamically to handle varying load, especially during peak shopping seasons.

### Non-Functional Requirements

* Support secure transactions for up to 1,000 daily transactions.
* Achieve a Service Level Agreement (SLA) of 99.9% uptime.
* Implement robust fraud detection mechanisms to minimize unauthorized transactions.
* Ensure data integrity and confidentiality of customer financial information.
* Provide real-time transaction processing capabilities.

## Load and Memory Consumption Calculations

### Load Calculation

Given the estimate of around 1,000 transactions per day:

> [!NOTE]
>
> **Daily Transactions**: `1,000`  
> **Seconds per day**: `24 * 60 * 60` = `86,400`

> [!TIP]
>
> **Transactions Per Second (TPS)** = `1,000 / 86,400` ≈ `0.0116`

### Memory Consumption

Assuming each transaction record takes approximately `500 bytes` (including transaction details, customer information, and payment method):

> [!NOTE]
>
> **Storage Per Transaction**: `500 bytes`  
> **Daily New Data**: `1,000 transactions/day * 500 bytes/transaction` = `500,000 bytes/day` ≈ `0.5 MB/day`

> [!TIP]
>
> **Yearly Data**: `0.5 MB/day * 365 days` ≈ `182.5 MB/year`  
> **Five-Year Data**: `182.5 MB/year * 5` ≈ `912.5 MB`

### Considerations

* Planning for peak times, such as holiday seasons, where transaction volume may significantly increase, 
  requiring effective scaling and load-balancing strategies.
* Implementing data retention policies that balance between legal compliance, customer service, and storage optimization.
* Ensuring infrastructure and application security to protect against breaches and fraud.
* Designing the system for high availability and resilience, including fallback mechanisms for payment gateways 
  and real-time fraud detection updates.
