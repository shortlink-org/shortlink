SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../ops/Makefile/common.mk

# LANDSCAPE ============================================================================================================
build: ## Build landscape
	@landscape2 build \
      --data-file data.yml \
      --settings-file settings.yml \
      --guide-file guide.yml \
      --logos-path logos \
      --output-dir build

run: ## Run landscape
	@landscape2 serve --landscape-dir build

dev: clean build run ## Run landscape in dev mode

clean: ## Clean landscape
	-rm -rf build
