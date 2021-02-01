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
	@docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.36-alpine golangci-lint run ./...

test: ## Run all unit test
	@go test -coverprofile=coverage.txt -covermode atomic -race -tags=unit -v ./...

gitlab-test:  ## Run all unit test for GitLab CI
	@go test -coverprofile=coverage.txt -covermode atomic -race -tags=unit -v ./... 2>&1 | go-junit-report -set-exit-code > report.xml

bench: ## Run benchmark tests
	@go test -bench ./...
