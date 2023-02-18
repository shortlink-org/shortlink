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

### ADR

- **Organization**
  - [ADR-0001](./decisions/0001-record-architecture-decisions.md) - Record architecture decisions
  - [ADR-0006](./decisions/0006-codeowner.md) - Codeowner
  - [ADR-0011](./decisions/0011-application-architecture-documentation.md) - Application architecture documentation
  - [ADR-0012](./decisions/0012-use-monorepository.md) - Use monorepository
- **Lifecycle**
  - [ADR-0016](./decisions/0016-lifecycle-deploy.md) - Lifecycle deploy
- **Naming**
  - [ADR-0002](./decisions/0002-implement-as-event-naming.md) - Implement as event naming
  - [ADR-0009](./decisions/0009-naming-spans-and-attributes.md) - Naming spans and attributes
  - [ADR-0010](./decisions/0010-logger-format.md) - Logger format
- **Observability**
  - [ADR-0014](./decisions/0014-observability.md) - Observability common
  - [ADR-0003](./decisions/0003-observability-health-check.md) - Observability health check
- **Network**
  - [ADR-0007](./decisions/0007-optimization-network.md) - Optimization network
- **DataBase**
  - [ADR-0004](./decisions/0004-use-uuid-as-primary-keys.md) - Use UUID as primary keys
  - [ADR-0005](./decisions/0005-postgres-optimization.md) - Postgres optimization
- **Service**
  - [ADR-0008](./decisions/0008-product-metrics-by-services.md) - Product metrics by services
- **Security**
  - [ADR-0013](./decisions/0013-security.md) - Security
- **Golang**
  - [ADR-0017](./decisions/0017-profile-guided-optimization.md) - Profile-guided optimization
- **Third-party**
  - [ADR-0015](./decisions/0015-authentication.md) - Authentication
