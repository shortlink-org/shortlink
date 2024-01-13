## API Boundary

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

| Service         | Description                    | Language/Framework | Docs                                                          | Status                                                                                                                                                                  |
|-----------------|--------------------------------|--------------------|---------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| api-cloudevents | Internal GateWay               | Go                 | [docs](./internal/boundaries/api/api-gateway/README.md)             | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-cloudevents&revision=true)](https://argo.shortlink.best/applications/shortlink-api-cloudevents) |
| api-graphql     | Internal GateWay               | Go                 | [docs](./internal/boundaries/api/api-gateway/README.md)             | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-graphql&revision=true)](https://argo.shortlink.best/applications/shortlink-api-graphql)         |
| api-grpc-web    | Internal GateWay               | Go                 | [docs](./internal/boundaries/api/api-gateway/README.md)             | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-grpc-web&revision=true)](https://argo.shortlink.best/applications/shortlink-api-grpc-web)       |
| api-ws          | Websocket service              | Go                 | [docs](./internal/boundaries/api/api-gateway/gateways/ws/README.md) | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-api-ws&revision=true)](https://argo.shortlink.best/applications/shortlink-api-ws)                   |
| bff-web         | BFF for web                    | Go                 | [docs](./internal/boundaries/api/bff-web/README.md)                 | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-bff-web&revision=true)](https://argo.shortlink.best/applications/shortlink-bff-web)                 |

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the API Boundary
- [README.md](./docs/ADR/README.md) - Architecture Decision Records

