## HTTP client/server

### Middleware

| Name         | Description                                                                           |
|--------------|---------------------------------------------------------------------------------------|
| Auth         | This middleware authenticates the request.                                            |
| Logger       | This middleware logs the request.                                                     |
| Pprof Labels | This middleware adds route labels to pprof.                                           |
| NewMetrics   | This middleware creates a new prometheus metrics.                                     |
| RequestSize  | This middleware limits the request size.                                              |
| SingleFlight | This middleware shares the response.                                                  |
| Span         | This middleware set `trace_id` to the response header.                                 |
