# DOCS TASKS ===========================================================================================================
.PHONY: docs
docs: ## Generate documentation
	@swag init \
		-g server.go \
		--dir ./internal/boundaries/api/api-gateway/application/http-chi \
		--output internal/boundaries/api/api-gateway/docs \
		--parseDependency
