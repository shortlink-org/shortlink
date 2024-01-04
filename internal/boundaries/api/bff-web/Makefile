SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../../ops/Makefile/common.mk

# INTEGRATION TESTS ====================================================================================================
k6:
	@k6 run ./tests/integrations/k6-link-crud.js

# DOCUMENTATION =======================================================================================================
docs: ### Prepare k6 test
	@go run github.com/shortlink-org/shortlink/internal/boundaries/platform/shortctl/cmd --o ./docs/env.md --include-dir "." --exclude-dir "vendor"

	@docker run --rm -v ${PWD}:/local/ openapitools/openapi-generator-cli generate \
		-i /local/infrastructure/http/api/api.yaml \
		-g openapi-yaml \
		-o /local/infrastructure/http/api
