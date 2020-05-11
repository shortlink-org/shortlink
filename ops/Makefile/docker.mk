# DOCKER TASKS =========================================================================================================
CI_REGISTRY_IMAGE := batazor/${PROJECT_NAME}

docker: docker-login docker-build docker-push ## docker login > build > push

docker-login: ## Docker login
	@echo docker login as ${DOCKER_USERNAME}
	@echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin

docker-build: ## Build the container
	@echo docker build image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG} -f ops/dockerfile/shortlink.Dockerfile .

	@echo docker build image ${CI_REGISTRY_IMAGE}-logger:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}-logger:${CI_COMMIT_TAG} -f ops/dockerfile/logger.Dockerfile .

	@echo docker build image ${CI_REGISTRY_IMAGE}-ui-nuxt:${CI_COMMIT_TAG}
	@docker build -t ${CI_REGISTRY_IMAGE}-ui-nuxt:${CI_COMMIT_TAG} -f ops/dockerfile/ui-nuxt.Dockerfile .

docker-push: ## Publish the container
	@echo docker push image ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}:${CI_COMMIT_TAG}

	@echo docker push image ${CI_REGISTRY_IMAGE}-logger:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}-logger:${CI_COMMIT_TAG}

	@echo docker push image ${CI_REGISTRY_IMAGE}-ui-nuxt:${CI_COMMIT_TAG}
	@docker push ${CI_REGISTRY_IMAGE}-ui-nuxt:${CI_COMMIT_TAG}
