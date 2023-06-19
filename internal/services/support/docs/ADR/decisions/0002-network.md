# 2. Network Setup for Support Service

Date: 2023-06-19

## Status

Accepted

## Context

Our "Support Service" needs a high-performance network setup to handle customer requests effectively. 
This service needs to handle high traffic volumes while providing responsive and reliable support to our users.

## Decision

After considering several options, we have decided to use Nginx as a reverse proxy due to its high performance, stability, 
and efficient handling of both static and dynamic content. We have configured Nginx to listen 
on port 8080 for both IPv4 and IPv6 traffic.

```nginx configuration
server {
    listen      8080;
    listen [::]:8080;
    server_name _;
}
```

For PHP script execution, we've chosen PHP-FPM due to its capability to handle heavy loads and its compatibility 
with OPCache, which we are using for performance optimization. We have enabled the PHP-FPM status page 
for easy monitoring and management of our PHP service.

```php-fpm configuration
# Enable the PHP-FPM status page
pm.status_path = /status
pm.status_listen = 127.0.0.1:9001
``````

## Consequences

Our setup using Nginx, PHP-FPM, and OPCache will offer enhanced performance, effective traffic management,
and faster PHP processing. This setup requires careful configuration and monitoring to maintain optimal performance.
While this introduces additional tools that the team will need to familiarize themselves with,
the benefits of improved performance and reliability outweigh the initial learning curve.

### Implementation Strategy

The implementation will involve setting up and configuring the Nginx server, installing and setting up PHP-FPM with OPCache, 
and ensuring the correct routing of requests. The development team will be responsible for this setup and ongoing management.

```mermaid
graph LR;
  A[Client Request] --> B[Nginx-Proxy];
  B --> C[PHP-FPM & OPCache];
  C --> D[PHP Processing];
  D --> E[Server Response];
```

This flowchart visualizes a client request journey from Nginx proxy to PHP-FPM with OPCache for processing, 
and finally a server response is sent back to the client.
