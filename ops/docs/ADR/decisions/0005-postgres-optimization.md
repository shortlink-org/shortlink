# 5. Prometheus stack

Date: 2022-12-10

## Status

Accepted

## Context

We want to have a fast and highly available PostgreSQL instance.

#### Links

1. [An Introduction to PostgreSQL Performance Tuning and Optimization](https://www.enterprisedb.com/postgres-tutorials/introduction-postgresql-performance-tuning-and-optimization)

## Decision

### Tool

1. **LoadBalancer**
    - [pgBouncer](https://www.pgbouncer.org/)
2. **Optimization**
    - For optimizing PostgreSQL configurations, we will use the tool available at [pgtune.leopard.in.ua](https://pgtune.leopard.in.ua/#/). 
      This tool generates configuration files that provide the optimal settings based on the system's hardware and expected workload, improving database performance.
3. **LoadTest**
    - [Percona-Lab/sysbench-tpcc](https://github.com/Percona-Lab/sysbench-tpcc)

### Tuning params

1. `max_connections`
2. `shared_buffers` - use for cache [15..25%]
3. `effective_cache_size` [50%…75%]
4. `work_mem` - RAM for request (**WARNING**: count_session * work_mem).

#### For one database

```bash
# Set timeout (in seconds) before a query cancels:
ALTER DATABASE mydatabase SET statement_timeout = '60s';
```

#### For write

1. `checkpoint_segments` [10-256]

## Consequences

1. We have a fast and highly available PostgreSQL instance.
2. We can use `pgBouncer` for load balancing.
3. We can use `pgTune` for tuning PostgreSQL.
