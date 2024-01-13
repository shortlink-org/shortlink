## Delivery Boundary

> [!NOTE]
> The Delivery Boundary in a software system focuses on the logistics and management of delivering goods and services. 
> It encompasses functionalities related to geolocation for tracking deliveries, support systems for customer service, 
> and mechanisms for managing merchandise. This boundary is essential for ensuring efficient and accurate delivery 
> of products to customers, providing real-time updates on the status of deliveries, and offering customer support 
> to address any issues or inquiries related to the delivery process. It integrates various aspects of logistics, 
> order management, and customer feedback to enhance the overall delivery experience.

| Service     | Description         | Language/Framework | Docs                                              | Status                                                                                                                                                          |
|-------------|---------------------|--------------------|---------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| merch       | Merch store         | Go (Dapr)          | [docs](./internal/boundaries/delivery/merch/README.md)       | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-merch&revision=true)](https://argo.shortlink.best/applications/shortlink-merch)             |                                                                   
| support     | Support service     | PHP                | [docs](./internal/boundaries/delivery/support/README.md)     | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-support&revision=true)](https://argo.shortlink.best/applications/shortlink-support)         |                                                                 
| geolocation | Geolocation service | Go                 | [docs](./internal/boundaries/delivery/geolocation/README.md) | [![App Status](https://argo.shortlink.best/api/badge?name=shortlink-geolocation&revision=true)](https://argo.shortlink.best/applications/shortlink-geolocation) |

### Docs

- [GLOSSARY.md](./GLOSSARY.md) - Ubiquitous Language of the Delivery Boundary
