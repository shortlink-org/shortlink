# 6. Permissions API for Zero-Trust Security

Date: 2025-01-XX

## Status

Accepted

## Context

The Proxy Service operates in a production environment and must follow Zero-Trust security principles. It is necessary to restrict application access to system resources to minimize the attack surface.

Problems without restrictions:

- The application has full access to the file system
- Access to all environment variables (including secrets)
- Ability to connect to any network hosts
- If the code is compromised, an attacker gains full access

Node.js 25+ provides a stable Permissions API that allows restricting access to:

- File system (`fs`)
- Environment variables (`env`)
- Network hosts (`net`)
- Processes (`process`)

## Decision

Use the Node.js Permissions API to restrict Proxy Service access to system resources.

### Configuration

Permissions are set through Node.js command-line flags:

```bash
node --permission \
  --allow-fs-read=/app \
  --allow-net
```

**Node.js 25 Limitations:**

- `--allow-fs-read` - supports specifying specific paths
- `--allow-net` - experimental flag, allows all network access (detailed host control is not available)
- `--allow-env` - not supported in Node.js 25, environment variables are controlled at the Kubernetes/Docker level

The `permissions.json` file is used as a permissions reference (for documentation).

### Implementation

1. **Launch with restrictions**: Uses `--permission`, `--allow-fs-read`, `--allow-net` flags
2. **Production**: Restrictions are automatically applied through the `pnpm prod` script
3. **Development**: Uses unrestricted mode (`pnpm start:permissive`) for development convenience
4. **Runtime check**: In production, permissions are logged at startup for audit

### Permissions

**File System:**

- Read: `/app`, `/app/.env`, `/app/dist`, `/app/prisma`
- Write: prohibited (proxy service does not write files)

**Environment Variables:**

- Only necessary: `PORT`, `NODE_ENV`, `SERVICE_NAME`, `SERVICE_USER_ID`, `LINK_SERVICE_GRPC_URL`, etc.
- Full list in `permissions.json`

**Network:**

- Only necessary hosts allowed: Link Service, OpenTelemetry, Pyroscope, localhost
- All other hosts are prohibited

## Consequences

### Positive

- **Zero-Trust security** - minimal privileges by default
- **Leak protection** - even if code is compromised, access is limited
- **Standards compliance** - complies with Zero-Trust architecture principles
- **Audit** - all permissions are explicitly documented in `permissions.json`
- **Simplicity** - configuration in a single file

### Negative

- **Additional configuration** - need to maintain a permissions list
- **Blocking risk** - when adding new dependencies, permissions need to be updated
- **Debugging** - access errors may not be obvious

### Risks and Mitigation

**Risk**: Forgetting to add necessary permission when adding a new dependency

- **Mitigation**: Runtime permission check in production, logging at startup

**Risk**: Too strict restrictions block legitimate operations

- **Mitigation**: Development mode without restrictions for testing, gradual tightening

## Implementation Details

### Files

- `permissions.json` - permissions reference (for documentation)
- `package.json` - scripts with `--permission`, `--allow-fs-read`, `--allow-net` flags
- `ops/Dockerfile` - using permissions in container
- `src/infrastructure/permissions.ts` - runtime permission check
- `src/application/bootstrap.ts` - logging permissions at startup

### Scripts

- `pnpm prod` - production with restrictions
- `pnpm start:permissive` - development without restrictions

## Alternatives Considered

### Alternative 1: Use only Docker security context

**Rejected** - Docker security context restricts system calls, but does not restrict Node.js access to files/env/network inside the container

### Alternative 2: Use SELinux/AppArmor

**Rejected** - more complex setup, requires root privileges, less flexible for containers

### Alternative 3: Do not use restrictions

**Rejected** - violates Zero-Trust principles, increases attack surface

## References

- [Node.js Permissions API](https://nodejs.org/api/permissions.html)
- [Zero-Trust Architecture](https://www.nist.gov/publications/zero-trust-architecture)
- `permissions.json` - permissions reference
- `src/infrastructure/permissions.ts` - implementation check
