# GO TASKS =============================================================================================================

generate: ## Code generation
	# Generate from .go code
	@go generate -tags=wireinject ./...

	@make proto-lint
	@make proto-generate
	@make fmt

.PHONY: fmt
fmt: ## Format source using gofmt
	# Apply go fmt
	@gofmt -l -s -w cmd pkg internal

gosec: ## Golang security checker
	@docker run --rm -it -v $(pwd):/app -w /app/ securego/gosec:latest /app/...

golint: ## Linter for golang
	@docker run --rm -it -v $(pwd):/app -w /app/ golangci/golangci-lint:v1.44.2-alpine golangci-lint run ./...

test: ## Run all unit test
	@go test -coverprofile=coverage.txt -covermode atomic -race -tags=unit -v ./...

bench: ## Run benchmark tests
	@go test -bench ./...

godoc-serve: ## Serve documentation (godoc format) for this package at port HTTP 9090
	@godoc -http=":9090"
