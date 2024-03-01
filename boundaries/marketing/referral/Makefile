SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/common.mk

# REFERRAL TASKS =======================================================================================================
dep: ## Install dependencies
	# Create a virtual environment at .venv
	uv venv

	# Install dependencies
	uv pip install --no-cache -r requirements.txt

test: ## Run tests
	@pytest --fixtures tests

lint: ## Run linter
	@ruff format
	@ruff check --fix .
