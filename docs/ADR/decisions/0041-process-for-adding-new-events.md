# 41. Process for adding new events

Date: 2025-01-26

## Status

Accepted

## Context

After migrating to `go-sdk/cqrs`, we need a standardized process for adding new events and commands to services. This ensures consistency, proper registration, and adherence to ADR-0002 canonical naming standards.

Without a clear process, developers may:

- Skip event registration in the CQRS registry
- Use incorrect naming conventions
- Miss required metadata fields
- Create events that don't follow protobuf schema evolution rules
- Fail to update consumers or documentation

## Decision

We establish a standardized step-by-step process for adding new events and commands to any service using `go-sdk/cqrs`.

### Process Overview

The process consists of 6 main steps:

1. **Define Protobuf Event/Command Message**
2. **Generate Protobuf Code**
3. **Register in CQRS Registry**
4. **Update Event Publishing/Subscribing Code**
5. **Update DI Wiring (if needed)**
6. **Test and Document**

### Step-by-Step Process

#### Step 1: Define Protobuf Event/Command Message

Create or update the protobuf file in the service's domain package:

**Location**: `boundaries/{service}/internal/domain/{domain}/v1/{domain}_events.proto`

**Example for Link Service**:

```protobuf
syntax = "proto3";

package domain.link.v1;

option go_package = "github.com/shortlink-org/shortlink/boundaries/link/internal/domain/link/v1";

import "google/protobuf/timestamp.proto";

// LinkArchived event - canonical name: link.link.archived.v1
// Published when a link is archived (soft delete)
message LinkArchived {
  // Hash of the archived link
  string hash = 1;
  // Reason for archiving
  string reason = 2;
  // OccurredAt is the timestamp when the event occurred
  google.protobuf.Timestamp occurred_at = 3;
}
```

**Requirements**:

- Follow ADR-0002 canonical naming: `{service}.{aggregate}.{event}.{version}`
- Include `occurred_at` timestamp field
- Use semantic field names (lowercase with underscores)
- Add comments describing the event purpose
- Follow protobuf schema evolution rules (no field removal, only additions)

#### Step 2: Generate Protobuf Code

Run protobuf code generation:

```bash
cd boundaries/{service}
make proto
# or
buf generate
```

Verify that `.pb.go` files are generated in the same directory.

#### Step 3: Register in CQRS Registry

Update the registry file: `boundaries/{service}/internal/infrastructure/cqrs/registry.go`

**Example**:

```go
func NewEventRegistry() (*bus.TypeRegistry, error) {
  registry := bus.NewTypeRegistry()

  // Register existing events
  if err := registry.RegisterEvent(&linkpb.LinkCreated{}); err != nil {
    return nil, err
  }
  // ... other events ...

  // Register new event
  if err := registry.RegisterEvent(&linkpb.LinkArchived{}); err != nil {
    return nil, err
  }

  return registry, nil
}
```

**Important**:

- Always check for errors when registering
- Register both events and commands if applicable
- The registry must be created before EventBus/CommandBus construction

#### Step 4: Update Event Publishing/Subscribing Code

**For Publishing Events**:

Update the usecase that triggers the event:

```go
// boundaries/{service}/internal/usecases/{usecase}/{action}.go

func (uc *UC) Archive(ctx context.Context, hash string, reason string) error {
  // Business logic...

  // Publish event using EventBus
  event := &linkpb.LinkArchived{
    Hash:       hash,
    Reason:     reason,
    OccurredAt: timestamppb.Now(),
  }

  if err := uc.eventBus.Publish(ctx, event); err != nil {
    uc.log.Error("Failed to publish link archived event",
      slog.String("error", err.Error()),
      slog.String("event_type", "link.link.archived.v1"),
      slog.String("link_hash", hash),
    )
    return err
  }

  uc.log.Info("Link archived event published successfully",
    slog.String("event_type", "link.link.archived.v1"),
    slog.String("link_hash", hash),
  )

  return nil
}
```

**For Subscribing to Events**:

Update the CQRS service event handlers:

```go
// boundaries/{service}/internal/usecases/{service}_cqrs/event.go

func (s *Service) EventHandlers(ctx context.Context) error {
  // Existing subscriptions...

  // Subscribe to new event
  if err := s.subscribeToEvent(ctx, "link.link.archived.v1", s.handleLinkArchived); err != nil {
    return err
  }

  return nil
}

func (s *Service) handleLinkArchived(ctx context.Context, msg *message.Message) error {
  // Resolve event type from registry
  eventType, ok := s.registry.ResolveEvent("link.link.archived.v1")
  if !ok {
    return errors.New("event type not found in registry")
  }

  // Unmarshal using ProtoMarshaler
  eventValue := reflect.New(eventType.Elem()).Interface()
  event, ok := eventValue.(proto.Message)
  if !ok {
    return errors.New("event does not implement proto.Message")
  }

  if err := s.marshaler.Unmarshal(msg, event); err != nil {
    return fmt.Errorf("failed to unmarshal event: %w", err)
  }

  // Type assert to specific event type
  linkArchived, ok := event.(*linkpb.LinkArchived)
  if !ok {
    return errors.New("invalid event type")
  }

  // Process event
  return s.processLinkArchived(ctx, linkArchived)
}
```

#### Step 5: Update DI Wiring (if needed)

If you added new dependencies or constructors, update Wire configuration:

```go
// boundaries/{service}/internal/di/wire.go

// CQRSSet already includes registry, no changes needed if only adding events
// If adding new services or handlers, update service sets accordingly
```

Run Wire generation:

```bash
cd boundaries/{service}
make wire
```

#### Step 6: Test and Document

**Testing**:

1. Unit tests for event creation and publishing
2. Integration tests with Kafka
3. Verify canonical naming in topic names
4. Verify metadata injection (trace_id, occurred_at, etc.)

**Documentation**:

1. Update service README if event contracts changed
2. Document event in event schema documentation
3. Update any API documentation that references events

### Cross-Service Events

When adding events that other services will consume:

1. **Publish via buf.build**: Ensure protobuf definitions are published to buf.build registry
2. **Update Consumer Services**: Other services must register the event in their registry
3. **Update Metadata Service**: If metadata service needs to consume, update `boundaries/metadata/internal/infrastructure/cqrs/registry.go`

**Example for Metadata Service consuming Link Service events**:

```go
// boundaries/metadata/internal/infrastructure/cqrs/registry.go

import (
  linkpb "buf.build/gen/go/shortlink-org/shortlink-link-link/protocolbuffers/go/domain/link/v1"
)

func NewEventRegistry() (*bus.TypeRegistry, error) {
  registry := bus.NewTypeRegistry()

  // Register Link Service events that metadata service consumes
  if err := registry.RegisterEvent(&linkpb.LinkCreated{}); err != nil {
    return nil, err
  }

  // Register new event from Link Service
  if err := registry.RegisterEvent(&linkpb.LinkArchived{}); err != nil {
    return nil, err
  }

  return registry, nil
}
```

### Versioning Events

When creating a new version of an existing event (breaking change):

1. **Create new protobuf message**: `LinkArchivedV2` (or keep same name if using versioned packages)
2. **Update canonical name**: `link.link.archived.v2`
3. **Register both versions**: Keep v1 registered for backward compatibility
4. **Dual-write period**: Publish both v1 and v2 during migration
5. **Update consumers gradually**: Migrate consumers to v2 one by one
6. **Deprecate v1**: After all consumers migrated, remove v1 registration

### Checklist

Use this checklist when adding a new event:

- [ ] Protobuf message defined following ADR-0002 naming
- [ ] Protobuf code generated successfully
- [ ] Event registered in CQRS registry
- [ ] Event publishing code updated (if producer)
- [ ] Event subscription code updated (if consumer)
- [ ] DI wiring updated (if needed)
- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Documentation updated
- [ ] Cross-service consumers updated (if applicable)
- [ ] Kafka topic created/verified (automatic via canonical naming)

## Consequences

### Benefits

- **Consistency**: All developers follow the same process
- **Type Safety**: Registry ensures type-safe event handling
- **Discoverability**: Clear process makes it easy to find where events are defined and used
- **Compliance**: Ensures adherence to ADR-0002 canonical naming
- **Maintainability**: Standardized process reduces errors and technical debt
- **Testing**: Clear testing requirements ensure event reliability

### Drawbacks

- **Process Overhead**: Multiple steps required (but necessary for correctness)
- **Learning Curve**: New developers need to learn the process
- **Cross-Service Coordination**: Adding events consumed by multiple services requires coordination

### Mitigations

- **Templates**: Create code templates for common event patterns
- **Documentation**: Maintain clear examples in service READMEs
- **CI Checks**: Add automated checks for registry registration
- **Code Reviews**: Ensure process is followed in PR reviews

## Examples

### Complete Example: Adding LinkArchived Event

**Protobuf** (`boundaries/link/internal/domain/link/v1/link_events.proto`):

```protobuf
message LinkArchived {
  string hash = 1;
  string reason = 2;
  google.protobuf.Timestamp occurred_at = 3;
}
```

**Registry** (`boundaries/link/internal/infrastructure/cqrs/registry.go`):

```go
if err := registry.RegisterEvent(&linkpb.LinkArchived{}); err != nil {
  return nil, err
}
```

**Publishing** (`boundaries/link/internal/usecases/link/archive.go`):

```go
event := &linkpb.LinkArchived{
  Hash: hash,
  Reason: reason,
  OccurredAt: timestamppb.Now(),
}
if err := uc.eventBus.Publish(ctx, event); err != nil {
  return err
}
```

**Subscription** (`boundaries/link/internal/usecases/link_cqrs/event.go`):

```go
if err := s.subscribeToEvent(ctx, "link.link.archived.v1", s.handleLinkArchived); err != nil {
  return err
}
```

## References

- [ADR-0002: Implement as event naming](./0002-implement-as-event-naming.md) - Canonical naming standard
- [go-sdk/cqrs documentation](https://github.com/shortlink-org/go-sdk/tree/main/cqrs) - CQRS package documentation
