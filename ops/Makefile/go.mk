# GO TASKS =============================================================================================================

generate: ## Code generation
	# proto generation link entity
	@protoc -I. \
	--gotemplate_out=all=true,template_dir=pkg/api/graphql/template:pkg/api/graphql \
	--go_out=Minternal/api/domain/link/link.proto=.:. \
	--go-grpc_out=Minternal/api/domain/link/link.proto=.:. \
	--go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
	internal/api/domain/link/link.proto

	@protoc -I/usr/local/include -I. \
    	--gotemplate_out=all=true,template_dir=internal/api/infrastructure/store/query/template:internal/api/infrastructure/store/query \
    	internal/api/domain/link/link.proto

	# proto generation metadata entity
	@protoc -I/usr/local/include -I. \
	--go_out=Minternal/metadata/domain/rpc.proto=.:. \
	--go-grpc_out=Minternal/metadata/domain/rpc.proto=.:. \
	--go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
	internal/metadata/domain/rpc.proto

	# proto generation gRPC-web
	@protoc -I. \
	-I=pkg/api/grpc-web \
	-I=third_party/googleapis \
	--plugin=protoc-gen-grpc-gateway=${GOPATH}/bin/protoc-gen-grpc-gateway \
	--go_out=Mpkg/api/grpc-web/api.proto=./internal/proto/grpc_service_config:. \
	--go-grpc_out=Mpkg/api/grpc-web/api.proto=./internal/proto/grpc_service_config:. \
	--go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=logtostderr=true,allow_delete_body=true:. \
	--openapiv2_out=logtostderr=true:. \
	pkg/api/grpc-web/api.proto
	@mv pkg/api/grpc-web/api.swagger.json docs/api.swagger.json

	# Generate from .go code
	@go generate internal/store/postgres/postgres.go
	@go generate internal/db/mongo/mongo.go
	@go generate internal/di/wire.go
	@go generate internal/bot/di/wire.go
	@go generate pkg/csi/di/wire.go

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
