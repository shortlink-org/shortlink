SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../ops/Makefile/common.mk

# INTEGRATION TESTS ====================================================================================================
test-dep: ### Prepare k6
	@docker run --rm -it -e GOOS=darwin -u "$(id -u):$(id -g)" -v "${PWD}:/xk6" grafana/xk6 build \
		--with github.com/grafana/xk6-browser@latest \
		--with github.com/grafana/xk6-distributed-tracing@latest \
		--with github.com/grafana/xk6-dashboard@master

test-e2e: ### Run k6 integration tests
	@./k6 run --out web-dashboard=export=tests/test-report.html ./tests/integrations/k6-link-crud.js

# DOCUMENTATION ========================================================================================================
docs: ### Generate OpenAPI documentation
	@go run github.com/shortlink-org/shortlink/boundaries/platform/shortctl/cmd --o ./docs/env.md --include-dir "." --exclude-dir "vendor"

	@docker run --rm -v ${PWD}:/local/ openapitools/openapi-generator-cli generate \
		-i /local/internal/infrastructure/http/api/api.yaml \
		-g openapi-yaml \
		-o /local/internal/infrastructure/http/api/render \
		--skip-validate-spec

# BFF ==================================================================================================================
dep: ### Get the dependencies
	@go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go install golang.org/x/text/cmd/gotext@latest

generate: dep ### Generate BFF
	@go generate -tags=wireinject ./...
