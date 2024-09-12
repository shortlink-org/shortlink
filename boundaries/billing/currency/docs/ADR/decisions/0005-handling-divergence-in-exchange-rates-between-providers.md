# 5. Handling Divergence in Exchange Rates Between Providers

Date: 2024-09-12

## Status

Accepted

## Context

In the Currency Service, exchange rates for a given currency pair may diverge between external providers (e.g., Bloomberg and Yahoo). 
This divergence can lead to inconsistent results for users. 
To handle this, we need a strategy that ensures both reliability and transparency.

## Decision

We will adopt the **Weighted Average Approach** for determining the exchange rate when there is a divergence between 
the rates provided by Bloomberg and Yahoo. Additionally, both rates will be **stored and shown for audit purposes**.

### Details

1. **Weighted Average Approach**:  
  - Each provider (Bloomberg and Yahoo) will be assigned a weight based on their perceived reliability.
  - The system will calculate a weighted average of the rates from both providers and use that as the exchange rate.
  - Weights will be configurable, allowing us to adjust based on changing conditions or business preferences.

    **Example Calculation**:
     - Bloomberg rate: 1.12 (weight: 0.7)
     - Yahoo rate: 1.10 (weight: 0.3)
     - Weighted average rate = (1.12 * 0.7) + (1.10 * 0.3) = **1.116**

2. **Store and Show Both Rates for Audit**:  
  - Both exchange rates (from Bloomberg and Yahoo) will be stored in the database.
  - For auditing purposes, these rates will be logged and accessible via the service API.
  - Users will be able to see both rates, along with the weighted average rate used by the system.
  - This ensures transparency and accountability, particularly useful for financial audits or regulatory reporting.

### Rationale

- **Weighted Average**: Balances the potential discrepancies between providers by considering the reliability of each. 
This approach smooths out outliers and provides a fair rate to users.
- **Storing Both Rates**: Adds transparency and auditability to the system. Users or administrators can verify 
the source rates and understand how the final exchange rate was derived.

### Alternatives Considered

1. **Primary Supplier Preference**:

- Designate one provider (e.g., Bloomberg) as the primary source and use its rate by default.
- **Reason for rejection**: This approach would ignore valuable data from the secondary provider (Yahoo), and in cases of temporary issues with the primary provider, it could lead to inaccuracies.

2. **Tolerance Threshold for Divergence**:

- Define a threshold for the acceptable divergence between rates (e.g., 0.5%). If the difference exceeds the threshold, an alternative strategy would be triggered (e.g., averaging the rates or falling back to manual review).
- **Reason for rejection**: While this approach helps detect significant discrepancies, it adds complexity to rate handling and might not address minor but consistent differences between the providers.

3. **Historical Trend Comparison**:

- Compare the rates from both providers against historical averages. If one rate significantly diverges from historical trends, it could be disregarded in favor of the other.
- **Reason for rejection**: This requires additional complexity in calculating and maintaining historical data, which might be overkill for real-time currency conversion.

4. **Manual Review**:

- In cases of significant divergence, the system could log the issue and notify an administrator for manual intervention.
- **Reason for rejection**: Manual review introduces latency in responding to rate requests and is not feasible in a real-time system.

## Consequences

- **Consistency**: The system provides a consistent exchange rate based on the weighted average, smoothing out discrepancies between providers.
- **Transparency**: Storing both rates ensures that the system can audit and justify the rates provided to users, allowing for better traceability.
- **Flexibility**: The weights can be adjusted over time as the reliability or performance of the providers changes.
