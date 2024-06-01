## Notification Boundary

> [!NOTE]
> The Notification Boundary in a software system is dedicated to managing and delivering various types of alerts 
> and messages to users. It covers a range of notification methods, including push notifications for immediate attention, 
> email and SMS notifications for broader communication, and in-app notifications for app-specific information. 
> This boundary also includes user preferences for notification management, ensuring that users receive relevant alerts 
> in their preferred manner. The effective implementation of this boundary is crucial for keeping users informed and engaged, 
> enhancing user experience, and ensuring timely communication of important information or updates.

| Service    | Description                          | Language/Framework | Docs                           | Status                                                                                                                                                        |
|------------|--------------------------------------|--------------------|--------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|
| bot        | Telegram bot                         | JAVA               | [docs](./bot/README.md)        |                                                                                                                                                               |                                                                    
| newsletter | Newsletter service                   | Rust               | [docs](./newsletter/README.md) | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-newsletter&revision=true)](https://argo.shortlink.best/applications/shortlink-newsletter) |                                                              
| notify     | Send notify to smtp, slack, telegram | Go                 | [docs](./notify/README.md)     | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-notify&revision=true)](https://argo.shortlink.best/applications/shortlink-notify)         |                                                                  

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the Notification Boundary
