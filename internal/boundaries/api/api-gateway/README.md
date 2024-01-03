# API Gateway Service

> [!NOTE]
> This is the gateway service that serves as an interface for external clients, 
> leveraging multiple communication protocols and API designs.

### ADR

- [ADR-0001](./docs/ADR/decisions/0001-init.md) - Init project
- [ADR-0002](./docs/ADR/decisions/0002-split-api-gateway-service-into-separate-services-for-each-protocol.md) - Split API Gateway Service into Separate Services for Each Protocol
- [ADR-0003](./docs/ADR/decisions/0003-adding-websocket-protocol-support.md) - Adding WebSocket Protocol Support

### HTTP API

### Supported Protocols:

- **GraphQL API:** An excellent alternative to REST APIs for handling complex, interrelated data structures.
- **WebSocket API:** A protocol that enables two-way persistent communication channels over TCP connections.
    - [WebSocket API Docs](./gateways/ws/README.md)
- **gRPC-gateway:** Provides a way to call gRPC methods over HTTP, making your gRPC service more accessible to more client types.
    - Swagger [docs](./docs/server/v1/grpc_api.swagger.yaml)
- **CloudEvents (Optional):** An open-standard specification for defining event data in a common way, facilitating interoperability across services, platforms, and systems.
    - [CloudEvents Docs](https://cloudevents.io/)

### Architecture

We use the C4 model to describe architecture.

#### System context diagram

```plantuml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/common.puml
!include https://raw.githubusercontent.com/shortlink-org/shortlink/main/docs/c4/containers/preset/c1.puml

!include actors/customer.puml

!include boundaries/gateway.puml
!include boundaries/auth.puml
!include boundaries/link.puml

customer --> gatewayBoundary : uses
gatewayBoundary --> authBoundary : check auth
gatewayBoundary --> linkBoundary : create link
```
