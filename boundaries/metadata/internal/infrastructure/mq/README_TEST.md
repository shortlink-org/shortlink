# Metadata Service CQRS Integration Tests

Integration tests for Metadata Service event subscription using testcontainers and real Kafka.

## Prerequisites

- Docker must be running (testcontainers requires Docker)
- Go 1.25.4+

## Running Tests

### Run all integration tests

```bash
go test -tags=integration -v ./internal/infrastructure/mq/...
```

### Run specific test

```bash
go test -tags=integration -v ./internal/infrastructure/mq/... -run TestMetadataMQ_SubscribeLinkCreated
```

## Test Structure

### `subscribe_integration_test.go`

Tests for Metadata Service event subscription:
- `TestMetadataMQ_SubscribeLinkCreated` - Tests subscription to LinkCreated events
- `TestMetadataMQ_EventUnmarshaling` - Tests event unmarshaling using ProtoMarshaler

## Test Flow

1. **Setup**: Testcontainers starts Kafka container
2. **Subscribe**: Metadata Service subscribes to `link.link.created.v1` topic
3. **Publish**: Test publishes LinkCreated event to Kafka
4. **Verify**: Test verifies:
   - Event is received
   - Event can be unmarshaled using TypeRegistry and ProtoMarshaler
   - Event data is correct

## Notes

- Tests use `//go:build integration` tag - they won't run with `go test` by default
- Tests require Docker to be running
- Each test starts its own Kafka container for isolation
- Tests verify protobuf unmarshaling works correctly with go-sdk/cqrs

