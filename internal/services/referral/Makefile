SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/common.mk

# REFERRAL TASKS =======================================================================================================
test: ## Run tests
	@pytest tests

lint: ## Run linter
	@ruff check --fix .
