# 🏛 Architectural Best Practices & Rules

This document defines the architectural rules and best practices for the Proxy Service following Clean Architecture and DDD principles.

## 1. Application Boundary Contracts (ABC)

### Rule: Single Input & Output DTOs

**Every use case MUST have a SINGLE input DTO and SINGLE output DTO.**

They are declared as TypeScript interfaces in the Application Layer.

### Pattern

```typescript
// ✅ CORRECT: Interface-based DTOs in Application Layer
// src/proxy/application/dto/GetLinkRequest.ts
export interface GetLinkRequest {
  hash: string;
}

// src/proxy/application/dto/GetLinkResponse.ts
export interface GetLinkResponse {
  link: Link;
}

// src/proxy/application/use-cases/GetLinkByHashUseCase.ts
export class GetLinkByHashUseCase {
  async execute(request: GetLinkRequest): Promise<GetLinkResponse> {
    // Implementation
    return { link };
  }
}
```

### ❌ Anti-Patterns

```typescript
// ❌ WRONG: Using classes for DTOs
export class GetLinkRequest {
  constructor(public readonly hash: string) {}
}

// ❌ WRONG: Multiple inputs/outputs
async execute(hash: string, options: Options, metadata: Metadata): Promise<Link | null | Error>

// ❌ WRONG: DTOs in Infrastructure layer
// src/proxy/infrastructure/http/fastify/dto/GetLinkRequest.ts
export interface GetLinkRequest { ... }
```

### Why This Matters

1. **Stabilizes Interface**: Single, well-defined contract for each use case
2. **Simplifies Testing**: Easy to mock and verify inputs/outputs
3. **Simplifies Documentation**: Clear contract = clear documentation
4. **One Use Case = One Contract**: 1:1 mapping between use case and DTOs
5. **Simple Serialization/Validation**: DTOs are plain objects, easy to serialize/validate

### Implementation Guidelines

1. **Location**: All DTOs must be in `src/proxy/application/dto/`
2. **Format**: TypeScript interfaces (NOT classes)
3. **Naming**: `{UseCaseName}Request` and `{UseCaseName}Response`
4. **Structure**: Plain object interfaces with primitive types or domain entities
5. **Imports**: DTOs may import from Domain Layer (entities, value objects) but NOT from Infrastructure

### Examples

#### ✅ Correct Implementation

```typescript
// GetLinkByHashUseCase DTOs
export interface GetLinkRequest {
  hash: string;
}

export interface GetLinkResponse {
  link: Link; // Domain entity is OK
}

// PublishEventUseCase DTOs
export interface PublishEventRequest {
  event: DomainEvent; // Domain event is OK
}

export interface PublishEventResponse {
  success: boolean;
}
```

#### ✅ Correct Usage

```typescript
// In Application Service
const result = await this.getLinkByHashUseCase.execute({ hash: "abc123" });

// In Use Case
async execute(request: GetLinkRequest): Promise<GetLinkResponse> {
  // Process request.hash
  return { link };
}
```

#### ❌ Incorrect Patterns

```typescript
// ❌ Multiple parameters
async execute(hash: string, options: Options): Promise<GetLinkResponse>

// ❌ Using classes
export class GetLinkRequest {
  constructor(public readonly hash: string) {}
}

// ❌ DTOs with infrastructure dependencies
export interface GetLinkRequest {
  requestId: string; // HTTP-specific
  headers: Record<string, string>; // HTTP-specific
}

// ❌ DTOs in wrong layer
// src/proxy/infrastructure/http/fastify/dto/GetLinkRequest.ts
```

### File Structure

```
src/proxy/application/
├── dto/
│   ├── GetLinkRequest.ts      # ✅ Interface
│   ├── GetLinkResponse.ts     # ✅ Interface
│   └── index.ts
├── use-cases/
│   ├── GetLinkByHashUseCase.ts
│   └── PublishEventUseCase.ts
└── services/
    └── LinkApplicationService.ts
```

### Benefits

1. **Type Safety**: TypeScript ensures contracts are followed
2. **Testability**: Easy to create mock DTOs
3. **Documentation**: DTOs serve as self-documenting contracts
4. **Evolution**: Easy to evolve interfaces (add optional fields)
5. **Validation**: DTOs can be validated using Zod at infrastructure boundary

## 2. Dependency Rules

### Rule: Layer Dependencies

- **Domain Layer**: No dependencies (on nothing)
- **Application Layer**: Only depends on Domain Layer
- **Infrastructure Layer**: Depends on Domain and Application Layers
- **Interfaces Layer**: Depends on Application and Infrastructure Layers

### ✅ Correct

```typescript
// Application Layer DTO can use Domain entities
import { Link } from "../../domain/entities/Link.js";

export interface GetLinkResponse {
  link: Link; // ✅ OK: Domain entity
}
```

### ❌ Incorrect

```typescript
// Application Layer DTO cannot use Infrastructure
import { LinkDto } from "../../infrastructure/http/dto/LinkDto.js"; // ❌

export interface GetLinkResponse {
  link: LinkDto; // ❌ Wrong layer
}
```

## 3. Use Case Interface

### Rule: IUseCase Generic Interface

Every use case must implement `IUseCase<TRequest, TResponse>`:

```typescript
export interface IUseCase<TRequest, TResponse> {
  execute(request: TRequest): Promise<TResponse>;
}
```

### ✅ Correct

```typescript
export class GetLinkByHashUseCase implements IUseCase<GetLinkRequest, GetLinkResponse> {
  async execute(request: GetLinkRequest): Promise<GetLinkResponse> {
    // Implementation
  }
}
```

## 4. Naming Conventions

### Use Cases

- Name: `{Action}{Entity}UseCase`
- Examples: `GetLinkByHashUseCase`, `PublishEventUseCase`

### DTOs

- Request: `{UseCaseName}Request`
- Response: `{UseCaseName}Response`
- Location: `src/proxy/application/dto/`

### Application Services

- Name: `{Entity}ApplicationService`
- Examples: `LinkApplicationService`

## 5. Error Handling

### Rule: Use Result Type or Throw Domain Exceptions

```typescript
// Option 1: Result type (neverthrow)
async execute(request: GetLinkRequest): Promise<Result<GetLinkResponse, LinkNotFoundError>>

// Option 2: Throw domain exceptions
async execute(request: GetLinkRequest): Promise<GetLinkResponse> {
  if (!link) {
    throw new LinkNotFoundError(hash); // Domain exception
  }
}
```

## Summary Checklist

When creating a new use case, ensure:

- [ ] DTOs are TypeScript interfaces (NOT classes)
- [ ] DTOs are in `src/proxy/application/dto/`
- [ ] Use case has exactly ONE input DTO and ONE output DTO
- [ ] DTOs only use types from Domain Layer (if needed)
- [ ] Use case implements `IUseCase<TRequest, TResponse>`
- [ ] DTOs follow naming convention: `{UseCaseName}Request` / `{UseCaseName}Response`
- [ ] No infrastructure dependencies in DTOs

## References

- Clean Architecture (Uncle Bob)
- Domain-Driven Design (Eric Evans)
- Application Boundary Contracts Pattern

