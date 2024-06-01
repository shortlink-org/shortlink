## Link Boundary

> [!NOTE]
> The Link Boundary in software systems primarily deals with the management and functionality of URLs, 
> especially in the context of shortening and tracking them. This includes creating shortlinks for easier sharing and 
> readability, redirecting users from the shortlink to the original URL, and analyzing link usage for insights. 
> Features like QR code generation for easy access and custom domains for branding are also part of this boundary. 
> It's essential for optimizing user experience in sharing and accessing web resources, providing valuable data through 
> link analytics, and enhancing brand visibility with custom domains.

| Service          | Description                                | Language/Framework | Docs                                 | Status                                                                                                                                                              |
|------------------|--------------------------------------------|--------------------|--------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| bff-link         | BFF for link boundary                      | Go                 | [docs](./bff/README.md)              | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-link-bff&revision=true)](https://argo.shortlink.best/applications/shortlink-link-bff)           |
| link             | Link manage service                        | Go                 | [docs](./link/README.md)             | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-link-link&revision=true)](https://argo.shortlink.best/applications/shortlink-link-link)         |                                                                    
| proxy            | Proxy service for redirect to original URL | TypeScript         | [docs](./proxy/README.md)            | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-link-proxy&revision=true)](https://argo.shortlink.best/applications/shortlink-link-proxy)       |                                                                   
| metadata         | Parser site by API                         | Go                 | [docs](./metadata/README.md)         | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-link-metadata&revision=true)](https://argo.shortlink.best/applications/shortlink-link-metadata) |                                                                
| chrome-extension | Chrome extension                           | JavaScript         | [docs](./chrome-extension/README.md) | -                                                                                                                                                                   |
| supplies         | Supplies service                           | Go                 | [docs](./supplies/README.md)         | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-link-supplies&revision=true)](https://argo.shortlink.best/applications/shortlink-link-supplies) |
| mobile           | Mobile UI (expo)                           | JS/React           | [docs](./mobile/README.md)           | -                                                                                                                                                                   |

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the Link Boundary
- [README.md](./docs/ADR/README.md) - Architecture Decision Records
