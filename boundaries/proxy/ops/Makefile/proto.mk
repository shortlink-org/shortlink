# PROTO TASKS ==========================================================================================================

proto-generate: ## Generate proto-files
	@echo "Generating proto files..."
	@buf generate --template ops/proto/buf.gen.yaml src/proto

proto-lint: ## Check proto lint
	@echo "Checking proto lint..."
	@buf lint --config ops/proto/buf.yaml

proto-lock: ## Lock proto dependencies
	@echo "Locking proto dependencies..."
	@buf build --config ops/proto/buf.yaml -o ops/proto/proto-lock.json

