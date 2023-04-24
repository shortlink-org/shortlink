## PostgreSQL

### Load Balancer

#### [pgbouncer](https://www.pgbouncer.org/)

```mermaid
graph LR
  A[Client] -->|1. Connect| B(PgBouncer)
  B -->|2. Check Load Balancer| LB[Load Balancer]
  LB -->|3. Choose Connection Pool| C[Connection Pool]
  C -->|4. Assign| D[Server Connection]
  D -->|5. Process| E[PostgreSQL Server]
  E -->|6. Response| D
  D -->|7. Return| C
  C -->|8. Forward| LB
  LB -->|9. Send to PgBouncer| B
  B -->|10. Send| A

  subgraph "PgBouncer Setup"
    B
    LB
    C
  end

  subgraph "Load Balancing Algorithms"
    L1[Round Robin]
    L2[Least Connection]
    L3[Custom Algorithm]
  end

  LB -.-> L1
  LB -.-> L2
  LB -.-> L3

  style A fill:#e3ebf3,stroke:#1a73e8
  style B fill:#e3ebf3,stroke:#1a73e8
  style C fill:#f7e3d3,stroke:#e18728
  style D fill:#e3ebf3,stroke:#1a73e8
  style E fill:#e3ebf3,stroke:#1a73e8
  style LB fill:#f7e3d3,stroke:#e18728
  style L1 fill:#d4f0c4,stroke:#50b840
  style L2 fill:#d4f0c4,stroke:#50b840
  style L3 fill:#d4f0c4,stroke:#50b840
```

### Backup

- ❌ [wal-g](https://github.com/wal-g/wal-g/blob/master/docs/PostgreSQL.md) is a tool for PostgreSQL backup and restore.
  But bitnami deprecated docker image, so if we want to use it, we need to build it by ourselves.
