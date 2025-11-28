# DOCUMENTATION ========================================================================================================
docs: ### Generate documentation about use ENV variables
	@go run github.com/shortlink-org/shortlink/boundaries/platform/shortctl/cmd --o ./docs/env.md --include-dir "." --exclude-dir "vendor"
	@docker run --rm -v ${PWD}:/local/ openapitools/openapi-generator-cli generate \
		-i /local/internal/infrastructure/http/api/api.yaml \
		-g openapi-yaml \
		-o /local/internal/infrastructure/http/api/render \
		--skip-validate-spec

