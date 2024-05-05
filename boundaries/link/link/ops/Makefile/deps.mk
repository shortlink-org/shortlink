# DEPENDENCIES =========================================================================================================
deps: ### Install dependencies
	@echo "Installing dependencies..."
	@go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
