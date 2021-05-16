# PROTO TASKS ==========================================================================================================

proto-lint: ## Check lint (use buf check)
	@buf ls-files
	@buf lint
	@buf breaking --against ops/proto-lock.json

proto-lock: ## Lock proto dependencies
	@buf build -o ops/proto-lock.json
