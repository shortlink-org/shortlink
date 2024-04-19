SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../ops/Makefile/common.mk

# DOCUMENTATION ========================================================================================================
docs: ### Generate documentation about use ENV variables
	@go run github.com/shortlink-org/shortlink/boundaries/platform/shortctl/cmd --o ./docs/env.md --include-dir "." --exclude-dir "vendor"

# PROTO TASKS ==========================================================================================================

proto-lint: ## Check lint
	@buf ls-files
	@buf lint
	@buf breaking --against ops/proto/proto-lock.json

proto-lock: ## Lock proto dependencies
	@buf build -o ops/proto/proto-lock.json

proto-generate: ## Generate proto-files
	# Link service ================================================================================
	# domain --------------------------------------------------------------------------------------
	@buf generate \
		--path=domain \
		--template=ops/proto/domain.buf.gen.yaml

	@buf generate \
		--path=domain \
		--template=ops/proto/buf.gen.tag.yaml

	# repository ----------------------------------------------------------------------------------
	@buf generate \
		--path=domain \
		--path=infrastructure \
		--template=ops/proto/repository.buf.gen.yaml

    # rpc -----------------------------------------------------------------------------------------
	@buf generate \
		--path=infrastructure \
		--template=ops/proto/infrastructure.buf.gen.yaml
