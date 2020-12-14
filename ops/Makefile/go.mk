# GO TASKS =============================================================================================================

generate: ## Code generation
	# Generate from .go code
	@go generate ./...

	@make fmt

.PHONY: fmt
fmt: ## Format source using gofmt
	# Apply go fmt
	@gofmt -l -s -w cmd pkg internal

gosec: ## Golang security checker
	@gosec -exclude=G104,G110 ./...

golint: ## Linter for golang
	@golangci-lint run ./...

test: ## Run all test
	@sh ./ops/scripts/coverage.sh

bench: ## Run benchmark tests
	go test -bench ./...
