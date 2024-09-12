# 4. API Design

Date: 2024-09-12

## Status

Accepted

## Context

The Currency Service API must handle both commands (e.g., updating exchange rates) 
and queries (e.g., retrieving real-time or historical exchange rates) efficiently. 
To ensure scalability, performance, and separation of concerns, we will adopt the **CQRS (Command Query Responsibility Segregation)** pattern. 
This will help optimize the system by separating the command operations (which modify data) from query operations (which read data).

## Decision


We will implement the API using the CQRS pattern to divide the system into two distinct parts:

- **Command side**: Responsible for handling operations that change the state of the system (e.g., updating exchange rates).
- **Query side**: Responsible for handling read operations (e.g., fetching real-time or historical exchange rates).

The API specification will be documented using **Swagger** and can be found [here](../../api/openapi.yaml).

### Command Side

- The command side will handle incoming requests to update or modify exchange rates.
- These requests will be processed asynchronously to ensure smooth handling of high volumes.
- Updated exchange rates will be written to both the **Rate Database** and the **Cache Store**.

### Query Side

- The query side will retrieve data from the **Cache Store** first, for performance reasons, and fall back to the **Rate Database** if necessary.
- If the requested data is not available in the cache, it will be fetched from external sources (e.g., Bloomberg, Yahoo) and cached for future requests.

## Consequences

- **Separation of Concerns**: Clear separation between read and write operations enhances system maintainability.
- **Scalability**: The system will scale better, especially under high load, by handling queries and commands independently.
- **Performance**: The use of caching and optimized read pathways will ensure that query performance is fast and efficient.
- **Asynchronous Processing**: Command operations will be processed asynchronously, avoiding delays for users during data updates.
