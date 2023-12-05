## BFF: web

> [!NOTE]
> This is a BFF service for web clients

### Docs

- [API](./docs/api.yaml) - API documentation in Swagger format

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-use-oapi-codegen.md) - Use `oapi-codegen` to generate API code

### Architecture

#### Component diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml
!include services/bff-web/service.puml
!include services/auth/ext.puml

title C1: BFF

System_Boundary(internalServices, "Internal Services") {
    !include services/billing/ext.puml
    !include services/link/ext.puml
    
    link -down-> permission : Permission
    billing -down-> permission : Permission
}

customer -> bff_web : Request 
bff_web -down-> auth : Proxy
bff_web -right-> internalServices : Proxy
```

### Use cases

- [UC-1](./usecases/api/README.md) - API implementation
