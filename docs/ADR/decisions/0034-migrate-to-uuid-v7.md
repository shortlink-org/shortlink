# 34. Migrate to UUID v7

Date: 2024-07-20

## Status

Accepted

## Context

UUID v4 has been used for primary keys in our databases. However, UUID v7 offers several advantages over UUID v4, including improved sortability and time-based generation, which can enhance performance and scalability in distributed systems.

## Decision

We will migrate our primary keys from UUID v4 to UUID v7. This decision involves updating the database schema and modifying the codebase to generate and handle UUID v7 keys.

## Consequences

### Positive

- **Improved Sortability**: UUID v7 allows for more efficient indexing and sorting operations.
- **Time-based Generation**: Facilitates chronological ordering and improves caching strategies.
- **Scalability**: Enhanced support for distributed systems and clustering.

### Negative

- **Migration Effort**: Requires significant changes to the database schema and application logic.
- **Potential Downtime**: There may be some downtime required during the migration process.

## Implementation

1. **Schema Migration**: Update the database schema to support UUID v7 as primary keys.
2. **Code Changes**: Modify the codebase to generate and handle UUID v7 keys.
3. **Testing**: Thoroughly test the changes in a staging environment before deploying to production.
4. **Deployment**: Carefully plan and execute the migration to minimize downtime and impact on users.

## References

- [UUID v7 Specification](https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format)
