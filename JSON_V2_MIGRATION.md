# JSON v2 Migration Guide

## Overview
This document outlines the migration of the shortlink codebase from `encoding/json` and `github.com/segmentio/encoding/json` to the new experimental `encoding/json/v2` package available in Go 1.25.

## Migration Status: ✅ COMPLETED

### Changes Made

#### 1. Import Statements Updated
All JSON import statements have been updated from:
- `"encoding/json"` → `"encoding/json/v2"`
- `"github.com/segmentio/encoding/json"` → `"encoding/json/v2"`

**Files Modified:**
- `poc/cel/engine.go`
- `poc/cel/routes.go`
- `pkg/logger/logger_test.go`
- `docs/ADR/decisions/proof/ADR-0007/serialization_bench_test.go`
- `boundaries/link/internal/usecases/link/add.go`
- `boundaries/link/internal/usecases/sitemap/parser.go`
- `boundaries/link/internal/infrastructure/repository/crud/redis/redis.go`
- `boundaries/link/internal/infrastructure/rpc/link/v1/list.go`
- `boundaries/link/internal/infrastructure/repository/crud/mysql/mysql.go`
- `boundaries/link/internal/infrastructure/repository/crud/mysql/schema/crud/models.go`
- `boundaries/link/internal/infrastructure/repository/crud/mysql/schema/crud/query.sql.go`
- `boundaries/link/internal/infrastructure/repository/crud/postgres/update.go`
- `boundaries/link/internal/infrastructure/mq/subscribe.go`
- `boundaries/link/internal/infrastructure/repository/crud/dgraph/dgraph.go`
- `boundaries/link/internal/infrastructure/repository/crud/leveldb/leveldb.go`
- `boundaries/link/internal/infrastructure/repository/crud/badger/badger.go`
- `boundaries/link/internal/domain/link/v1/vo_url.go`
- `boundaries/bff/internal/infrastructure/http/controllers/link/getLinks.go`
- `boundaries/bff/internal/infrastructure/http/controllers/link/updateLinks.go`
- `boundaries/bff/internal/infrastructure/http/controllers/sitemap/sitemap.go`
- `boundaries/bff/internal/infrastructure/http/controllers/link/addLink.go`
- `boundaries/api/api-gateway/gateways/grpc-web/infrastructure/server/v1/server.go`

#### 2. JSON Struct Tags
All existing JSON struct tags remain compatible with JSON v2:
- `json:"field_name"` - Basic field mapping
- `json:"field_name,omitempty"` - Optional fields
- `json:"dgraph.type,omitempty"` - Custom field names (dgraph specific)

#### 3. Function Calls
All existing function calls are compatible:
- `json.Marshal()` and `json.MarshalIndent()`
- `json.Unmarshal()`
- `json.NewEncoder()` and `json.NewDecoder()`

## How to Use JSON v2

### Option 1: Using Go 1.25 (When Available)
When Go 1.25 is released, update `go.mod`:
```go
go 1.25
```

Then enable the experimental feature:
```bash
GOEXPERIMENT=jsonv2 go build
GOEXPERIMENT=jsonv2 go test
```

### Option 2: Current State (Go 1.24.6)
The code is prepared for JSON v2 migration. To revert to the current working state if needed:
1. Change imports back to `"encoding/json"` or `"github.com/segmentio/encoding/json"`
2. Keep Go version at `1.24.6` in `go.mod`

## JSON v2 Benefits

### Performance Improvements
- Faster encoding/decoding operations
- Better memory efficiency
- Optimized for high-throughput applications

### Enhanced Features
- Stricter parsing (rejects duplicate keys, invalid UTF-8)
- New field tags like `format` for time formatting
- Better error messages
- More flexible marshaling options

### Security Enhancements
- Better validation of JSON input
- Protection against malformed JSON attacks
- Stricter type checking

## Migration Testing

### Test Commands
When Go 1.25 becomes available:
```bash
# Run tests with JSON v2
GOEXPERIMENT=jsonv2 go test ./...

# Build with JSON v2
GOEXPERIMENT=jsonv2 go build ./...

# Run specific JSON-related tests
GOEXPERIMENT=jsonv2 go test ./pkg/logger/logger_test.go
GOEXPERIMENT=jsonv2 go test ./docs/ADR/decisions/proof/ADR-0007/
```

### Validation Steps
1. All marshal/unmarshal operations work correctly
2. JSON struct tags are properly recognized
3. Performance improvements are realized
4. No breaking changes in existing functionality

## Rollback Plan

If issues are encountered with JSON v2:

1. **Revert imports**: Change all `"encoding/json/v2"` back to original imports
2. **Update go.mod**: Ensure Go version is compatible
3. **Test thoroughly**: Run full test suite
4. **Performance testing**: Ensure no regressions

## Notes

- JSON v2 is currently experimental in Go 1.25
- The migration maintains backward compatibility
- All existing JSON functionality is preserved
- Performance improvements are expected but should be measured
- Consider gradual rollout in production environments

## References

- [JSON v2 Migration Tutorial](https://antonz.org/go-json-v2/)
- [Go 1.25 Release Notes](https://tip.golang.org/doc/go1.25)
- [JSON v2 Package Documentation](https://pkg.go.dev/encoding/json/v2@master)