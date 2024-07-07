# 20. Research Summary and Decision for PostgreSQL Operators: Choosing Crunchy Postgres Operator

Date: 2023-06-09

## Status

Accepted

## Context

As our system has continued to evolve, the management of our PostgreSQL databases has remained a critical aspect of 
our operations. We sought an automated, scalable solution to streamline our operations and enable our team to focus 
on delivering our services. 

The Zalando Postgres Operator was our initial solution, but we encountered certain limitations, specifically its 
lack of support for ArgoCD, Prometheus, and its inability to generate URIs in secret.

We then considered the CloudNativePG Postgres Operator, developed by EnterpriseDB, which addressed some of 
the limitations we identified with Zalando's Operator. However, after further research and evaluation, 
we identified the Crunchy Postgres Operator as an even more fitting solution due to its ability 
to handle user management more effectively.

## Decision

We reassessed a range of PostgreSQL operators, including the previously adopted Zalando Postgres Operator, 
our temporary choice of CloudNativePG, Bitnami's PostgreSQL, and the Crunchy Postgres Operator from Crunchy Data. 
After careful evaluation, the Crunchy Postgres Operator emerged as the superior option due to its robust feature set, 
including more effective user management.

We have decided to adopt the [Crunchy Postgres Operator](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes).

We will continue our strategy of creating a unique schema and user for each service in our PostgreSQL databases. 
This approach is designed to enhance data security, simplify access management, and provide more granular control.

## Consequences

Our transition to the Crunchy Postgres Operator brings about several significant changes:

1. **Operational Efficiency**: The Crunchy Postgres Operator automates complex database operations within our 
  Kubernetes environment, reducing the manual workload for our teams.

2. **Scalability**: The Crunchy Postgres Operator, designed to handle the lifecycle of large database environments, 
  is a scalable solution that can grow with our needs.

3. **Effective User Management**: The Crunchy Postgres Operator outshines other options with its effective management 
  of users and roles.

4. **Security and Control**: The Crunchy Postgres Operator offers enhanced security measures and provides granular 
  control over our databases. By continuing to implement unique schemas and users for each service, 
  we can prevent unauthorized access and potential data breaches.

5. **Transition**: We have discontinued the use of the Zalando Postgres Operator and CloudNativePG Operator following 
  our transition to the Crunchy Postgres Operator. This transition required careful planning and execution to avoid 
  any disruption to our services.

6. **Maintainability**: While the Crunchy Postgres Operator provides a robust solution, it's important to remember 
  that maintaining this system requires understanding and expertise in PostgreSQL and Kubernetes operations. 
  It is a complex tool that may require additional resources for effective use.

#### Research Summary

We are currently using the [Crunchy Postgres Operator](https://www.crunchydata.com/products/crunchy-postgresql-for-kubernetes) 
to manage our PostgreSQL databases. This operator offers a feature set that aligns with our requirements, including:

> Example configuration for PGO: [link](./proof/ADR-0020/postgres-cluster.yaml)

- Effective management of users and roles.
- Seamless sharing of secrets across Kubernetes namespaces.
- Capability to generate URIs in secrets.

Before settling on the Crunchy operator, we assessed alternative solutions:

- The [Zalando Postgres Operator](https://github.com/zalando/postgres-operator) was our initial choice, but its limitations, 
  including the lack of support for ArgoCD, Prometheus, and inability to generate URIs in secret, led us to look for other options.

> Example configuration for Zalando Postgres Operator: [link](./proof/ADR-0020/zalando/postgres-cluster.yaml)

- The [CloudNativePG Postgres Operator](https://cloudnative-pg.io/) was briefly adopted, but we found 
  the Crunchy Postgres Operator's features and user management capabilities more beneficial for our needs.

> Example configuration for Zalando Postgres Operator: [link](./proof/ADR-0020/cloudnative-pg/postgres-cluster.yaml)

- **bitnami/postgresql** was previously in use, but due to the superior feature set of the Crunchy Postgres Operator, 
  we decided to discontinue its use.

The selection of the Crunchy Postgres Operator enhances our PostgreSQL management capabilities within 
our Kubernetes environment, providing us with a more complete, feature-rich solution.

### Alternatives

- [tembo](https://tembo.io/) - Postgres Operator for Kubernetes
