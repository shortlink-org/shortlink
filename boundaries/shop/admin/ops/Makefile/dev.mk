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

dump: ## Dump migrations
	@python src/migration.py dumpdata goods.good > fixtures/good.json

restore: ## Restore migrations
	@python src/migration.py loaddata fixtures/good.json

# STATIC TASKS =========================================================================================================
static: ## Collect static files
	@python src/made.py collectstatic
