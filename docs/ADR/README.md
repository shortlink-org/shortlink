# ADR

### How use

```shell
export ADR_TEMPLATE=${PWD}/docs/ADR/template/template.md
adr new Implement as Unix shell scripts
```

### Docs

- [How to install](https://github.com/npryce/adr-tools/blob/master/INSTALL.md)

### Update

For update ADR we use GIT, so we can get date each updated and we use git message
for information team about cases for updated ADR.

### Architecture Decision Log (ADL):

- **Organization**
  - [ADR-0001](./decisions/0001-record-architecture-decisions.md) - Record architecture decisions
  - [ADR-0006](./decisions/0006-codeowner.md) - Codeowner
  - [ADR-0011](./decisions/0011-application-architecture-documentation.md) - Application architecture documentation
  - [ADR-0012](./decisions/0012-use-monorepository.md) - Use monorepository
  - [ADR-0021](./decisions/0021-microservice-structure.md) - Microservice structure
  - [ADR-0024](./decisions/0024-selecting-a-development-tool-for-remote-machine.md) - Selecting a Development Tool for Remote Machine
- **Naming**
  - [ADR-0002](./decisions/0002-implement-as-event-naming.md) - Implement as event naming
  - [ADR-0009](./decisions/0009-naming-spans-and-attributes.md) - Naming spans and attributes
  - [ADR-0023](./decisions/0023-naming-prometheus-metrics.md) - Prometheus Metrics Naming
  - [ADR-0010](./decisions/0010-logger-format.md) - Logger format
- **Observability**
  - [ADR-0003](./decisions/0003-observability-health-check.md) - Observability health check
  - [ADR-0014](./decisions/0014-observability.md) - Standardizing Observability Tools
- **Network**
  - [ADR-0007](./decisions/0007-optimization-network.md) - Optimization network
  - **Proposed**: [ADR-0018](./decisions/0018-kubernetes-network.md) - Kubernetes network
- **DataBase**
  - [ADR-0004](./decisions/0004-use-uuid-as-primary-keys.md) - Use UUID as primary keys
  - [ADR-0005](./decisions/0005-postgres-optimization.md) - Postgres optimization
  - [ADR-0020](./decisions/0020-postgresql-operator.md) - Research Summary and Decision for PostgreSQL Operators: Choosing Crunchy Postgres Operator
  - [ADR-0026](./decisions/0026-pattern-database-per-service.md) - Pattern: database per service
- **Service**
  - [ADR-0008](./decisions/0008-product-metrics-by-services.md) - Product metrics by services
  - [ADR-0015](./decisions/0015-authentication.md) - Authentication
- **Security**
  - [ADR-0013](./decisions/0013-security.md) - Enhancing Security Measures
  - [ADR-0025](./decisions/0025-configuration.md) - Application Configuration Redesign
- **Golang**
  - [ADR-0017](./decisions/0017-profile-guided-optimization.md) - Profile-guided optimization
- **Front-end**
  - [ADR-0019](./decisions/0019-front-end-testing.md) - Front-end testings
- **Kubernetes**
  - [ADR-0022](./decisions/0022-kubernetes.md) - Kubernetes
  - [ADR-0016](./decisions/0016-lifecycle-deploy.md) - Lifecycle deploy
