# CSI TASKS ============================================================================================================
csi-build: ## Build the CSI container
	@echo docker build image ${CI_REGISTRY_IMAGE}-csi:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}-csi:${CI_COMMIT_TAG} -f ops/dockerfile/csi.Dockerfile .
