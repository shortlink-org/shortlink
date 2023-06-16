## BFF: web

### Docs

- [API](./docs/api.yaml) - API documentation in Swagger format

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project

### Architecture

#### Component diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml
!include services/bff-web/service.puml
!include services/auth/ext.puml
!include services/billing/ext.puml

customer -> bff_web : Request 
bff_web -down-> auth : Proxy
bff_web -> billing : Proxy
```

### Use cases

- [UC-1](./usecases/api/README.md) - API implementation
