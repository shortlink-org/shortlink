# 43. Protobuf workflow

Date: 2026-08-31

## Status

Accepted

## Context

The buf tooling in this repository was modern on paper — buf v2 configs, remote plugins, managed
mode, BSR dependencies — but the workflow around it had rotted, and two parts of it quietly
contradicted the registry model we thought we were following.

`ops/proto/` held nine buf configs of which eight pointed at directories deleted long ago, and the
root `ops/Makefile/proto.mk` ran `buf lint` from a directory with no `buf.yaml` at all. Breaking
changes were detected against five hand-refreshed `proto-lock.json` images committed to git, two of
which had silently diverged from each other. Nothing in CI ran buf, so none of this was visible.

More importantly, the contract distribution model was only half applied. `boundaries/proxy` kept a
local copy of link's `.proto` and generated TypeScript from it, while every one of its call sites
actually imported the BSR package — and that local copy had drifted wire-incompatibly from the
canonical contract (`field_mask` numbered 6 instead of 7, on top of a number the canonical contract
had since given to `allowed_emails`). `boundaries/link` generated metadata's types locally from a
BSR module name that does not exist.

## Decision

**Each boundary owns its buf module.** The module config lives at the boundary root, the generation
templates and make targets under `<boundary>/ops/`. There is no buf module, no `buf.work.yaml`, and
no proto make target at the repository root, because no `.proto` lives there. The root Makefile is
the make root for repo-level ops, not an orchestrator of boundaries.

**Contracts are distributed through the BSR, in one direction only.** A boundary publishes its own
module; consumers depend on the generated SDK from the registry (`buf.build/gen/go/...` in `go.mod`,
`@buf/...` in `package.json`). No boundary keeps a copy of someone else's `.proto`, and no boundary
generates someone else's contract locally.

**Publishing is manual.** `make -C <boundary> proto-push` is run by a developer after the contract
change is merged. CI deliberately does not publish.

**Breaking changes are detected against git, not against a committed image.**
`buf breaking --against '.git#branch=main,subdir=<module>'` reads the previous state of the module
out of history, so there is no image to refresh and nothing to diverge.

**CI runs the same make targets a developer runs** — `proto-lint`, `proto-format`, `proto-breaking`
— one matrix entry per module, in `.github/workflows/buf.yaml`. `buf format` is the single authority
on `.proto` style; super-linter's protolint is disabled.

**Remote plugin versions are pinned.** Unpinned `remote:` entries had already produced committed
output from four different plugin versions.

## Consequences

Generated code and the contract it comes from are committed together, and CI fails a change that
breaks the contract before it merges rather than after a consumer notices.

Because publishing is manual, a merged contract change is not visible to consumers until someone
runs `proto-push` and the consumers bump their SDK. CI protects the contract but does not police the
freshness of consumer pins; that remains a human step.

Two things stayed out of scope. `boundaries/api/api-gateway` is deprecated and was excluded
entirely. `protoc-gen-go-orm` no longer has its source in this repository, so the `*.orm.go` files
are frozen at v1.6.0 and their generation template is annotated and left disabled.
