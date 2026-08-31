# PROTO TASKS ==========================================================================================================

BUF  ?= buf
HASH := \#

# Compare this module against its state on main. The `#` must go through a variable:
# GNU Make strips everything after a literal `#`, even inside an assignment.
#
# Override it when main has no comparable state yet, e.g.:
#   make proto-breaking BREAKING_AGAINST='../../.git$(HASH)ref=HEAD~1,subdir=boundaries/metadata'
BREAKING_AGAINST ?= ../../.git$(HASH)branch=main,subdir=boundaries/metadata

.PHONY: proto-lint proto-format proto-format-fix proto-breaking proto-check proto-generate proto-dep-update proto-push

proto-lint: ## Check lint
	@$(BUF) lint

proto-format: ## Check that .proto files are in canonical buf format
	@$(BUF) format --diff --exit-code

proto-format-fix: ## Rewrite .proto files in canonical buf format
	@$(BUF) format --write

proto-breaking: ## Detect breaking changes against main
	@$(BUF) breaking --against '$(BREAKING_AGAINST)'

proto-check: proto-lint proto-format proto-breaking ## Run every read-only proto check

proto-generate: ## Generate proto-files
	set -e

	# domain --------------------------------------------------------------------------------------
	@$(BUF) generate \
		--path=internal/domain \
		--template=ops/proto/domain.buf.gen.yaml

	# rpc -----------------------------------------------------------------------------------------
	@$(BUF) generate \
		--path=internal/infrastructure \
		--template=ops/proto/rpc.buf.gen.yaml

proto-dep-update: ## Refresh buf.lock from the deps declared in buf.yaml
	@$(BUF) dep update

proto-push: ## Publish this module to the BSR (run manually during development)
	@$(BUF) push
