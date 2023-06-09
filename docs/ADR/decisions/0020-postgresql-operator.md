# 20. Research Summary and Decision for PostgreSQL Operators: Choosing CloudNativePG

Date: 2023-06-09

## Status

Accepted

## Context

As our system has continued to grow, the management of our PostgreSQL databases has become a critical aspect 
of our operations. We sought an automated, scalable solution that would streamline our operations and enable our team 
to focus on the development and delivery of our services. Our earlier solution, Zalando's Postgres Operator, 
served us well, but certain limitations were encountered, specifically its lack of support 
for ArgoCD, Prometheus out of the box, and its inability to generate URIs in secret.

The CloudNativePG Postgres Operator, developed by EnterpriseDB, offers a Kubernetes-native solution and effectively 
addresses the limitations we identified with the Zalando's Operator.

## Decision

After assessing various PostgreSQL operators, including the previously adopted Zalando Postgres Operator, PGO from Crunchy Data, Bitnami's PostgreSQL operator, and the new contender, CloudNativePG, we decided to adopt a new solution.

Despite the initial adoption of Zalando's Postgres Operator, the CloudNativePG Postgres Operator emerged as a superior option. This decision was motivated by its extensive feature set, which includes support for ArgoCD and Prometheus, as well as its ability to generate URIs in secrets.

We have decided to adopt the [CloudNativePG Postgres Operator](https://cloudnative-pg.io/).

Moreover, we will maintain our decision to create a unique schema and user for each service in our PostgreSQL databases. This strategy is designed to enhance data security, simplify access management, and provide more granular control.

## Consequences

The transition to the CloudNativePG Postgres Operator brings about several significant changes:

1. **Operational Efficiency**: The CloudNativePG Postgres Operator automates complex database operations within our Kubernetes environment, reducing the manual workload for our teams.

2. **Scalability**: The CloudNativePG operator, designed to handle the lifecycle of large database environments, is a scalable solution that can grow with our needs.

3. **Enhanced Features**: The CloudNativePG Postgres Operator provides out-of-the-box support for ArgoCD and Prometheus and is capable of generating URIs in secrets, features not provided by the Zalando Postgres Operator.

4. **Security and Control**: The CloudNativePG operator offers enhanced security measures and provides granular control over our databases. By continuing to implement unique schemas and users for each service, we can prevent unauthorized access and potential data breaches.

5. **Transition**: We have discontinued the use of the Zalando Postgres Operator following our transition to the CloudNativePG Postgres Operator. This switch required careful planning and execution to avoid service disruptions.

6. **Maintainability**: While the CloudNativePG Postgres Operator provides a robust solution, it's essential to remember that maintaining this system requires understanding and expertise in PostgreSQL and Kubernetes operations. It's a complex tool that might require additional resources for effective use.

#### Research Summary

We are currently using the [CloudNativePG Postgres Operator](https://cloudnative-pg.io/) to manage our PostgreSQL databases. This operator offers a feature set that aligns with our requirements, including:

- Effective management of users and roles.
- Seamless sharing of secrets across Kubernetes namespaces.
- Out-of-the-box support for ArgoCD and Prometheus.
- Capability to generate URIs in secrets.

Before settling on the CloudNativePG operator, we assessed alternative solutions:

- The [Zalando Postgres Operator](https://github.com/zalando/postgres-operator) was initially adopted, but its limitations, including
lack of support for ArgoCD, Prometheus, and inability to generate URIs in secret, led us to seek a more suitable option.
  - Example configuration for PGO: [link](./proof/ADR-0020/zalando/postgres-cluster.yaml)
- The [PGO, the Postgres Operator from Crunchy Data](https://access.crunchydata.com/documentation/postgres-operator/v5/)
  was evaluated but was found lacking in features necessary for our needs, such as control over users and roles and
  the ability to share secrets between Kubernetes namespaces.
  - Example configuration for PGO: [link](./proof/ADR-0020/postgres-cluster.yaml)
- **bitnami/postgresql** was previously in use, but due to the superior feature set of the Zalando Postgres Operator,
  we decided to discontinue its use.

The selection of the CloudNativePG Postgres Operator enhances our PostgreSQL management capabilities within our Kubernetes environment, providing us with a more complete, feature-rich solution.
