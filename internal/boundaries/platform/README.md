## Platform Boundary

> [!NOTE]
> The Platform Boundary in a software system encompasses the core infrastructure and tools that provide 
> the foundational support for application development and operation. It includes aspects like developer portals 
> for centralized tool management, containerization for efficient application deployment, and logging systems 
> for monitoring and troubleshooting. This boundary is crucial for maintaining the overall health and efficiency of 
> the software ecosystem, offering the necessary tools and services for seamless development, deployment, 
> and operational processes.

| Service   | Description                                                   | Language/Framework | Docs                                         | Status                                                                                                                                                |
|-----------|---------------------------------------------------------------|--------------------|----------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| csi       | CSI example                                                   | Go                 | [docs](./internal/boundaries/platform/csi/README.md)    | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-csi&revision=true)](https://argo.shortlink.best/applications/shortlink-csi)       |                                                                     
| logger    | Logger service                                                | Go                 | [docs](./internal/boundaries/platform/logger/README.md) | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-logger&revision=true)](https://argo.shortlink.best/applications/shortlink-logger) |                                                                  
| shortctl  | ShortLink CLI                                                 | Go                 | [docs](./internal/boundaries/platform/cli/README.md)    |                                                                                                                                                       |                                                                   

### Third-party Service

| Service   | Description                                                   | Language/Framework | Docs                                         | Status                                                                                                                                                |
|-----------|---------------------------------------------------------------|--------------------|----------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|
| backstage | Backstage is an open platform for building developer portals. | TypeScript         | [docs](https://backstage.io/docs/)           | [![App Status](https://argo.shortlink.best/api/badge?name=backstage&revision=true)](https://argo.shortlink.best/applications/backstage)               |    

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the Platform Boundary
