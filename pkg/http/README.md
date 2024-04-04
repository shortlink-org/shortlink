## HTTP client/server

### Middleware

| Name                                      | Description                                            |
|-------------------------------------------|--------------------------------------------------------|
| [Auth](./middleware/auth)                 | This middleware authenticates the request.             |
| [Logger](./middleware/logger)             | This middleware logs the request.                      |
| [Metrics](./middleware/metrics)           | This middleware creates a new prometheus metrics.      |
| [Pprof Labels](./middleware/pprof_labels) | This middleware adds route labels to pprof.            |
| [RequestSize](./middleware/request_size)  | This middleware limits the request size.               |
| [SingleFlight](./middleware/singleflight) | This middleware shares the response.                   |
| [Span](./middleware/span)                 | This middleware set `trace_id` to the response header. |
