# GO TASKS =============================================================================================================

generate: ## Code generation
	# Generate from .go code
	@go generate -tags=wireinject ./...

	@make proto-lint
	@make proto-generate
	@make fmt

.PHONY: fmt
fmt: ## Format source using goimports
	# Apply go fmt
	@goimports -l -local -w internal

golint: ## Linter for golang
	@docker run --rm -it -v $(pwd):/app -w /app/ golangci/golangci-lint:v1.58.0-alpine golangci-lint run ./internal/...

test: ## Run all unit test
	export CGO_ENABLED=1
	@go test -coverprofile=coverage.txt -covermode atomic -race -tags=unit -v ./...

bench: ## Run benchmark tests
	@go test -bench ./internal/...

godoc-serve: ## Serve documentation (godoc format) for this package at port HTTP 9090
	@godoc -http=":9090"
