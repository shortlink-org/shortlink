SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/common.mk

# REFERRAL TASKS =======================================================================================================
dep: ## Install dependencies
	@pip install -r requirements.txt

test: ## Run tests
	@pytest tests

lint: ## Run linter
	@ruff check --fix .
