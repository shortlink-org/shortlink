# 20. Research summary of PostgreSQL Operators

Date: 2023-06-09

## Status

Accepted

## Context

As our system scaled, the management of our PostgreSQL databases became increasingly complex. We identified the need 
for an automated, scalable solution that could manage this complexity and free up our development team's time. 
Our system operates within a Kubernetes environment, and we wanted a solution that could handle PostgreSQL management 
seamlessly within this environment. 

The Zalando Postgres Operator appeared to be an ideal choice with its Kubernetes-native design, 
but its initial inability to manage users and roles and to share secrets across Kubernetes namespaces 
presented challenges.

## Decision

We considered a range of PostgreSQL operators, including PGO from Crunchy Data, Bitnami's PostgreSQL operator, 
and Zalando's offering.

Despite the limitations we initially identified with the Zalando Postgres Operator, its robustness, scalability, 
and Kubernetes-native design made it stand out from the other options. 

We decided to adopt the [Zalando Postgres Operator](https://github.com/zalando/postgres-operator).

![Zalando Postgres Operator](https://postgres-operator.readthedocs.io/en/latest/diagrams/operator.png)

Moreover, we decided to create a unique schema and user for each service in our PostgreSQL databases. 
This strategy was designed to enhance data security, simplify access management, and provide more granular control.

## Consequences

Adopting the Zalando Postgres Operator as our database management system has brought about several significant changes:

1. **Operational Efficiency**: The Zalando Postgres Operator is an automated solution that reduces the manual workload 
  for our teams. By managing complex database operations natively in our Kubernetes environment, 
  we can concentrate on other critical aspects of our system.

2. **Scalability**: Zalando's operator is designed to handle the lifecycle of large database environments, 
  including provisioning, scaling, and configuration. This makes it a scalable solution that can grow with our needs and 
  handle the increased complexity that comes with scale.

3. **Security and Control**: The Zalando Postgres Operator has enhanced security measures in place and provides 
  granular control over our databases. By implementing unique schemas and users for each service, we can help prevent 
  unauthorized access and potential data breaches.

4. **Transition**: We have discontinued the use of **bitnami/postgresql** following our transition 
  to the Zalando Postgres Operator. This transition required careful planning and execution to avoid 
  any disruption to our services.

5. **Maintainability**: While the Zalando Postgres Operator provides a robust solution, it's important to remember 
  that maintaining this system requires understanding and expertise in PostgreSQL and Kubernetes operations. 
  It is a complex tool that may require additional resources for effective use.

#### Research Summary

We are currently using the [Zalando Postgres Operator](https://github.com/zalando/postgres-operator) 
to manage our PostgreSQL databases. This operator offers crucial features that match our requirements, including:

 - Effective management of users and roles.
 - Seamless sharing of secrets across Kubernetes namespaces.

Before settling on the Zalando operator, we assessed alternative solutions:

- The [PGO, the Postgres Operator from Crunchy Data](https://access.crunchydata.com/documentation/postgres-operator/v5/) 
  was evaluated but was found lacking in features necessary for our needs, such as control over users and roles and 
  the ability to share secrets between Kubernetes namespaces.
- **bitnami/postgresql** was previously in use, but due to the superior feature set of the Zalando Postgres Operator, 
  we decided to discontinue its use.

Our current choice of operator better aligns with our needs, enhancing our PostgreSQL management capabilities 
within our Kubernetes environment.
