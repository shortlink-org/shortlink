# PROTO TASKS ==========================================================================================================

proto-lint: ## Check lint
	@buf ls-files
	@buf lint
	@buf breaking --against ops/proto/proto-lock.json

proto-lock: ## Lock proto dependencies
	@buf build -o ops/proto/proto-lock.json

proto-generate: ## Generate proto-files
    # rpc -----------------------------------------------------------------------------------------
	@buf generate \
		--path=internal/infrastructure \
		--template=ops/proto/rpc.buf.gen.yaml
