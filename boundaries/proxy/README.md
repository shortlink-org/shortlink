### Proxy

<img width='200' height='200' src="./docs/public/logo.svg">

> [!NOTE]
> Proxy service for redirecting to the original URL

##### Require

```
pnpm add grpc-tools grpc_tools_node_protoc_ts --global
```

##### Security: Permissions API (Zero-Trust)

Proxy Service использует Node.js Permissions API для обеспечения Zero-Trust безопасности.

- **Production**: Автоматически применяются ограничения через `permissions.json`
- **Development**: Используйте `pnpm start:permissive` для разработки без ограничений

##### Observability

- Все события, публикуемые в AMQP, включают заголовок `traceparent` (и при наличии `tracestate`) в message headers, что позволяет downstream-сервисам связывать трейс при обработке сообщения через OpenTelemetry `propagation.extract`.
