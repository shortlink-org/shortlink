# APPLICATION TASKS ====================================================================================================
dep: ## Install dependencies for this project
	# install protoc addons
	@go install moul.io/protoc-gen-gotemplate@latest
	@go install github.com/cloudflare/cfssl/cmd/...@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/vektra/mockery/v2@v2.42.2
	@go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
	@go install github.com/shortlink-org/shortlink/pkg/protoc/protoc-gen-go-orm

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
	# Metadata service -----------------------------------------------------------------------------
	@buf generate \
		--path=boundaries/link/metadata/domain \
		--path=boundaries/link/metadata/infrastructure \
		--template=ops/proto/metadata/buf.gen.yaml \
		--config=ops/proto/metadata/buf.yaml

	# Proxy service --------------------------------------------------------------------------------
	@buf generate \
		--path=boundaries/link/proxy/src/proto/domain \
		--path=boundaries/link/proxy/src/proto/infrastructure \
		--template=ops/proto/proxy/buf.gen.yaml \
		--config=ops/proto/proxy/buf.yaml

	# Billing service -------------------------------------------------------------------------------
	@buf generate \
		--path=boundaries/billing/billing/internal/domain \
		--path=boundaries/billing/billing/internal/infrastructure \
		--template=ops/proto/billing/buf.gen.yaml \
		--config=ops/proto/billing/buf.yaml

	# Referral service ------------------------------------------------------------------------------
	@buf generate \
		--path=boundaries/marketing/referral/src/domain \
		--template=ops/proto/referral/buf.gen.yaml \
		--config=ops/proto/referral/buf.yaml

	# Eventsourcing package -------------------------------------------------------------------------
	@buf generate \
		--path=pkg/pattern/eventsourcing/domain \
		--template=ops/proto/eventsourcing/buf.gen.yaml \
		--config=ops/proto/eventsourcing/buf.yaml

	# Shortdb service -------------------------------------------------------------------------------
	@buf generate \
		--path=boundaries/shortdb/shortdb/parser/v1 \
		--path=boundaries/shortdb/shortdb/domain/query/v1 \
		--path=boundaries/shortdb/shortdb/domain/index/v1 \
		--path=boundaries/shortdb/shortdb/domain/page/v1 \
		--path=boundaries/shortdb/shortdb/domain/table/v1 \
		--path=boundaries/shortdb/shortdb/domain/field/v1 \
		--path=boundaries/shortdb/shortdb/domain/database/v1 \
		--path=boundaries/shortdb/shortdb/domain/session/v1 \
		--template=ops/proto/shortdb/buf.gen.yaml \
		--config=ops/proto/shortdb/buf.yaml

	# API-gateway service ---------------------------------------------------------------------------
	@buf generate \
		--path=boundaries/api/api-gateway/gateways/grpc-web/infrastructure/server/v1 \
		--template=ops/proto/api-gateway/buf.gen.yaml \
		--config=ops/proto/api-gateway/buf.yaml

	# Transactional Outbox package ------------------------------------------------------------------
	@buf generate \
		--path=pkg/pattern/transactional_outbox/domain \
		--template=ops/proto/transactional_outbox/buf.gen.yaml \
		--config=ops/proto/transactional_outbox/buf.yaml

	# Protoc plugins package ------------------------------------------------------------------
	@buf generate \
		--path=pkg/protoc/protoc-gen-rich-model \
		--template=ops/proto/protoc_plugin/buf.gen.yaml \
		--config=ops/proto/protoc_plugin/buf.yaml
