# INCLUDE ==============================================================================================================
# Include Makefile
include $(ROOT_DIR)/ops/Makefile/docker.mk

### Runnings ===========================================================================================================
up: ## Run for development mode
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f $(ROOT_DIR)/docker-compose.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/database/redis/redis.yaml \
		up -d --remove-orphans --build

down: confirm ## Down docker compose
	@COMPOSE_PROFILES=dns,observability,gateway docker compose \
		-f $(ROOT_DIR)/docker-compose.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/tooling/services/coredns/coredns.yaml \
		-f $(ROOT_DIR)/ops/docker-compose/database/redis/redis.yaml \
	    down --remove-orphans
	@docker network prune -f

### Code style =========================================================================================================
lint: ## Lint code
	@cargo fmt
	@cargo clippy --fix --allow-dirty

### Testing ============================================================================================================
test: ## Run tests
    @cargo test