# 2. Store Provider Selection

Date: 2023-05-21

## Status

Accepted

## Context

Our ShortLink service requires a reliable and efficient data store for managing links. The choice of data store influences factors like the support for batch operations, filtering, and the overall scalability of the system. A variety of data stores are available each with its own benefits and trade-offs.

## Decision

After evaluating the features and scalability of different data stores, we have decided to support RAM, MongoDB, Postgres, Redis, LevelDB, Badger, SQLite, and DGraph.

Here is the support and scalability summary of each:

| Name     | Support | Scale |
|----------|---------|-------|
| RAM      | ✅       | ❌     |
| MongoDB  | ✅       | ✅     |
| Postgres | ✅       | ✅     |
| Redis    | ❌       | ✅     |
| LevelDB  | ❌       | ❌     |
| Badger   | ❌       | ❌     |
| SQLite   | ❌       | ❌     |
| DGraph   | ❌       | ✅     |

However, the decision is also influenced by our past experiences. We have decided to drop support for MySQL, scylla, and cassandra due to various reasons, as reflected in our changelog:

- [19.09.2022] Drop support database: MySQL
- [04.08.2021] Drop support database: scylla, cassandra

## Consequences

By choosing to support a range of data stores, we allow for a wider selection for users of the ShortLink service depending on their specific needs and contexts. This increases the flexibility and applicability of our service.

However, maintaining support and ensuring optimal integration with all these data stores will add to the complexity of our system. It increases the workload for our development team and the potential points of failure in our service.

Moreover, we will need to monitor and reevaluate our choice of data stores continually, as we've done in the past with MySQL, scylla, and cassandra, to ensure they continue to meet our service requirements and provide the best experience to our users.
