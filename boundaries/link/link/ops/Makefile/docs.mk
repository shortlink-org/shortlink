# DOCUMENTATION ========================================================================================================
docs: ### Generate documentation about use ENV variables
	@go run github.com/shortlink-org/shortlink/boundaries/platform/shortctl/cmd --o ./docs/env.md --include-dir "." --exclude-dir "vendor"
