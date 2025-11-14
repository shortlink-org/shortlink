# ADR

### How to use

```shell
export ADR_TEMPLATE=${PWD}/docs/ADR/template/template.md
adr new Implement as Unix shell scripts
```

### Docs

- [How to install](https://github.com/npryce/adr-tools/blob/master/INSTALL.md)

### Update

For update ADR we use GIT, so we can get date each updated, and we use a git message
for an information team about cases for updated ADR.

### Architecture Decision Log (ADL):

- [ADR-0001](./decisions/0001-connect-interceptors.md) - Connect Interceptors для cross-cutting concerns
- [ADR-0002](./decisions/0002-clean-architecture.md) - Clean Architecture с разделением на слои
- [ADR-0003](./decisions/0003-use-case-pipeline.md) - Use Case Pipeline для cross-cutting concerns
- [ADR-0004](./decisions/0004-anti-corruption-layer.md) - Anti-Corruption Layer для внешних сервисов
- [ADR-0005](./decisions/0005-event-driven-architecture.md) - Event-Driven Architecture
- [ADR-0006](./decisions/0006-permissions-api-zero-trust.md) - Permissions API для Zero-Trust безопасности
- [ADR-0007](./decisions/0007-redis-caching.md) - Redis кэширование для оптимизации производительности
