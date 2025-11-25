# CQRS Integration Tests

Integration tests for CQRS EventBus using testcontainers and real Kafka.

## Prerequisites

- Docker must be running (testcontainers requires Docker)
- Go 1.25.4+

## Running Tests

### Run all integration tests

```bash
go test -tags=integration -v ./internal/infrastructure/cqrs/...
```

### Run specific test

```bash
go test -tags=integration -v ./internal/infrastructure/cqrs/... -run TestEventBus_PublishLinkCreated
```

### Run E2E test

```bash
go test -tags=integration -v ./internal/infrastructure/cqrs/... -run TestE2E
```

## Test Structure

### `eventbus_integration_test.go`

Tests for EventBus publishing:
- `TestEventBus_PublishLinkCreated` - Tests publishing LinkCreated event
- `TestEventBus_PublishLinkUpdated` - Tests publishing LinkUpdated event
- `TestEventBus_PublishLinkDeleted` - Tests publishing LinkDeleted event

### `e2e_test.go`

End-to-end tests:
- `TestE2E_LinkCreatedToMetadataService` - Full flow: Link Service publishes → Metadata Service consumes
- `TestE2E_CanonicalTopicNames` - Verifies all events use canonical topic names

## Test Flow

1. **Setup**: Testcontainers starts Kafka container
2. **Publish**: Link Service publishes event via EventBus
3. **Verify**: Test consumes from Kafka and verifies:
   - Correct topic name (canonical format)
   - Event can be unmarshaled
   - Event data is correct

## Notes

- Tests use `//go:build integration` tag - they won't run with `go test` by default
- Tests require Docker to be running
- Each test starts its own Kafka container for isolation
- Tests verify canonical event naming (ADR-0002)

