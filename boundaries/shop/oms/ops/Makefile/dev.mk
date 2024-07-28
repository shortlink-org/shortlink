# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/docker.mk

### Runnings ===========================================================================================================
up: ## Run for development mode
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f $(SELF_DIR)/../../../docker-compose.yaml \
		-f $(SELF_DIR)/../../../ops/docker-compose/tooling/services/coredns/coredns.yaml \
		up -d --remove-orphans --build

down: confirm ## Down docker compose
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f $(SELF_DIR)/../../../docker-compose.yaml \
		-f $(SELF_DIR)/../../../ops/docker-compose/tooling/services/coredns/coredns.yaml \
	down --remove-orphans
	@docker network prune -f
