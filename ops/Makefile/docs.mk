# DOCS TASKS ===========================================================================================================
.PHONY: docs
docs: ## Generate documentation
	@swag init \
		-g server.go \
		--dir ./boundaries/api/api-gateway/application/http-chi \
		--output boundaries/api/api-gateway/docs \
		--parseDependency

check-link: ## Check if all links in the documentation are valid
	-npm install -g markdown-link-check
	@find . \( -type d -name '*vendor*' -o -type d -name '*node_modules*' \) -prune -o -name \*.md -print0 | xargs -0 -n1 markdown-link-check -q
