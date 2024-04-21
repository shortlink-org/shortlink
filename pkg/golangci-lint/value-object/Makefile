SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../ops/Makefile/common.mk

build: ## Build the value-object plugin
	@go build -buildmode=plugin -o bin/value-object.so main.go
