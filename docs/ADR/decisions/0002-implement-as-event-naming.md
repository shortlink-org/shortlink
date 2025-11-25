# 2. Implement as event naming

Date: 2022-08-24

## Status

Accepted

## Context

The Shortlink platform relies on Event-Driven Architecture for communication between services:

- Shop Boundary
- OMS
- Billing
- AB Platform
- Feed Service
- Pricer
- Proxy
- Auth Boundary
- Botfarm
- Currency Service
- Newsletter
- Many others

Originally, event names were created ad-hoc by developers without a unified standard.  
This led to inconsistencies:

- `order.created`, `orderCreated`, `order_new`
- broken analytics/CDC pipelines
- difficulties integrating with Kafka subscriptions
- problems with retry / DLQ policies
- inability to validate or version messages
- ambiguity in cross-language consumers (Go / TS / Python)

We are introducing a new `go-sdk/cqrs` package — a platform-level CQRS layer built on top of Watermill, Kafka, and protobuf.  
To support this architecture, a stable **canonical event naming** standard is required.

## Decision

Define a single mandatory event/command naming strategy for the entire Shortlink platform.

### Canonical naming schemas:

### **Events**

```
{service}.{aggregate}.{event}.{version}
```

### **Commands**

```
{service}.{command}.{version}
```

All parts MUST be lowercase, dot-separated.

Examples:

```
billing.order.created.v1
billing.create_payment.v1
shop.goods.price_changed.v1
oms.reservation.expired.v2
newsletter.subscription.confirmed.v1
```

This standard applies to:

- Kafka topic names
- Protobuf message metadata
- Tracing
- EventBus
- CommandBus
- Any cross-service contract

## Versioning

Only major versions are supported:

```
v1, v2, v3...
```

Version is included in:

- event name
- Kafka topic
- metadata

Reasons:

- clear backward compatibility
- easy rollout of breaking schema changes
- safe coexistence of v1/v2 consumers

## Kafka Topic Naming

Kafka topics MUST match the canonical event name:

```
Topic = {service}.{aggregate}.{event}.{version}
```

No prefixes or suffixes.

Forbidden:

- `events.billing.order.created.v1`
- `shortlink.billing.order.created`
- `order.created`
- `Billing.OrderCreated.v1`

## Payload Format — Protobuf Only

All commands and events MUST be defined as protobuf messages.

Content type:

```
application/x-protobuf
```

Motivation:

- strict type contracts
- schema evolution
- cross-language compatibility (Go/TS/Python/Java)
- lowest payload size
- no JSON drift
- simple automatic marshaling in `go-sdk/cqrs`

## Metadata Requirements

All events and commands MUST contain the following metadata fields:

| Key                        | Description              |
| -------------------------- | ------------------------ |
| `shortlink.event_name`     | Canonical event name     |
| `shortlink.event_version`  | Version (`v1`, `v2`...)  |
| `shortlink.service_name`   | Producing service        |
| `shortlink.trace_id`       | OpenTelemetry trace span |
| `shortlink.span_id`        | OTel span ID             |
| `shortlink.occurred_at`    | RFC3339 timestamp        |
| `shortlink.schema_version` | Protobuf schema version  |
| `shortlink.content_type`   | `application/x-protobuf` |

These fields are injected automatically by `go-sdk/cqrs` using:

- `namer.go`
- `marshaler.go`
- `metadata.go`
- OpenTelemetry propagators

## Envelope Structure

The `go-sdk/cqrs` package wraps every event/command in a standard envelope:

```go
type EventEnvelope struct {
    Name     string
    Version  string
    Payload  any            // protobuf message
    Metadata map[string]string
}

type CommandEnvelope struct {
    Name     string
    Version  string
    Payload  any
    Metadata map[string]string
}
```

This ensures uniformity and compatibility across all services.

## Registry

All protobuf events/commands MUST be registered in the global CQRS registry:

```
cqrs.RegisterEvent(&pb.OrderCreated{})
cqrs.RegisterCommand(&pb.CreateOrder{})
```

The registry ensures:

- automatic unmarshalling
- type safety
- backward-compatible event handling
- tooling to validate event schemas

## Forward Compatibility

Event evolution must always remain backward-compatible and controlled.

Rules:

1. **Any breaking change requires a new version.**  
   If the payload structure changes (field removed, type changed, meaning changed), a new major version MUST be created:

```
billing.order.created.v2
```

2. **Old versions MUST continue to exist.**  
   Existing consumers must continue processing `v1` events until they migrate.

3. **Versioning applies to topics.**  
   A new version always implies a new Kafka topic:

```
billing.order.created.v1
billing.order.created.v2
```

4. **Producers may publish both versions during migration period.**

5. **Consumers decide when to upgrade.**  
   New projections/handlers can subscribe to the new version independently.

## Examples

### ✔️ Correct Event

**Protobuf:**

```proto
message OrderCreated {
string order_id = 1;
uint64 amount = 2;
}
```

### Kafka Topic:

```
billing.order.created.v1
```

### Metadata:

```
shortlink.event_name=billing.order.created.v1
shortlink.event_version=v1
shortlink.service_name=billing
shortlink.trace_id=bd12123c…
shortlink.span_id=aa9123f1…
shortlink.occurred_at=2025-03-24T10:12:00Z
shortlink.content_type=application/x-protobuf
```

### ❌ Incorrect Examples

```
order.created
ORDER_CREATED
Billing.OrderCreated
billing.order_created.v1
billing.order-created.v1
billing.order.created
```

## Consequences

### Benefits:

- Platform-wide consistency of event naming
- Predictable versioning and schema evolution
- Reliable Kafka topic structure for all services
- Clear interoperability between Go / TypeScript / Python clients
- Strong observability through standardized metadata fields
- Easier troubleshooting, DLQ processing, and replay workflows
- Fully compatible with the `go-sdk/cqrs` message registry and protobuf marshaler
- Enables automated tooling and validation of event schemas
- Reduces risk of schema drift and contract fragmentation across teams

### Drawbacks:

- Legacy services may require migration to canonical naming
- Existing Kafka topics may need aliases or phased deprecation
- Teams must align their protobuf definitions to adhere to naming conventions
- Extra effort required during initial rollout

## Next Steps

1. Update `go-sdk/cqrs/message/namer.go` to strictly apply this ADR.
2. Implement mandatory metadata injection inside `ProtoMarshaler`.
3. Add unit tests validating canonical naming, versioning, and metadata structure.
4. Migrate initial services (Billing, OMS, Shop) to the new naming format.
5. Provide developer documentation and examples for all teams.
6. Add CI checks to ensure that new protobuf events follow ADR-0002.
7. Prepare a deprecation plan for legacy event names (if any).

## Status

This ADR is considered **Accepted** and is now the authoritative event naming standard across the Shortlink platform.

The implementation begins in 2025 and is enforced through:

- `go-sdk/cqrs`
- protobuf-based payload schemas
- Watermill/Kafka transport integration

All new services must follow this ADR.  
Existing services should adopt it during their next release cycle.
