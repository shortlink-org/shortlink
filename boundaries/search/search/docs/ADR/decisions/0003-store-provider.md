# 3. Store Provider Selection

Date: 2024-02-17

## Status

Accepted

## Context

Selecting a search database is critical for supporting our application's full-text search, real-time indexing, 
and scalability requirements. Elasticsearch is well-regarded for these capabilities. 
However, the emergence of vector databases, such as pgvector, offers new possibilities for similarity searches.

## Decision

We have chosen Elasticsearch as our primary search database due to its proven performance, scalability, 
and full-text search capabilities. To address future needs for similarity searches, we plan to explore integration 
with vector databases like pgvector, leveraging the best of both worlds.

## Consequences

+ **Elasticsearch**: Ensures robust full-text search, scalability, and real-time processing.
+ **Vector Database (pgvector)**: Offers potential for advanced similarity searches, enhancing search functionalities.

### References

Selecting Elasticsearch addresses our current requirements effectively while keeping the door open 
for future enhancements with vector databases to meet evolving search scenarios.

- [Elasticsearch](https://www.elastic.co/elasticsearch/)
- [paradedb](https://www.paradedb.com/) - A vector database
- [pgvector](https://github.com/pgvector/pgvector) - A vector database extension for PostgreSQL
