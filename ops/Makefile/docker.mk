# DOCKER TASKS =========================================================================================================
# This is the default. It can be overridden in the main Makefile after
# including docker.mk

# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME := shortlink

# Export such that its passed to shell functions for Docker to pick up.
export PROJECT_NAME

DOCKER_USERNAME := "batazor"
DOCKER_BUILDKIT := 1
CI_REGISTRY_IMAGE := batazor/${PROJECT_NAME}
CI_COMMIT_TAG := latest
SHORTLINK_SERVICES := api auth bot csi landing link logger metadata notify proxy ui-next

docker: docker-login docker-build docker-push ## docker login > build > push

docker-login: ## Docker login
	@echo docker login as ${DOCKER_USERNAME}
	@echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin

docker_build:
	@echo "Building ${CI_REGISTRY_IMAGE}-$(SERVICE):${CI_COMMIT_TAG}"
	@docker buildx build --platform=linux/amd64,linux/arm64 \
		--force-rm \
		--push \
		-t ${CI_REGISTRY_IMAGE}-$(SERVICE):${CI_COMMIT_TAG} \
		-f ops/dockerfile/$(SERVICE).Dockerfile .
	@docker push ${CI_REGISTRY_IMAGE}-$(SERVICE):${CI_COMMIT_TAG}
	@docker rmi ${CI_REGISTRY_IMAGE}-$(SERVICE):${CI_COMMIT_TAG}

docker-build: ## Build the container
	for i in $(SHORTLINK_SERVICES); do \
		make docker_build SERVICE=$$i; \
  	done

### Helpers ============================================================================================================

docker_ip: ## View docker ip and container name
	@docker ps -q | xargs docker inspect --format "{{range .NetworkSettings.Networks}}{{print .IPAddress}} {{end}}{{.Name}}"
