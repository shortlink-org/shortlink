## API Boundary

> [!WARNING]
> These services **DEPRECATED** and will be removed in the future.

> [!NOTE]
> The API-Gateway boundary in a software architecture is a critical component, serving as the primary 
> interface between the external world and the internal systems or services of an application. 
> It typically includes both an API (Application Programming Interface) and a BFF (Backend for Frontend). 
> The API part of this boundary is responsible for defining how external entities interact with the system, 
> specifying the protocols and methods for data exchange. The BFF, on the other hand, 
> is tailored to specific user interface requirements, handling the server-side logic necessary for those interfaces. 
> This boundary is essential for managing requests, ensuring secure and efficient communication, 
> and providing a clear separation between external interfaces and internal service logic.

### Services

| Service         | Description       | Language/Framework | Docs                                                       |
|-----------------|-------------------|--------------------|------------------------------------------------------------|
| api-cloudevents | Internal GateWay  | Go                 | [docs](./boundaries/api/api-gateway/README.md)             |
| api-graphql     | Internal GateWay  | Go                 | [docs](./boundaries/api/api-gateway/README.md)             |
| api-grpc-web    | Internal GateWay  | Go                 | [docs](./boundaries/api/api-gateway/README.md)             |
| api-ws          | Websocket service | Go                 | [docs](./boundaries/api/api-gateway/gateways/ws/README.md) |

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the API Boundary
- [README.md](./docs/ADR/README.md) - Architecture Decision Records

