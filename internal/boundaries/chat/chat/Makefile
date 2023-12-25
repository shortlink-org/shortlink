SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../../ops/Makefile/common.mk

deps: ## Install dependencies
	@mix local.hex
	@mix deps.get

migration: ## Create and migrate your database
	@mix ecto.setup

start: ## Start the application
	@echo "Open http://localhost:4000"
	@mix phx.server
