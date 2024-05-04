SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../ops/Makefile/common.mk

# PLUGINS ==============================================================================================================
build: ## Build Istio Extension
	@GOOS=wasip1 GOARCH=wasm go build -o main.wasm ./shortlink_go/main.go
