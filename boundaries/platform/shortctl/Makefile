SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/common.mk

# CLI TASKS ============================================================================================================
build: ## Build the CLI
	@go build \
		-a \
		-ldflags "-s -w" \
		-installsuffix cgo \
		-trimpath \
		-o bin/cli ./cmd

docs: ## Generate CLI docs
	@make cli-build
	@bin/cli docs generate --o ./docs/env.md \
		--include-dir "cmd,internal,pkg" \
		--exclude-dir "vendor,node_modules,dist,ui"
