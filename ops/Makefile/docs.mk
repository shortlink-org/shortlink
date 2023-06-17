# DOCS TASKS ===========================================================================================================
.PHONY: docs
docs: ## Generate documentation
	@swag init \
		-g server.go \
		--dir ./internal/services/api-gateway/application/http-chi \
		--output internal/services/api-gateway/docs \
		--parseDependency
