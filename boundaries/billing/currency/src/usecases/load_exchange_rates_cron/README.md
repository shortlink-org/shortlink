# Use Case: Scheduled Updates of Exchange Rate Data via Cron Job

> **Note**  
> This use case covers the scheduled update process where a cron job runs at predefined intervals to refresh exchange rate data from Bloomberg and Yahoo APIs, ensuring that the data is up to date.

## Flow

1. **Scheduled Trigger**:  
   A cron job is set up to run at regular intervals (e.g., hourly). This job triggers the **Currency Service** to refresh exchange rate data from external providers. The purpose is to keep the system's cache and database populated with the latest rates, reducing the need to make API calls for every user request.

2. **Fetch Exchange Rates**:  
   The **Currency Service** executes the `load_exchange_rates` use case to fetch exchange rates from both **Bloomberg** and **Yahoo** APIs.  
   By using this [Load Exchange Rate Data from Subscriptions](#use-case-load-exchange-rate-data-from-subscriptions) use case, it ensures that the system maintains an up-to-date set of exchange rates.

3. **Handle API Responses**:
   - **Success**: If the API calls to **Bloomberg** and **Yahoo** are successful, the service retrieves the updated exchange rate data.
   - **Failure and Fallback**: If one API fails or is unavailable, the service will attempt to use the alternative provider. Additionally, it implements retries with exponential backoff to handle transient failures and avoid overwhelming the external providers.

4. **Store Rates in Cache and Database**:  
   Once the service successfully retrieves the updated rates, it stores them in the **Cache** for quick access and in the **Rate Database** for long-term storage and auditing.

5. **Logging and Monitoring**:  
   Throughout the process, the service logs the activities and any errors encountered, such as API failures or rate limit exceedances. This logging is crucial for monitoring the health of the system and for debugging issues when they occur.

6. **Completion**:  
   Once the data has been fetched, processed, and stored, the cron job execution is considered complete. The service waits until the next scheduled interval to perform another update.

## Sequence Diagram

### Cron Job Interaction Diagram

```plantuml
@startuml
skinparam actorBackgroundColor #ADD8E6
skinparam participantBackgroundColor #90EE90
skinparam noteBackgroundColor #FFFFE0
skinparam sequenceLifeLineBackgroundColor #FF6347

actor CronJob as cron
participant CurrencyService as service
participant CacheStore as cache
participant BloombergAPI as bloomberg
participant YahooAPI as yahoo
participant RateDatabase as db

cron -> service: Trigger scheduled rate refresh
service -> bloomberg: Fetch latest rates from Bloomberg
bloomberg --> service: Return updated rates
service -> yahoo: Fetch latest rates from Yahoo
yahoo --> service: Return updated rates
service -> cache: Update cache with new rates
service -> db: Update database with new rates

note over cron, service
1. Cron job triggers periodic refresh.
2. Service fetches latest rates from Bloomberg and Yahoo.
3. Updates cache and database with new rates.
end note
@enduml
```

## Error Handling

### 1. API Failure

- **Scenario**: One of the external APIs (Bloomberg or Yahoo) fails to return data or is unavailable.
- **Handling**:
   - The service retries the request with exponential backoff.
   - If the API continues to fail, the system falls back to the alternative provider.
   - The failure is logged, and an alert may be triggered depending on the severity and duration of the outage.

### 2. Rate Limit Exceeded

- **Scenario**: The number of requests sent to the external API exceeds its rate limits.
- **Handling**:
   - The service tracks API usage and implements rate limiting to prevent further requests from exceeding the limits.
   - Retries are scheduled based on the API's rate limit reset time.
   - Errors are logged for monitoring and further analysis.

### 3. Cache or Database Unavailable

- **Scenario**: The cache or database is unavailable when the service attempts to store the updated exchange rates.
- **Handling**:
   - If the **Cache** is unavailable, the service will skip updating the cache and log the error.
   - If the **Rate Database** is unavailable, the system attempts to retry at a later time.
   - All failures are logged, and alerts may be triggered if necessary.

### 4. Invalid or Inconsistent Data

- **Scenario**: The data retrieved from the external provider is invalid or inconsistent.
- **Handling**:
   - The service discards invalid data and logs the issue.
   - It attempts to fetch the data again during the next scheduled cron job.
   - The system will continue serving cached or previously stored data to clients until valid data is retrieved.
