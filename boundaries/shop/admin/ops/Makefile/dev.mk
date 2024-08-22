# ADMIN TASKS ==========================================================================================================
dep: ## Install dependencies
	# Create a virtual environment at .venv
	uv venv

	# Install dependencies
	uv pip install -r pyproject.toml --no-deps

lock: ## Lock dependencies
	-rm requirements.txt
	@uv pip compile pyproject.toml --generate-hashes -o requirements.txt --no-deps

run: ## Run server
	@python src/manage.py runserver

test: ## Run tests
	@pytest --fixtures tests

lint: ## Run linter
	@uvx ruff format
	@uvx ruff check --fix .

# MIGRATION TASKS ======================================================================================================
migrate: ## Run migrations
	@python src/migration.py migrate
