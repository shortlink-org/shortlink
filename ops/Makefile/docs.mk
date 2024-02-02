# DOCS TASKS ===========================================================================================================
.PHONY: docs
docs: ## Generate documentation
	@swag init \
		-g server.go \
		--dir ./boundaries/api/api-gateway/application/http-chi \
		--output boundaries/api/api-gateway/docs \
		--parseDependency
