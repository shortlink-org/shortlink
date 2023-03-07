# 7. Optimization

Date: 2022-12-14

## Status

Accepted

## Context

Optimization is a key consideration in the design and implementation of microservices.
This is because microservices are often used in high-performance, scalable applications
that require the efficient use of resources.

## Decision

The change that we're proposing or have agreed to implement.

## Consequences

- [Network]
  - [gRPC]
    - [Using Protobuf FieldMask](https://netflixtechblog.com/practical-api-design-at-netflix-part-1-using-protobuf-fieldmask-35cfdc606518)
    - [Protobuf FieldMask for Mutation Operations](https://netflixtechblog.com/practical-api-design-at-netflix-part-2-protobuf-fieldmask-for-mutation-operations-2e75e1d230e4)
- [Serialization]
  - Use `github.com/segmentio/encoding/json` instead of `encoding/json` [proof](./proof/ADR-0007)
