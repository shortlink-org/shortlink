# GO TASKS =============================================================================================================

dep: ## Install dev-tooling used by go:generate across the boundaries
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/vektra/mockery/v2@latest
	@go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go install github.com/google/wire/cmd/wire@latest

	# i18n
	@go install golang.org/x/text/cmd/gotext@latest

# NOTE: .proto lives inside the boundaries, each of which owns its own buf module and
# ops/Makefile/proto.mk. Run proto targets there, e.g. `make -C boundaries/link proto-check`.
generate: ## Code generation
	# Generate from .go code
	@go generate -tags=wireinject ./...

	@make fmt

.PHONY: fmt
fmt: ## Format source using goimports
	# Apply go fmt
	@goimports -l -local -w internal

golint: ## Linter for golang
	@docker run --rm -it -v $(pwd):/app -w /app/ golangci/golangci-lint:v1.64.6 golangci-lint run ./internal/...

test: ## Run all unit test
	export CGO_ENABLED=1
	@go test -coverprofile=coverage.txt -covermode atomic -race -tags=unit -v ./...

bench: ## Run benchmark tests
	@go test -bench ./internal/...

godoc-serve: ## Serve documentation (godoc format) for this package at port HTTP 9090
	@godoc -http=":9090"
