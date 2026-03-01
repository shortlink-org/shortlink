# AGENTS.md

DDD-first operating guide for coding agents in this repository.

## 1) Core Principle

Model behavior around the domain first. Protect domain language, business invariants, and bounded-context boundaries before optimizing frameworks, transport, or storage.

## 2) Bounded Contexts (Monorepo Map)

The `boundaries/` directory represents bounded contexts and related interfaces:

- `link` (Go) - core short-link lifecycle domain
- `metadata` (Go) - enrichment and screenshot/metadata domain
- `bff` (Go) - client-facing application/API composition for web
- `api/api-gateway` (Go) - deprecated API boundary context
- `proxy` (TypeScript) - redirect/runtime edge service
- `ui` (Next.js/TypeScript) - web client
- `mobile` (Expo + Nx) - mobile client
- `chrome-extension` (WXT + React/TS) - browser client

Read first:

- `boundaries/README.md`
- boundary-local `README.md`
- `docs/ADR/README.md`

## 3) Rule Precedence (Most Specific Wins)

If instructions conflict, follow:

1. Files in the exact service/context being edited
2. Boundary-local rules:
   - `boundaries/proxy/.cursor/rules/architecture.md`
   - `boundaries/ui/.cursor/rules/front-end-cursor-rules.mdc`
3. Repo-wide rules:
   - `.cursor/.cursorrules`
   - `.cursor/rules/go-microservices.mdc`
4. This document

## 4) DDD Architecture Rules

### 4.1 Layer Responsibilities

- **Domain layer**
  - Entities, Value Objects, Aggregates, Domain Services, Domain Events
  - No infrastructure dependencies
  - Business invariants enforced here
- **Application layer**
  - Use cases/application services
  - Orchestrates domain behavior and transactions
  - DTOs for boundary contracts
- **Infrastructure layer**
  - DB, MQ, HTTP clients, file systems, third-party SDKs
  - Implements repository/gateway interfaces
- **Interface/transport layer**
  - HTTP/gRPC handlers, CLI, messaging consumers
  - Map external contracts to application DTOs

Dependency direction must always point inward (toward domain).

### 4.2 Tactical DDD Guidance

- Keep aggregate roots small and consistency-focused.
- Modify aggregate state only through aggregate behavior.
- Prefer immutable Value Objects.
- Repositories should load/save aggregate roots, not random tables.
- Domain events should express business facts in ubiquitous language.
- Cross-context integration should use explicit contracts (events, APIs, ACL/adapters), never hidden coupling.

### 4.3 Ubiquitous Language

- Use terms from boundary glossary/README/ADR in names.
- If domain meaning changes, update docs and naming in the same PR.
- Avoid generic names like `manager`, `data`, `util` for domain logic.

## 5) Agent Workflow (DDD-Oriented)

1. Identify the target bounded context.
2. Identify affected aggregate/use case.
3. Write change at domain/application level first; adapt infra second.
4. Keep changes local to one context unless integration is required.
5. Validate with the narrowest relevant tests/checks.
6. Summarize domain impact (invariant, event, contract, behavior).

## 6) Commands (Validated in this repo)

### Root

- `make help`
- `make dep`
- `make proto-lint`
- `make proto-generate`
- `make up` / `make down` (heavy full stack)

### Go contexts

Run from context directory:

- `go test ./...`
- `go test -run <Name> ./...`
- `make help`

Useful local targets:

- `boundaries/link`: `make dev`, `make proto-lint`, `make proto-generate`, `make docs`, `make e2e`
- `boundaries/metadata`: proto/docs via `make help`
- `boundaries/bff`: `make dep`, `make generate`, `make docs`

### TypeScript/Frontend contexts

- `boundaries/proxy`: `pnpm install && pnpm build && pnpm lint && pnpm test`
- `boundaries/ui`: `pnpm install && pnpm type-check && pnpm lint && pnpm test:run`
- `boundaries/mobile`: `pnpm install && pnpm nx serve shortlink` (or `pnpm nx build shortlink`)
- `boundaries/chrome-extension`: `pnpm install && pnpm build && pnpm compile`

## 7) Generation and Contracts

- If `.proto` contracts change, run corresponding `proto-lint` and `proto-generate`.
- Regenerate derived artifacts in the same change when contracts/interfaces change.
- Do not hand-edit generated files unless unavoidable; document reason when doing so.

## 8) Anti-Patterns (Avoid)

- Business logic inside handlers/controllers/repositories.
- Infrastructure types leaking into domain entities/value objects.
- Cross-context direct DB coupling.
- “God” services that bypass aggregate invariants.
- Renaming domain concepts without updating ubiquitous language docs.

## 9) Git Hygiene

- Check current tree before editing: `git status`.
- Keep commits small and domain-focused.
- Use clear English commit messages.
- Never revert unrelated user changes.

## 10) Definition of Done (DDD Checklist)

Before finalizing:

- Bounded context and aggregate boundaries were respected.
- Domain invariants are preserved and tested.
- External contracts/events are updated and validated if changed.
- Relevant tests/lint/type-check were run (or explicitly justified).
- ADR/README/glossary updates are included when domain language changed.
