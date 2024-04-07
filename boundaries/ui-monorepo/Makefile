SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../ops/Makefile/common.mk

# UI-MONOREPO ==========================================================================================================
clean: ## Clean cache
	-rm -rf node_modules .nx
	-rm -rf packages/*/node_modules packages/*/.next packages/*/out packages/*/dist
