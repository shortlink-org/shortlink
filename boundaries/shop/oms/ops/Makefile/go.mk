# GO TASKS =============================================================================================================

generate: ## Generate go-code
	@go generate -tags=wireinject ./...
