SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(SELF_DIR)

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/ops/Makefile/common.mk
include $(SELF_DIR)/ops/Makefile/docs.mk

# INTEGRATION TESTS ====================================================================================================
test-dep: ### Prepare k6
	@docker run --rm -it -e GOOS=darwin -u "$(id -u):$(id -g)" -v "${PWD}:/xk6" grafana/xk6 build \
		--with github.com/grafana/xk6-browser@latest \
		--with github.com/grafana/xk6-distributed-tracing@latest \
		--with github.com/grafana/xk6-dashboard@master

test-e2e: ### Run k6 integration tests
	@./k6 run --out web-dashboard=export=tests/test-report.html ./tests/integrations/k6-link-crud.js


# BFF ==================================================================================================================
dep: ### Get the dependencies
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go install golang.org/x/text/cmd/gotext@latest

generate: dep ### Generate BFF
	@go generate -tags=wireinject ./...
