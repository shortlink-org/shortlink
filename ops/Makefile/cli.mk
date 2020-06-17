# CLI TASKS ============================================================================================================
cli-build:
	@go build \
		-a \
		-mod vendor \
		-installsuffix cgo \
		-trimpath \
		-o bin/cli ./cmd/cli

cli-generate-docs:
	@make cli-build
	@bin/cli docs generate --o ./docs/env.md \
		--include-dir "cmd,internal,pkg" \
		--exclude-dir "vendor,node_modules,dist,ui"
