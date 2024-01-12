# 31. PostgreSQL Scale

Date: 2024-01-11

## Status

Accepted

## Context

We need to have a strategy for scale PostgreSQL.

## Decision

#### We have seen these options:

- **Hardware**
  - [Network-attached block storage](https://www.crunchydata.com/blog/an-overview-of-distributed-postgresql-architectures#network-attached-block-storage)
- **LoadBalancer**
  - [PGBouncer](https://www.pgbouncer.org/) - Lightweight connection pooler for PostgreSQL.
- **Replication**
  - [Read replicas](https://www.crunchydata.com/blog/an-overview-of-distributed-postgresql-architectures#read-replicas)
  - [Active-active](https://www.crunchydata.com/blog/an-overview-of-distributed-postgresql-architectures#active-active)
- **Partitioning**
- **Sharding (Transparent sharding)**
  - [Citus](https://www.citusdata.com/) - Citus is an open-source extension to Postgres that transforms Postgres into a distributed database.

#### Type of load:

- **OLTP** - Online Transaction Processing
- **OLAP** - Online Analytical Processing
- **HTAP** - Hybrid Transactional/Analytical Processing

## Consequences

We have a strategies for scale PostgreSQL.

### References

- [An Overview of Distributed PostgreSQL Architectures](https://www.crunchydata.com/blog/an-overview-of-distributed-postgresql-architectures)
