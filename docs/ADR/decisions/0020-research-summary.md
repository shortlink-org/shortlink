# 20. research summary

Date: 2023-03-30

## Status

Accepted

## Context

We want save results of research in one place.

## Decision

We'll use this document to save results of research.

## Consequences

### DataBase

#### PostgreSQL

<details>

<summary>Click to expand</summary>
<p>

We try to use PostgreSQL Operator:

- [PGO, the Postgres Operator from Crunchy Data](https://access.crunchydata.com/documentation/postgres-operator/v5/)
  - Don't control users and roles
  - Don't share secrets between k8s-namespaces
- [zalando/postgres-operator](https://github.com/zalando/postgres-operator)
  - Don't control users and roles
  - Don't share secrets between k8s-namespaces

In summary, we can't use any of these operators. At the moment we're using **bitnami/postgresql**

</p>
</details>
