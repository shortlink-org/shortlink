## HTTP client/server

### Middleware

| Name         | Description                                       |
|--------------|---------------------------------------------------|
| Logger       | This middleware logs the request.                 |
| NewMetrics   | This middleware creates a new prometheus metrics. |
| RequestSize  | This middleware limits the request size.          |
| SingleFlight | This middleware shares the response.              |
