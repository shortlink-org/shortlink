# AGENTS.md

Operational guide for coding agents working in this repository.

## 1) Start Here

- This repository is a **multi-boundary monorepo** for the ShortLink project.
- Prefer **small, boundary-scoped changes** over broad cross-repo edits.
- Before changing code:
  1. Read this file.
  2. Read the closest boundary README and ADR docs.
  3. Read local Cursor rules (listed below) when working in that area.

## 2) Instruction Sources (Most Specific Wins)

When instructions differ, follow the most specific scope:

1. Files in the exact boundary/service you are editing.
2. Boundary-local Cursor rules:
   - `boundaries/proxy/.cursor/rules/architecture.md`
   - `boundaries/ui/.cursor/rules/front-end-cursor-rules.mdc`
3. Repo-wide Cursor rules:
   - `.cursor/.cursorrules`
   - `.cursor/rules/go-microservices.mdc`
4. This `AGENTS.md`.

## 3) Repository Map

Top-level boundaries in `boundaries/`:

- `link` (Go)
- `metadata` (Go)
- `bff` (Go)
- `api/api-gateway` (Go, deprecated API boundary context)
- `proxy` (TypeScript, Node.js, pnpm)
- `ui` (Next.js/TypeScript, pnpm)
- `mobile` (Expo + Nx, pnpm)
- `chrome-extension` (WXT + React/TS, pnpm)

Useful docs:

- Root contributing guide: `CONTRIBUTING.md`
- Architecture decisions: `docs/ADR/README.md`
- Boundary overview: `boundaries/README.md`

## 4) Standard Agent Workflow

1. Identify the smallest affected boundary.
2. Read `README.md` and docs in that boundary first.
3. Implement only what is required for the request.
4. Run the **narrowest relevant checks** for touched code.
5. Summarize what changed and which checks were executed.

## 5) Command Reference

### Root-level

- Show available tasks: `make help`
- Proto tooling and generation:
  - `make dep` (installs shared generators/tools)
  - `make proto-lint`
  - `make proto-generate`
- Full local stack (heavy):
  - `make up`
  - `make down`

### Go services (link, metadata, bff, api-gateway)

Run commands from the service directory whenever possible:

- `go test ./...`
- `go test -run <Name> ./...` for targeted tests
- `make help` for service-specific targets (where available)

Service-specific notes:

- `boundaries/link`: `make dev`, `make proto-lint`, `make proto-generate`, `make docs`, `make e2e`
- `boundaries/metadata`: proto/docs targets via `make help`
- `boundaries/bff`: `make dep`, `make generate`, `make docs`

### Proxy service (`boundaries/proxy`)

- Use pnpm only.
- Typical commands:
  - `pnpm install`
  - `pnpm build`
  - `pnpm lint`
  - `pnpm test`
  - `pnpm test:cov`

### UI service (`boundaries/ui`)

- Use pnpm only.
- Typical commands:
  - `pnpm install`
  - `pnpm type-check`
  - `pnpm lint`
  - `pnpm test:run`

### Mobile service (`boundaries/mobile`)

- Use pnpm only.
- Typical commands:
  - `pnpm install`
  - `pnpm nx serve shortlink`
  - `pnpm nx build shortlink`

### Chrome extension (`boundaries/chrome-extension`)

- Use pnpm only.
- Typical commands:
  - `pnpm install`
  - `pnpm dev`
  - `pnpm build`
  - `pnpm compile`

## 6) Coding Rules to Preserve

- Go code should follow clean architecture and interface-driven boundaries (`.cursor/rules/go-microservices.mdc`).
- Proxy application layer contracts must follow the DTO/use-case rules in `boundaries/proxy/.cursor/rules/architecture.md`.
- UI changes should follow front-end rules in `boundaries/ui/.cursor/rules/front-end-cursor-rules.mdc`.
- Do not edit vendored dependencies manually.

## 7) Generation and Derived Artifacts

- If `.proto` files change, run the relevant `proto-lint` and `proto-generate` targets.
- If API/docs generators are used, regenerate outputs in the same change.
- Avoid manual edits to generated files unless the generation pipeline requires post-processing (document why if so).

## 8) Git and Change Hygiene

- Check current changes before editing (`git status`).
- Keep commits focused and message text in English.
- Prefer one logical change per commit.
- Do not rewrite unrelated local edits.

## 9) Definition of Done for Agents

Before finalizing, ensure:

- The change is scoped to the requested task.
- Relevant tests/lint/type-checks for touched areas were run (or explain why not).
- Docs/comments were updated when behavior or interfaces changed.
- Any follow-up risks or assumptions are clearly called out.
