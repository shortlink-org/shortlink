SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../../ops/Makefile/common.mk

# DOCUMENTATION ========================================================================================================
docs: ### Generate documentation about use ENV variables
	@go run github.com/shortlink-org/shortlink/internal/boundaries/platform/shortctl/cmd --o ./docs/env.md --include-dir "$(PWD)" --exclude-dir "vendor"
