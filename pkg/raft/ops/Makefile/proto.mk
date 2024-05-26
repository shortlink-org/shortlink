# PROTO TASKS ==========================================================================================================

proto-lint: ## Check lint
	@buf ls-files
	@buf lint
	@buf breaking --against ops/proto/proto-lock.json

proto-lock: ## Lock proto dependencies
	@buf build -o ops/proto/proto-lock.json

proto-generate: ## Generate proto-files
	@buf generate \
		--path=domain \
		--template=ops/proto/domain.buf.gen.yaml
	@buf generate \
		--path=rpc \
		--template=ops/proto/rpc.buf.gen.yaml
