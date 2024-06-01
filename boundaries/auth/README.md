## Auth boundary

> [!NOTE]
> The Auth boundary in software architecture is dedicated to handling authentication and authorization processes. 
> Authentication is about verifying user identities, ensuring that users are who they claim to be. 
> Authorization, on the other hand, determines what authenticated users are allowed to do within the system, 
> often managing permissions and access levels. This boundary is vital for security, controlling access to different 
> parts of the application based on user roles, permissions, and policies. It ensures that sensitive information and 
> critical functionalities are accessible only to authorized users, thereby safeguarding the integrity and 
> confidentiality of the system.

| Service | Description      | Language/Framework | Docs                     | Status                                                                                                                                                      |
|---------|------------------|--------------------|--------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------|
| auth    | Auth service     | Go                 | [docs](./auth/README.md) | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-auth-auth&revision=true)](https://argo.shortlink.best/applications/shortlink-auth-auth) |

### Third-party Service

| Service    | Description                   | Language/Framework | Docs                                    | Status                                                                                                                        |
|------------|-------------------------------|--------------------|-----------------------------------------|-------------------------------------------------------------------------------------------------------------------------------|
| ory/kratos | User management service       | Go                 | [docs](https://www.ory.sh/kratos/docs/) | [![App Status](https://argo.shortlink.best/api/badge?name=auth&revision=true)](https://argo.shortlink.best/applications/auth) |          
| ory/hydra  | OAuth 2.0 Provider            | Go                 | [docs](https://www.ory.sh/keto/docs/)   | [![App Status](https://argo.shortlink.best/api/badge?name=auth&revision=true)](https://argo.shortlink.best/applications/auth) |          
| spiceDB    | Permission management service | Go                 | [docs](https://authzed.com/docs)        | [![App Status](https://argo.shortlink.best/api/badge?name=auth&revision=true)](https://argo.shortlink.best/applications/auth) |

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the Auth Boundary
