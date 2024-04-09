SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../ops/Makefile/common.mk

# PLUGINS ==============================================================================================================
build: ## Build Istio Extension
	@tinygo build -o main.wasm -scheduler=none -target=wasi ./shortlink/main.go
