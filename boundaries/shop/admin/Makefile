SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/common.mk

# ADMIN TASKS ==========================================================================================================
dep: ## Install dependencies
	# Create a virtual environment at .venv
	uv venv

	# Install dependencies
	uv pip install --no-cache -r requirements.txt

lock: ## Lock dependencies
	@uv pip compile pyproject.toml -o requirements.txt

run: ## Run server
	@python src/manage.py runserver

test: ## Run tests
	@pytest --fixtures tests

lint: ## Run linter
	@ruff format
	@ruff check --fix .

# MIGRATION TASKS ======================================================================================================
migrate: ## Run migrations
	@python src/manage.py migrate
