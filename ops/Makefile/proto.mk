# APPLICATION TASKS ====================================================================================================
dep: ## Install dependencies for this project
	# install protoc addons
	@go install github.com/srikrsna/protoc-gen-gotag@latest
	@go install moul.io/protoc-gen-gotemplate@latest
	@go install github.com/cloudflare/cfssl/cmd/...@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/vektra/mockery/v2@v2.33.3
	@go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go install github.com/shortlink-org/shortlink/internal/pkg/protoc/protoc-gen-go-orm

	# for NodeJS
	@npm install -g grpc-tools grpc_tools_node_protoc_ts ts-protoc-gen protoc-gen-ts @bufbuild/protobuf @bufbuild/protoc-gen-es @bufbuild/buf

	# install wire
	@go install github.com/google/wire/cmd/wire@latest

	#i18n
	@go install golang.org/x/text/cmd/gotext@latest

# PROTO TASKS ==========================================================================================================

proto-lint: ## Check lint
	@buf ls-files
	@buf lint
	@buf breaking --against ops/proto-lock.json

proto-lock: ## Lock proto dependencies
	@buf build -o ops/proto-lock.json

proto-generate: ## Generate proto-files
	# Link service --------------------------------------------------------------------------------
	@buf generate \
		--path=internal/boundaries/link/link/domain \
		--path=internal/boundaries/link/link/infrastructure \
		--template=ops/proto/link/buf.gen.yaml \
		--config=ops/proto/link/buf.yaml

	@buf generate \
		--path=internal/boundaries/link/link/domain \
		--path=internal/boundaries/link/link/infrastructure \
		--template=ops/proto/link/buf.gen.tag.yaml \
		--config=ops/proto/link/buf.yaml

	# Metadata service -----------------------------------------------------------------------------
	@buf generate \
		--path=internal/boundaries/link/metadata/domain \
		--path=internal/boundaries/link/metadata/infrastructure \
		--template=ops/proto/metadata/buf.gen.yaml \
		--config=ops/proto/metadata/buf.yaml

	# Proxy service --------------------------------------------------------------------------------
	@buf generate \
		--path=internal/boundaries/link/proxy/src/proto/domain \
		--path=internal/boundaries/link/proxy/src/proto/infrastructure \
		--template=ops/proto/proxy/buf.gen.yaml \
		--config=ops/proto/proxy/buf.yaml

	# Billing service -------------------------------------------------------------------------------
	@buf generate \
		--path=internal/boundaries/payment/billing/domain \
		--path=internal/boundaries/payment/billing/infrastructure \
		--template=ops/proto/billing/buf.gen.yaml \
		--config=ops/proto/billing/buf.yaml

	# Referral service ------------------------------------------------------------------------------
	@buf generate \
		--path=internal/boundaries/marketing/referral/src/domain \
		--template=ops/proto/referral/buf.gen.yaml \
		--config=ops/proto/referral/buf.yaml

	# Eventsourcing service -------------------------------------------------------------------------
	@buf generate \
		--path=internal/pkg/eventsourcing/v1 \
		--template=ops/proto/eventsourcing/buf.gen.yaml \
		--config=ops/proto/eventsourcing/buf.yaml

	# Shortdb service -------------------------------------------------------------------------------
	@buf generate \
		--path=internal/boundaries/shortdb/shortdb/parser/v1 \
		--path=internal/boundaries/shortdb/shortdb/domain/query/v1 \
		--path=internal/boundaries/shortdb/shortdb/domain/index/v1 \
		--path=internal/boundaries/shortdb/shortdb/domain/page/v1 \
		--path=internal/boundaries/shortdb/shortdb/domain/table/v1 \
		--path=internal/boundaries/shortdb/shortdb/domain/field/v1 \
		--path=internal/boundaries/shortdb/shortdb/domain/database/v1 \
		--path=internal/boundaries/shortdb/shortdb/domain/session/v1 \
		--template=ops/proto/shortdb/buf.gen.yaml \
		--config=ops/proto/shortdb/buf.yaml

	# API-gateway service ---------------------------------------------------------------------------
	@buf generate \
		--path=internal/boundaries/api/api-gateway/gateways/grpc-web/infrastructure/server/v1 \
		--template=ops/proto/api-gateway/buf.gen.yaml \
		--config=ops/proto/api-gateway/buf.yaml
