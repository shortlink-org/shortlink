# PROTO TASKS ==========================================================================================================

proto-lint: ## Check lint
	@buf ls-files
	@buf lint
	@buf breaking --against ops/proto-lock.json

proto-lock: ## Lock proto dependencies
	@buf build -o ops/proto-lock.json

proto-generate: ## Generate proto-files
	@buf generate \
		--path=internal/services/link/domain \
		--path=internal/services/link/infrastructure \
		--template=ops/proto/link/buf.gen.yaml \
		--config=ops/proto/link/buf.yaml

	@buf generate \
		--path=internal/services/metadata/domain \
		--path=internal/services/metadata/infrastructure \
		--template=ops/proto/metadata/buf.gen.yaml \
		--config=ops/proto/metadata/buf.yaml

	@buf generate \
		--path=internal/services/proxy/src/proto/domain \
		--path=internal/services/proxy/src/proto/infrastructure \
		--template=ops/proto/proxy/buf.gen.yaml \
		--config=ops/proto/proxy/buf.yaml

	@buf generate \
		--path=internal/services/billing/domain \
		--path=internal/services/billing/infrastructure \
		--template=ops/proto/billing/buf.gen.yaml \
		--config=ops/proto/billing/buf.yaml

	@buf generate \
		--path=internal/pkg/eventsourcing/v1 \
		--template=ops/proto/eventsourcing/buf.gen.yaml \
		--config=ops/proto/eventsourcing/buf.yaml

