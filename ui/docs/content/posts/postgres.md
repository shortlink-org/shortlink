---
title: "Postgres"
date: 2021-08-23T02:43:03+03:00
draft: false
categories:
    - Database
tags:
    - database
    - postgres
---

### Tuning params

1. `max_connections`
2. `shared_buffers` - use for cache [15..25%]
3. `effective_cache_size` [50%…75%]
4. `work_mem` - RAM for request (**WARNING**: count_session * work_mem). 

#### For write

1. `checkpoint_segments` [10-256]

### Links

1. [An Introduction to PostgreSQL Performance Tuning and Optimization](https://www.enterprisedb.com/postgres-tutorials/introduction-postgresql-performance-tuning-and-optimization)
