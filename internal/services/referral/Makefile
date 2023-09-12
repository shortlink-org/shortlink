SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/common.mk

# REFERRAL TASKS =======================================================================================================
dep: ## Install dependencies
	@${VIRTUAL_ENV}/bin/pip install -r requirements.txt

test: ## Run tests
	@${VIRTUAL_ENV}/bin/pytest --fixtures tests

lint: ## Run linter
	@ruff check --fix .
