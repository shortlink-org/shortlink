# DOCKER TASKS =========================================================================================================
# This is the default. It can be overridden in the main Makefile after
# including docker.mk

DOCKER_USERNAME := "batazor"
DOCKER_BUILDKIT := 1
# disable becouse more images don't have signature
DOCKER_CONTENT_TRUST := 0
BUILDX_GIT_LABELS := 1
BUILDX_EXPERIMENTAL := 1
SOURCE_DATE_EPOCH := $(git log -1 --pretty=%ct)
export CURRENT_UID=$(id -u):$(id -g)

# DOCKER TASKS =========================================================================================================
docker-login: ## Docker login
	@echo docker login as ${DOCKER_USERNAME}
	@echo ${DOCKER_PASSWORD} | docker login -u ${DOCKER_USERNAME} --password-stdin

### Helpers ============================================================================================================

docker_ip: ## View docker ip and container name
	@docker ps -q | xargs docker inspect --format "{{range .NetworkSettings.Networks}}{{print .IPAddress}} {{end}}{{.Name}}"

